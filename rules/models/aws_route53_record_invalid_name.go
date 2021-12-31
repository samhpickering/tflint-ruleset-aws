// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53RecordInvalidNameRule checks the pattern is valid
type AwsRoute53RecordInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53RecordInvalidNameRule returns new rule with default attributes
func NewAwsRoute53RecordInvalidNameRule() *AwsRoute53RecordInvalidNameRule {
	return &AwsRoute53RecordInvalidNameRule{
		resourceType:  "aws_route53_record",
		attributeName: "name",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsRoute53RecordInvalidNameRule) Name() string {
	return "aws_route53_record_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53RecordInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53RecordInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53RecordInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53RecordInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 1024 characters or less",
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
