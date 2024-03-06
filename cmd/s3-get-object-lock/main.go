// usage: s3-get-bucket-location endpoint AKID SAK region bucket
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 5 {
		log.Fatalf("not enough/too many args (has %d, needs 5)", n)
	}

	client := s3.New(s3.Options{
		BaseEndpoint: aws.String(flag.Arg(0)),
		Credentials:  credentials.NewStaticCredentialsProvider(flag.Arg(1), flag.Arg(2), ""),
		Region:       flag.Arg(3),
	})

	out, err := client.GetObjectLockConfiguration(context.Background(), &s3.GetObjectLockConfigurationInput{
		Bucket: aws.String(flag.Arg(4)),
	})
	if err != nil {
		log.Panicf("GetObjectLockConfiguration: %v", err)
	}

	if out.ObjectLockConfiguration != nil {
		fmt.Printf("ObjectLockConfiguration.ObjectLockEnabled is %v\n", out.ObjectLockConfiguration.ObjectLockEnabled)
		if out.ObjectLockConfiguration.Rule != nil {
			fmt.Printf("ObjectLockConfiguration.Rule is %v\n", out.ObjectLockConfiguration.Rule)
		} else {
			fmt.Println("ObjectLockConfiguration.Rule is <nil>")
		}
	} else {
		fmt.Println("ObjectLockConfiguration is <nil>")
	}
}
