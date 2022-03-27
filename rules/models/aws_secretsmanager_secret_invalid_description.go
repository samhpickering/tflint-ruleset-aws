// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretInvalidDescriptionRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretInvalidDescriptionRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidDescriptionRule() *AwsSecretsmanagerSecretInvalidDescriptionRule {
	return &AwsSecretsmanagerSecretInvalidDescriptionRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "description",
		max:           2048,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Name() string {
	return "aws_secretsmanager_secret_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 2048 characters or less",
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
