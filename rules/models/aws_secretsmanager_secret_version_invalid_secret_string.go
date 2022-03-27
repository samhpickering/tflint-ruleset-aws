// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretVersionInvalidSecretStringRule checks the pattern is valid
type AwsSecretsmanagerSecretVersionInvalidSecretStringRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretVersionInvalidSecretStringRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretVersionInvalidSecretStringRule() *AwsSecretsmanagerSecretVersionInvalidSecretStringRule {
	return &AwsSecretsmanagerSecretVersionInvalidSecretStringRule{
		resourceType:  "aws_secretsmanager_secret_version",
		attributeName: "secret_string",
		max:           65536,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Name() string {
	return "aws_secretsmanager_secret_version_invalid_secret_string"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Check(runner tflint.Runner) error {
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
					"secret_string must be 65536 characters or less",
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
