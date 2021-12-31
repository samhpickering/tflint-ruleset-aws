// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule checks the pattern is valid
type AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule returns new rule with default attributes
func NewAwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule() *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule {
	return &AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule{
		resourceType:  "aws_ssoadmin_permission_set_inline_policy",
		attributeName: "instance_arn",
		max:           1224,
		min:           10,
		pattern:       regexp.MustCompile(`^arn:aws:sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`),
	}
}

// Name returns the rule name
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule) Name() string {
	return "aws_ssoadmin_permission_set_inline_policy_invalid_instance_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInstanceArnRule) Check(runner tflint.Runner) error {
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
					"instance_arn must be 1224 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"instance_arn must be 10 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws:sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`),
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
