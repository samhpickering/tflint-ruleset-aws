// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductInvalidDescriptionRule checks the pattern is valid
type AwsServicecatalogProductInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogProductInvalidDescriptionRule returns new rule with default attributes
func NewAwsServicecatalogProductInvalidDescriptionRule() *AwsServicecatalogProductInvalidDescriptionRule {
	return &AwsServicecatalogProductInvalidDescriptionRule{
		resourceType:  "aws_servicecatalog_product",
		attributeName: "description",
		max:           8191,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductInvalidDescriptionRule) Name() string {
	return "aws_servicecatalog_product_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 8191 characters or less",
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
