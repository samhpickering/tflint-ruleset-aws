// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightGroupInvalidGroupNameRule checks the pattern is valid
type AwsQuicksightGroupInvalidGroupNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightGroupInvalidGroupNameRule returns new rule with default attributes
func NewAwsQuicksightGroupInvalidGroupNameRule() *AwsQuicksightGroupInvalidGroupNameRule {
	return &AwsQuicksightGroupInvalidGroupNameRule{
		resourceType:  "aws_quicksight_group",
		attributeName: "group_name",
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightGroupInvalidGroupNameRule) Name() string {
	return "aws_quicksight_group_invalid_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightGroupInvalidGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightGroupInvalidGroupNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightGroupInvalidGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightGroupInvalidGroupNameRule) Check(runner tflint.Runner) error {
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
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"group_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0020}-\x{00FF}]+$`),
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
