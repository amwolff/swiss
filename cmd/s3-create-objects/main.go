// usage: s3-create-objects [endpoint] [AKID] [SAK] [bucket] [prefix (can be empty)] [number of objects to create]
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) < 7 {
		log.Fatalln("not enough args")
	}

	objects, err := strconv.Atoi(os.Args[6])
	if err != nil {
		log.Panicf("Atoi: %v", err)
	}

	session, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Args[1]),
		Region:   aws.String("global"),
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     os.Args[2],
					SecretAccessKey: os.Args[3],
				},
			}),
	})
	if err != nil {
		log.Panicf("NewSession: %v", err)
	}

	service := s3.New(session)

	for i := 0; i < objects; i++ {
		_, err := service.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(os.Args[4]),
			Key:    aws.String(os.Args[5] + strconv.Itoa(i)),
		})
		if err != nil {
			log.Panicf("PutObject: %v", err)
		}
	}
}
