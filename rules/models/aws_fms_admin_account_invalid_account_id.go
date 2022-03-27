// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFmsAdminAccountInvalidAccountIDRule checks the pattern is valid
type AwsFmsAdminAccountInvalidAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFmsAdminAccountInvalidAccountIDRule returns new rule with default attributes
func NewAwsFmsAdminAccountInvalidAccountIDRule() *AwsFmsAdminAccountInvalidAccountIDRule {
	return &AwsFmsAdminAccountInvalidAccountIDRule{
		resourceType:  "aws_fms_admin_account",
		attributeName: "account_id",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^[0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Name() string {
	return "aws_fms_admin_account_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Check(runner tflint.Runner) error {
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
					"account_id must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"account_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]+$`),
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
