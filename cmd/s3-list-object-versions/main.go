// usage: s3-list-object-versions AKID SAK region bucket versionid
package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 5 {
		log.Fatalf("not enough/too many args (has %d, needs 5)", n)
	}

	client := s3.New(s3.Options{
		Credentials: credentials.NewStaticCredentialsProvider(flag.Arg(0), flag.Arg(1), ""),
		Region:      flag.Arg(2),
	})

	out, err := client.ListObjectVersions(context.Background(), &s3.ListObjectVersionsInput{
		Bucket:          aws.String(flag.Arg(3)),
		VersionIdMarker: aws.String(flag.Arg(4)),
	})
	if err != nil {
		log.Panicf("error: %v", err)
	}

	spew.Config.MaxDepth = 3
	spew.Dump(out)
}
