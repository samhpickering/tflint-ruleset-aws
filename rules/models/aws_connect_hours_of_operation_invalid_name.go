// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConnectHoursOfOperationInvalidNameRule checks the pattern is valid
type AwsConnectHoursOfOperationInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConnectHoursOfOperationInvalidNameRule returns new rule with default attributes
func NewAwsConnectHoursOfOperationInvalidNameRule() *AwsConnectHoursOfOperationInvalidNameRule {
	return &AwsConnectHoursOfOperationInvalidNameRule{
		resourceType:  "aws_connect_hours_of_operation",
		attributeName: "name",
		max:           127,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConnectHoursOfOperationInvalidNameRule) Name() string {
	return "aws_connect_hours_of_operation_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConnectHoursOfOperationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConnectHoursOfOperationInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConnectHoursOfOperationInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConnectHoursOfOperationInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 127 characters or less",
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
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
