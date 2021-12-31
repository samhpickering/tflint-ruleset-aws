// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppmeshMeshInvalidNameRule checks the pattern is valid
type AwsAppmeshMeshInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshMeshInvalidNameRule returns new rule with default attributes
func NewAwsAppmeshMeshInvalidNameRule() *AwsAppmeshMeshInvalidNameRule {
	return &AwsAppmeshMeshInvalidNameRule{
		resourceType:  "aws_appmesh_mesh",
		attributeName: "name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshMeshInvalidNameRule) Name() string {
	return "aws_appmesh_mesh_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshMeshInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshMeshInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshMeshInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshMeshInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 255 characters or less",
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
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
