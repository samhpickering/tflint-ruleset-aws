// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudhsmV2HsmInvalidIPAddressRule checks the pattern is valid
type AwsCloudhsmV2HsmInvalidIPAddressRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2HsmInvalidIPAddressRule returns new rule with default attributes
func NewAwsCloudhsmV2HsmInvalidIPAddressRule() *AwsCloudhsmV2HsmInvalidIPAddressRule {
	return &AwsCloudhsmV2HsmInvalidIPAddressRule{
		resourceType:  "aws_cloudhsm_v2_hsm",
		attributeName: "ip_address",
		pattern:       regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Name() string {
	return "aws_cloudhsm_v2_hsm_invalid_ip_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`),
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
