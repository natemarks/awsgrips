package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
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


func ListSecrets() (secretList []types.SecretListEntry, err error) {
	//var input secretsmanager.ListSecretsInput 
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return secretList, err
	}

	client := *secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.ListSecretsInput{
		Filters: []types.Filter{
			{
				Key:   "tag-key",
				Values: []string{"purpose"},
			},
			{
				Key:   "tag-value",
				Values: []string{"postgr8_test_fixture"},
			},
		},
	}


	paginator := *secretsmanager.NewListSecretsPaginator(
		&client,
		input,
		func(o *secretsmanager.ListSecretsPaginatorOptions){o.Limit=3},
)
	for paginator.HasMorePages(){
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return secretList, err
		}
		secretList = append(secretList, output.SecretList...)
	}

	return secretList, err
}
