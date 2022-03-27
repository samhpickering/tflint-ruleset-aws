// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmDevicePoolInvalidProjectArnRule checks the pattern is valid
type AwsDevicefarmDevicePoolInvalidProjectArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDevicefarmDevicePoolInvalidProjectArnRule returns new rule with default attributes
func NewAwsDevicefarmDevicePoolInvalidProjectArnRule() *AwsDevicefarmDevicePoolInvalidProjectArnRule {
	return &AwsDevicefarmDevicePoolInvalidProjectArnRule{
		resourceType:  "aws_devicefarm_device_pool",
		attributeName: "project_arn",
		max:           1011,
		min:           32,
		pattern:       regexp.MustCompile(`^arn:.+`),
	}
}

// Name returns the rule name
func (r *AwsDevicefarmDevicePoolInvalidProjectArnRule) Name() string {
	return "aws_devicefarm_device_pool_invalid_project_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmDevicePoolInvalidProjectArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmDevicePoolInvalidProjectArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmDevicePoolInvalidProjectArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmDevicePoolInvalidProjectArnRule) Check(runner tflint.Runner) error {
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
					"project_arn must be 1011 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"project_arn must be 32 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:.+`),
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
