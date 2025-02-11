// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAccountAlternateContactInvalidTitleRule checks the pattern is valid
type AwsAccountAlternateContactInvalidTitleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAccountAlternateContactInvalidTitleRule returns new rule with default attributes
func NewAwsAccountAlternateContactInvalidTitleRule() *AwsAccountAlternateContactInvalidTitleRule {
	return &AwsAccountAlternateContactInvalidTitleRule{
		resourceType:  "aws_account_alternate_contact",
		attributeName: "title",
		max:           50,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAccountAlternateContactInvalidTitleRule) Name() string {
	return "aws_account_alternate_contact_invalid_title"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAccountAlternateContactInvalidTitleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAccountAlternateContactInvalidTitleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAccountAlternateContactInvalidTitleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAccountAlternateContactInvalidTitleRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"title must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"title must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
