// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule checks the pattern is valid
type AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule returns new rule with default attributes
func NewAwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule() *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule {
	return &AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule{
		resourceType:  "aws_imagebuilder_infrastructure_configuration",
		attributeName: "subnet_id",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule) Name() string {
	return "aws_imagebuilder_infrastructure_configuration_invalid_subnet_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderInfrastructureConfigurationInvalidSubnetIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"subnet_id must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"subnet_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
