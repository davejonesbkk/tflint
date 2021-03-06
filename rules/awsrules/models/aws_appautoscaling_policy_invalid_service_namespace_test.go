// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsAppautoscalingPolicyInvalidServiceNamespaceRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_appautoscaling_policy" "foo" {
	service_namespace = "eks"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsAppautoscalingPolicyInvalidServiceNamespaceRule(),
					Message: `"eks" is an invalid value as service_namespace`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_appautoscaling_policy" "foo" {
	service_namespace = "ecs"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsAppautoscalingPolicyInvalidServiceNamespaceRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
