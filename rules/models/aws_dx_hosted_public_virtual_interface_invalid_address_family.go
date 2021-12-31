// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule checks the pattern is valid
type AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule() *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule {
	return &AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule{
		resourceType:  "aws_dx_hosted_public_virtual_interface",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule) Name() string {
	return "aws_dx_hosted_public_virtual_interface_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as address_family`, truncateLongMessage(val)),
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
