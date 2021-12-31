// Package fixtures provide functions to create and clean up AWS resources
// required to run the project tests. Most functions interact with AWS
// resources, so fixtures are often required. We create the fixtures as
// cloudformation stacks and tag them with delete: true and delete_after:
// (next day) to make it easy to clean them up in the event that they get
// left around by test failures
package fixtures

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

// EnsureFixtureUploadS3 Create an S3 bucket that's tagged for deletion. This
// function ensures that there is a s3 bucket for cloudformation template
// uploads. Like all o the other test fixture assets, it's created as a temporary
// stack (it's tagged for deletion). Because this bucket is specifically intended
// for use with test automation, the bucket name and paths should conform to some
// predictable pattern. AS an example, when I run the tests for this project
// (awsgrips) it'll create a s3 bucket that has 'awsgrips' the AWS account
// number and aws region somewhere in the name.

//[SOMETHING UNIQUE].[PROJECT].[AWS ACCOUNT NUMBER].[AWS-REGION].cloudformation-fixtures
//ex. natemarks.awsgrips.0123456789.us-east-1.cloudformation-fixtures

// NOTE: I think the region is important because CFN has to run stacks from buckets in the same region

// Having created the bucket, we're going to reuse templates created in other projects because we're efficient. There are two ways to share
//cloudformation templates:

// difficult, clean way: use the cloudformation registry
// the quick and dirty way: publish on GitHub

// The difficult, clean way assumes that the modules we're using are all ready for some kind of release

// DownloadUrl Download an url to a local file in a temporary directory and return
// the absolute path to the downloaded file
func DownloadUrl(location string) (localFile string, err error) {
	downloadDir, err := ioutil.TempDir("", "")
	u, err := url.Parse(location)
	if err != nil {
		return
	}

	pathWords := strings.Split(u.Path, "/")
	fileName := pathWords[len(pathWords)-1]
	localFile = path.Join(downloadDir, fileName)
	resp, err := http.Get(location)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	out, err := os.Create(localFile)
	if err != nil {
		return
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return
}

// TempTagCreateStackInput Given a CreateStackInput, append the 'deleteme' and deleteme_after' tags
func TempTagCreateStackInput(input *cloudformation.CreateStackInput) (err error) {
	dTag := types.Tag{
		Key:   aws.String("deleteme"),
		Value: aws.String("true"),
	}
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	date := fmt.Sprintf("%04d-%02d-%02d-", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())
	dAfterTag := types.Tag{
		Key:   aws.String("deleteme_after"),
		Value: aws.String(date),
	}
	input.Tags = append(input.Tags, dTag)
	input.Tags = append(input.Tags, dAfterTag)
	return
}
