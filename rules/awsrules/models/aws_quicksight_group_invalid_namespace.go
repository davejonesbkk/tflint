// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsQuicksightGroupInvalidNamespaceRule checks the pattern is valid
type AwsQuicksightGroupInvalidNamespaceRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsQuicksightGroupInvalidNamespaceRule returns new rule with default attributes
func NewAwsQuicksightGroupInvalidNamespaceRule() *AwsQuicksightGroupInvalidNamespaceRule {
	return &AwsQuicksightGroupInvalidNamespaceRule{
		resourceType:  "aws_quicksight_group",
		attributeName: "namespace",
		pattern:       regexp.MustCompile(`^default$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightGroupInvalidNamespaceRule) Name() string {
	return "aws_quicksight_group_invalid_namespace"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightGroupInvalidNamespaceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightGroupInvalidNamespaceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightGroupInvalidNamespaceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightGroupInvalidNamespaceRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`namespace does not match valid pattern ^default$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}