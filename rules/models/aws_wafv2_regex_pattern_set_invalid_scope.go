// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafv2RegexPatternSetInvalidScopeRule checks the pattern is valid
type AwsWafv2RegexPatternSetInvalidScopeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsWafv2RegexPatternSetInvalidScopeRule returns new rule with default attributes
func NewAwsWafv2RegexPatternSetInvalidScopeRule() *AwsWafv2RegexPatternSetInvalidScopeRule {
	return &AwsWafv2RegexPatternSetInvalidScopeRule{
		resourceType:  "aws_wafv2_regex_pattern_set",
		attributeName: "scope",
		enum: []string{
			"CLOUDFRONT",
			"REGIONAL",
		},
	}
}

// Name returns the rule name
func (r *AwsWafv2RegexPatternSetInvalidScopeRule) Name() string {
	return "aws_wafv2_regex_pattern_set_invalid_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafv2RegexPatternSetInvalidScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafv2RegexPatternSetInvalidScopeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafv2RegexPatternSetInvalidScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafv2RegexPatternSetInvalidScopeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as scope`, truncateLongMessage(val)),
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
