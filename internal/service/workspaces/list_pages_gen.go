// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeIpGroups"; DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
)

func describeIPGroupsPages(ctx context.Context, conn workspacesiface.WorkSpacesAPI, input *workspaces.DescribeIpGroupsInput, fn func(*workspaces.DescribeIpGroupsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeIpGroupsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
