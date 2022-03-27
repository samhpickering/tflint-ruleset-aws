// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFmsPolicyInvalidNameRule checks the pattern is valid
type AwsFmsPolicyInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFmsPolicyInvalidNameRule returns new rule with default attributes
func NewAwsFmsPolicyInvalidNameRule() *AwsFmsPolicyInvalidNameRule {
	return &AwsFmsPolicyInvalidNameRule{
		resourceType:  "aws_fms_policy",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
	}
}

// Name returns the rule name
func (r *AwsFmsPolicyInvalidNameRule) Name() string {
	return "aws_fms_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFmsPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFmsPolicyInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFmsPolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFmsPolicyInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
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
