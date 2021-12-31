// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMemoryDBClusterInvalidACLNameRule checks the pattern is valid
type AwsMemoryDBClusterInvalidACLNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	pattern       *regexp.Regexp
}

// NewAwsMemoryDBClusterInvalidACLNameRule returns new rule with default attributes
func NewAwsMemoryDBClusterInvalidACLNameRule() *AwsMemoryDBClusterInvalidACLNameRule {
	return &AwsMemoryDBClusterInvalidACLNameRule{
		resourceType:  "aws_memorydb_cluster",
		attributeName: "acl_name",
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]*$`),
	}
}

// Name returns the rule name
func (r *AwsMemoryDBClusterInvalidACLNameRule) Name() string {
	return "aws_memorydb_cluster_invalid_acl_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMemoryDBClusterInvalidACLNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMemoryDBClusterInvalidACLNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMemoryDBClusterInvalidACLNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMemoryDBClusterInvalidACLNameRule) Check(runner tflint.Runner) error {
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
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"acl_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][a-zA-Z0-9\-]*$`),
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
