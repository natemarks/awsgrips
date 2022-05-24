package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// CreateSecret wrap secretsmanager.CreateSecret
func CreateSecret(secretInput *secretsmanager.CreateSecretInput) (result *secretsmanager.CreateSecretOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	result, err = client.CreateSecret(context.TODO(), secretInput)
	return result, err
}

// GetSecretString wrap secretsmanager.GetSecretValue. Assume we just want the
// sring from GetSecretValueOutput
func GetSecretString(secretID string) (secret string, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	return *result.SecretString, err
}

// DeleteSecret wrap secretsmanager.DeleteSecret
func DeleteSecret(input *secretsmanager.DeleteSecretInput) (output *secretsmanager.DeleteSecretOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return output, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	output, err = client.DeleteSecret(context.TODO(), input)
	return output, err
}
