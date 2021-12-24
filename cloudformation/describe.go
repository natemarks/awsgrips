package cloudformation

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"strings"
)

// GetStackByNameSubstring restore a given snapshot to a given instnace namd
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
