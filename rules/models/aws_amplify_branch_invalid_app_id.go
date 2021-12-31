// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyBranchInvalidAppIDRule checks the pattern is valid
type AwsAmplifyBranchInvalidAppIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyBranchInvalidAppIDRule returns new rule with default attributes
func NewAwsAmplifyBranchInvalidAppIDRule() *AwsAmplifyBranchInvalidAppIDRule {
	return &AwsAmplifyBranchInvalidAppIDRule{
		resourceType:  "aws_amplify_branch",
		attributeName: "app_id",
		max:           20,
		min:           1,
		pattern:       regexp.MustCompile(`^d[a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyBranchInvalidAppIDRule) Name() string {
	return "aws_amplify_branch_invalid_app_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyBranchInvalidAppIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyBranchInvalidAppIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyBranchInvalidAppIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyBranchInvalidAppIDRule) Check(runner tflint.Runner) error {
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
					"app_id must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"app_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^d[a-z0-9]+$`),
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
