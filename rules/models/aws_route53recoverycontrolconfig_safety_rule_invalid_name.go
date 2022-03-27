// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule checks the pattern is valid
type AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule returns new rule with default attributes
func NewAwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule() *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule {
	return &AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule{
		resourceType:  "aws_route53recoverycontrolconfig_safety_rule",
		attributeName: "name",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^\S+$`),
	}
}

// Name returns the rule name
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule) Name() string {
	return "aws_route53recoverycontrolconfig_safety_rule_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\S+$`),
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
