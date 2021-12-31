// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmActivationInvalidIAMRoleRule checks the pattern is valid
type AwsSsmActivationInvalidIAMRoleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSsmActivationInvalidIAMRoleRule returns new rule with default attributes
func NewAwsSsmActivationInvalidIAMRoleRule() *AwsSsmActivationInvalidIAMRoleRule {
	return &AwsSsmActivationInvalidIAMRoleRule{
		resourceType:  "aws_ssm_activation",
		attributeName: "iam_role",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsSsmActivationInvalidIAMRoleRule) Name() string {
	return "aws_ssm_activation_invalid_iam_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmActivationInvalidIAMRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmActivationInvalidIAMRoleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmActivationInvalidIAMRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmActivationInvalidIAMRoleRule) Check(runner tflint.Runner) error {
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
					"iam_role must be 64 characters or less",
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
