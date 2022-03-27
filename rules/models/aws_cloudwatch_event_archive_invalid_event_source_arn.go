// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventArchiveInvalidEventSourceArnRule checks the pattern is valid
type AwsCloudwatchEventArchiveInvalidEventSourceArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchEventArchiveInvalidEventSourceArnRule returns new rule with default attributes
func NewAwsCloudwatchEventArchiveInvalidEventSourceArnRule() *AwsCloudwatchEventArchiveInvalidEventSourceArnRule {
	return &AwsCloudwatchEventArchiveInvalidEventSourceArnRule{
		resourceType:  "aws_cloudwatch_event_archive",
		attributeName: "event_source_arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventArchiveInvalidEventSourceArnRule) Name() string {
	return "aws_cloudwatch_event_archive_invalid_event_source_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventArchiveInvalidEventSourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventArchiveInvalidEventSourceArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventArchiveInvalidEventSourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventArchiveInvalidEventSourceArnRule) Check(runner tflint.Runner) error {
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
					"event_source_arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"event_source_arn must be 1 characters or higher",
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
