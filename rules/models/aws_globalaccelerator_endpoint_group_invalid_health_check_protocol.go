// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule checks the pattern is valid
type AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule returns new rule with default attributes
func NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule() *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule {
	return &AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule{
		resourceType:  "aws_globalaccelerator_endpoint_group",
		attributeName: "health_check_protocol",
		enum: []string{
			"TCP",
			"HTTP",
			"HTTPS",
		},
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Name() string {
	return "aws_globalaccelerator_endpoint_group_invalid_health_check_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as health_check_protocol`, truncateLongMessage(val)),
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
