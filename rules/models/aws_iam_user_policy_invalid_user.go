// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMUserPolicyInvalidUserRule checks the pattern is valid
type AwsIAMUserPolicyInvalidUserRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMUserPolicyInvalidUserRule returns new rule with default attributes
func NewAwsIAMUserPolicyInvalidUserRule() *AwsIAMUserPolicyInvalidUserRule {
	return &AwsIAMUserPolicyInvalidUserRule{
		resourceType:  "aws_iam_user_policy",
		attributeName: "user",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMUserPolicyInvalidUserRule) Name() string {
	return "aws_iam_user_policy_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserPolicyInvalidUserRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMUserPolicyInvalidUserRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserPolicyInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserPolicyInvalidUserRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"user must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
