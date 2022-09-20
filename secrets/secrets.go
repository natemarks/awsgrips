package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

// UpdateSecret wrap secretsmanager.UpdateSecret
func UpdateSecret(secretInput *secretsmanager.UpdateSecretInput) (result *secretsmanager.UpdateSecretOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	result, err = client.UpdateSecret(context.TODO(), secretInput)
	return result, err
}

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

// ListSecrets Use ListSecretsInput to get a slice of secret entries
// This really just handles the pagination for me
func ListSecrets(input *secretsmanager.ListSecretsInput) (secretList []types.SecretListEntry, err error) {
	//var input secretsmanager.ListSecretsInput
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return secretList, err
	}

	client := *secretsmanager.NewFromConfig(cfg)

	paginator := *secretsmanager.NewListSecretsPaginator(
		&client,
		input,
		// use this syntax to specifiy the page size
		// func(o *secretsmanager.ListSecretsPaginatorOptions) { o.Limit = 3 },
	)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return secretList, err
		}
		secretList = append(secretList, output.SecretList...)
	}

	return secretList, err
}

// Given a secret Name, return true if a secret with a matching Name is in the list
func secretNameInList(secretName string, secretList []types.SecretListEntry) bool {
	for _, s := range secretList {
		if *s.Name == secretName {
			return true
		}
	}
	return false
}

// SecretNameExists Return true if a secret with teh give ID exists
func SecretNameExists(secretName string) (bool, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return false, err
	}
	input := &secretsmanager.ListSecretsInput{}
	client := *secretsmanager.NewFromConfig(cfg)

	paginator := *secretsmanager.NewListSecretsPaginator(
		&client,
		input,
		// use this syntax to specify the page size
		// func(o *secretsmanager.ListSecretsPaginatorOptions) { o.Limit = 3 },
	)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return false, err
		}
		if secretNameInList(secretName, output.SecretList) {
			return true, err
		}
	}

	return false, err
}
