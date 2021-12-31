// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOpsworksApplicationInvalidTypeRule checks the pattern is valid
type AwsOpsworksApplicationInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksApplicationInvalidTypeRule returns new rule with default attributes
func NewAwsOpsworksApplicationInvalidTypeRule() *AwsOpsworksApplicationInvalidTypeRule {
	return &AwsOpsworksApplicationInvalidTypeRule{
		resourceType:  "aws_opsworks_application",
		attributeName: "type",
		enum: []string{
			"aws-flow-ruby",
			"java",
			"rails",
			"php",
			"nodejs",
			"static",
			"other",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksApplicationInvalidTypeRule) Name() string {
	return "aws_opsworks_application_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksApplicationInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOpsworksApplicationInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksApplicationInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksApplicationInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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
