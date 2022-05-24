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

var testSecretData = testSecret{
	Username: "my_db_username",
	Password: "my_db_password",
	Engine:   "postgres",
	Host:     "my_db_hostname",
	Port:     5432,
}
var gotTestData testSecret

func TestGetSecret(t *testing.T) {
	type args struct {
		secretId string
	}
	tests := []struct {
		name       string
		args       args
		wantSecret testSecret
		wantErr    bool
	}{
		{name: "valid",
			wantSecret: testSecretData,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		now := time.Now()
		secretId := fmt.Sprintf("postgr8-test-%s", fmt.Sprint(now.UnixMilli()))

		testSecret, err := json.Marshal(testSecretData)
		if err != nil {
			t.Error("failed to marshall testSecretData")
		}

		var createInput = secretsmanager.CreateSecretInput{
			Name:         aws.String(secretId),
			SecretString: aws.String(string(testSecret)),
		}

		var deleteInput = secretsmanager.DeleteSecretInput{
			SecretId:                   aws.String(secretId),
			ForceDeleteWithoutRecovery: true,
		}

		t.Run(tt.name, func(t *testing.T) {
			cso, err := secrets.CreateSecret(&createInput)
			if err != nil {
				t.Error("Failed to create secret")
			}
			t.Log(cso.ARN)
			gotSecret, err := secrets.GetSecret(secretId)
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
			// cleanup test secret
			_, err = secrets.DeleteSecret(&deleteInput)
			if err != nil {
				t.Error("failed to delete test secret")
			}
		})
	}
}
