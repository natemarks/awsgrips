package cloudformation

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

// GetStackByNameSubstring Given a partial (substring) name of a stack, return a
// slice of stack object where each stack name contains the given substring
func GetStackByNameSubstring(sub string) ([]types.Stack, error) {
	var result []types.Stack
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return result, err
	}

	CFNClient := *cloudformation.NewFromConfig(cfg)

	input := &cloudformation.DescribeStacksInput{}
	p := cloudformation.NewDescribeStacksPaginator(&CFNClient, input)
	for p.HasMorePages() {

		// next page takes a context
		page, pageErr := p.NextPage(context.TODO())
		if pageErr != nil {
			return result, pageErr
		}

		for _, st := range page.Stacks {
			if strings.Contains(*st.StackName, sub) {
				result = append(result, st)
			}
		}
	}
	return result, err
}
