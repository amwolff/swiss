// usage: s3-cp-small-bucket [-d] path/to/config.json
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type creds struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type bucket struct {
	Endpoint    string `json:"endpoint"`
	Region      string `json:"region"`
	Credentials creds  `json:"credentials"`
	Name        string `json:"name"`
}

type config struct {
	Destination bucket `json:"destination"`
	Source      bucket `json:"source"`
}

func main() {
	dryRun := flag.Bool("d", true, "print the object keys it would upload")
	flag.Parse()

	ctx := context.TODO()

	if flag.NArg() != 2 {
		log.Fatalln("not enough/too many args")
	}

	b, err := os.ReadFile(flag.Arg(1))
	if err != nil {
		log.Panicf("ReadFile: %v", err)
	}

	var runningConfig config
	if err = json.Unmarshal(b, &runningConfig); err != nil {
		log.Panicf("Unmarshal: %v", err)
	}

	dst, err := initClient(runningConfig.Destination)
	if err != nil {
		log.Panicf("initClient (destination): %v", err)
	}
	src, err := initClient(runningConfig.Source)
	if err != nil {
		log.Panicf("initClient (source): %v", err)
	}

	dstBucket, srcBucket := runningConfig.Destination.Name, runningConfig.Source.Name

	paths, err := listPaths(ctx, src, srcBucket)
	if err != nil {
		log.Panicf("listPaths: %v", err)
	}
	if err = copyObjects(ctx, dst, src, dstBucket, srcBucket, paths, *dryRun); err != nil {
		log.Panicf("copyObjects: %v", err)
	}
}

func initClient(b bucket) (*s3.S3, error) {
	session, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(b.Endpoint),
		Region:   aws.String(b.Region),
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     b.Credentials.AccessKeyID,
					SecretAccessKey: b.Credentials.SecretAccessKey,
				},
			}),
	})
	if err != nil {
		return nil, err
	}

	return s3.New(session), nil
}

func listPaths(ctx context.Context, client *s3.S3, bucket string) (paths []string, err error) {
	err = client.ListObjectsV2PagesWithContext(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}, func(out *s3.ListObjectsV2Output, b bool) bool {
		for _, o := range out.Contents {
			paths = append(paths, aws.StringValue(o.Key))
		}
		return true
	})
	return paths, err
}

func copyObjects(ctx context.Context, dstClient, srcClient *s3.S3, dstBucket, srcBucket string, paths []string, dryRun bool) error {
	for _, p := range paths {
		srcOut, err := srcClient.GetObjectWithContext(ctx, &s3.GetObjectInput{
			Bucket: aws.String(srcBucket),
			Key:    aws.String(p),
		})
		if err != nil {
			return fmt.Errorf("failed while getting %s: %w", p, err)
		}

		release := func() error {
			if err = srcOut.Body.Close(); err != nil {
				return fmt.Errorf("failed while closing %s: %w", p, err)
			}
			return nil
		}

		if dryRun {
			log.Printf("Would upload %s (%s)", p, aws.StringValue(srcOut.ETag))
			if err = release(); err != nil {
				return err
			}
			continue
		}

		log.Printf("Uploading %s (%s)", p, aws.StringValue(srcOut.ETag))

		dstOut, err := dstClient.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(srcOut.Body),
			Bucket: aws.String(dstBucket),
			Key:    aws.String(p),
		})
		if err != nil {
			return fmt.Errorf("falied while putting %s: %w", p, errors.Join(err, release()))
		}

		if err = release(); err != nil {
			return err
		}

		log.Printf("Successfully uploaded %s (%s)", p, aws.StringValue(dstOut.ETag))
	}

	return nil
}
