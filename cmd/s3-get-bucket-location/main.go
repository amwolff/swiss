// usage: s3-get-bucket-location endpoint AKID SAK bucket
package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 4 {
		log.Fatalf("not enough/too many args (has %d, needs 4)", n)
	}

	session, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(flag.Arg(0)),
		Region:   aws.String("global"),
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     flag.Arg(1),
					SecretAccessKey: flag.Arg(2),
				},
			}),
	})
	if err != nil {
		log.Panicf("NewSession: %v", err)
	}

	out, err := s3.New(session).GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(flag.Arg(3)),
	})
	if err != nil {
		log.Panicf("GetBucketLocation: %v", err)
	}

	loc := aws.StringValue(out.LocationConstraint)
	if loc == "" {
		loc = "<nil>"
	}
	log.Printf("bucket's location: %s", loc)
}
