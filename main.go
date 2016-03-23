package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	dependency "github.com/brycefisher/glide-nested-dependency"
)

func main() {
	older := getOldS3Client("us-east-1", false)
	newer := dependency.GetNewerS3Client("us-east-1", false)
	fmt.Println(newer, older)
}

// v0.9.3rc4
func getOldS3Client(region string, forcePathStyle bool) *s3.S3 {
	return s3.New(&aws.Config{
		Region:           &region,
		S3ForcePathStyle: &forcePathStyle,
	})
}
