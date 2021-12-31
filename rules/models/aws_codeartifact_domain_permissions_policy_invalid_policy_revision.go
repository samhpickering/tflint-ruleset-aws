// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule checks the pattern is valid
type AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule returns new rule with default attributes
func NewAwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule() *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule {
	return &AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule{
		resourceType:  "aws_codeartifact_domain_permissions_policy",
		attributeName: "policy_revision",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^\S+$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule) Name() string {
	return "aws_codeartifact_domain_permissions_policy_invalid_policy_revision"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyRevisionRule) Check(runner tflint.Runner) error {
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
					"policy_revision must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy_revision must be 1 characters or higher",
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
