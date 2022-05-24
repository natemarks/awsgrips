package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func CreateSecret(secretInput *secretsmanager.CreateSecretInput) (result *secretsmanager.CreateSecretOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	result, err = client.CreateSecret(context.TODO(), secretInput)
	return result, err
}

// GetSecret dvbddfb
func GetSecret(secretId string) (secret string, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	return *result.SecretString, err
}

// DeleteSecret
func DeleteSecret(input *secretsmanager.DeleteSecretInput) (output *secretsmanager.DeleteSecretOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return output, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	output, err = client.DeleteSecret(context.TODO(), input)
	return output, err
}
