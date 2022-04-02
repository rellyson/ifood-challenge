package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func AwsConfig() aws.Config {
	r := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(r))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
