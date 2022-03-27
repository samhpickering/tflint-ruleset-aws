// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule checks the pattern is valid
type AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule returns new rule with default attributes
func NewAwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule() *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule {
	return &AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule{
		resourceType:  "aws_config_organization_conformance_pack",
		attributeName: "delivery_s3_bucket",
		max:           63,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule) Name() string {
	return "aws_config_organization_conformance_pack_invalid_delivery_s3_bucket"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3BucketRule) Check(runner tflint.Runner) error {
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
					"delivery_s3_bucket must be 63 characters or less",
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
