package secrets_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/natemarks/awsgrips/secrets"
)

// testSecret format of the secret generated by the AWS CDK when spinning up a
// new RDS instance
type testSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// define a test secret data object to create. This will be converted to JSON
// and saved as a string
var testSecretData = testSecret{
	Username: "my_db_username",
	Password: "my_db_password",
	Engine:   "postgres",
	Host:     "my_db_hostname",
	Port:     5432,
}

// define a variable to store the output of GetSecret
var gotTestData testSecret

// TestSecretFunctions create/read/delete a JSON string secret
func TestSecretFunctions(t *testing.T) {
	var listInput secretsmanager.ListSecretsInput
	tests := []struct {
		name       string
		wantSecret testSecret
		wantErr    bool
	}{
		// Just a single test case
		{name: "valid",
			wantSecret: testSecretData,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		now := time.Now()
		secretId := fmt.Sprintf("awsgrips-secret-test-%s", fmt.Sprint(now.UnixMilli()))

		// marshall test data object
		testSecret, err := json.Marshal(testSecretData)
		if err != nil {
			t.Error("failed to marshall testSecretData")
		}
		// generate CreateSecret params
		var createInput = secretsmanager.CreateSecretInput{
			Name:         aws.String(secretId),
			SecretString: aws.String(string(testSecret)),
		}

		// generate DeleteSecret params
		var deleteInput = secretsmanager.DeleteSecretInput{
			SecretId:                   aws.String(secretId),
			ForceDeleteWithoutRecovery: true,
		}

		t.Run(tt.name, func(t *testing.T) {
			// create a test secret
			cso, err := secrets.CreateSecret(&createInput)
			if err != nil {
				t.Error("Failed to create secret")
			}
			t.Log(cso.ARN)
			// get the secret we just created
			gotSecret, err := secrets.GetSecretString(secretId)
			if err != nil {
				t.Error("Failed to unmarshall JSON secret")
			}
			err = json.Unmarshal([]byte(gotSecret), &gotTestData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTestData != testSecretData {

				t.Errorf("GetSecret() = %v, want %v", gotSecret, tt.wantSecret)
			}
			listOutput, err := secrets.ListSecrets(&listInput)
			if err != nil {
				t.Error("Error listing secrets")
			}
			if len(listOutput) == 0 {
				t.Error("zero secrets found")
			}
			// cleanup test secret
			_, err = secrets.DeleteSecret(&deleteInput)
			if err != nil {
				t.Error("failed to delete test secret")
			}
		})
	}
}

func TestSecretNameExists(t *testing.T) {
	type args struct {
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "sdf", args: args{secretName: "sandbox/biometric/test/credentials"}, want: true, wantErr: false},
		{name: "sdf", args: args{secretName: "ZZZ/biometric/test/credentials"}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := secrets.SecretNameExists(tt.args.secretName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecretNameExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SecretNameExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
