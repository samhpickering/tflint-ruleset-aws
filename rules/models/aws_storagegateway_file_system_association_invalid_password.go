// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayFileSystemAssociationInvalidPasswordRule checks the pattern is valid
type AwsStoragegatewayFileSystemAssociationInvalidPasswordRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayFileSystemAssociationInvalidPasswordRule returns new rule with default attributes
func NewAwsStoragegatewayFileSystemAssociationInvalidPasswordRule() *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule {
	return &AwsStoragegatewayFileSystemAssociationInvalidPasswordRule{
		resourceType:  "aws_storagegateway_file_system_association",
		attributeName: "password",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^[ -~]+$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule) Name() string {
	return "aws_storagegateway_file_system_association_invalid_password"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayFileSystemAssociationInvalidPasswordRule) Check(runner tflint.Runner) error {
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
					"password must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"password must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`password does not match valid pattern ^[ -~]+$`,
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
