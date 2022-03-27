// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaLayerVersionPermissionInvalidLayerNameRule checks the pattern is valid
type AwsLambdaLayerVersionPermissionInvalidLayerNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaLayerVersionPermissionInvalidLayerNameRule returns new rule with default attributes
func NewAwsLambdaLayerVersionPermissionInvalidLayerNameRule() *AwsLambdaLayerVersionPermissionInvalidLayerNameRule {
	return &AwsLambdaLayerVersionPermissionInvalidLayerNameRule{
		resourceType:  "aws_lambda_layer_version_permission",
		attributeName: "layer_name",
		max:           140,
		min:           1,
		pattern:       regexp.MustCompile(`^(arn:[a-zA-Z0-9-]+:lambda:[a-zA-Z0-9-]+:\d{12}:layer:[a-zA-Z0-9-_]+)|[a-zA-Z0-9-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionPermissionInvalidLayerNameRule) Name() string {
	return "aws_lambda_layer_version_permission_invalid_layer_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionPermissionInvalidLayerNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaLayerVersionPermissionInvalidLayerNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionPermissionInvalidLayerNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionPermissionInvalidLayerNameRule) Check(runner tflint.Runner) error {
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
					"layer_name must be 140 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"layer_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(arn:[a-zA-Z0-9-]+:lambda:[a-zA-Z0-9-]+:\d{12}:layer:[a-zA-Z0-9-_]+)|[a-zA-Z0-9-_]+$`),
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
