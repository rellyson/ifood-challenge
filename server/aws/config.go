package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func AwsConfig() aws.Config {
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("AWS_ENDPOINT")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endpoint != "" && region != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
