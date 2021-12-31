// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule checks the pattern is valid
type AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule returns new rule with default attributes
func NewAwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule() *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule {
	return &AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule{
		resourceType:  "aws_storagegateway_file_system_association",
		attributeName: "gateway_arn",
		max:           500,
		min:           50,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule) Name() string {
	return "aws_storagegateway_file_system_association_invalid_gateway_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayFileSystemAssociationInvalidGatewayArnRule) Check(runner tflint.Runner) error {
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
					"gateway_arn must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"gateway_arn must be 50 characters or higher",
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
