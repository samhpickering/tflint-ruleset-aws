// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule checks the pattern is valid
type AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule returns new rule with default attributes
func NewAwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule() *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule {
	return &AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule{
		resourceType:  "aws_ecr_pull_through_cache_rule",
		attributeName: "ecr_repository_prefix",
		max:           20,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule) Name() string {
	return "aws_ecr_pull_through_cache_rule_invalid_ecr_repository_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrPullThroughCacheRuleInvalidEcrRepositoryPrefixRule) Check(runner tflint.Runner) error {
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
					"ecr_repository_prefix must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"ecr_repository_prefix must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
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
