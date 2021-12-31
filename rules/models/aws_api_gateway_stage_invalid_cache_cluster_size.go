// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAPIGatewayStageInvalidCacheClusterSizeRule checks the pattern is valid
type AwsAPIGatewayStageInvalidCacheClusterSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayStageInvalidCacheClusterSizeRule returns new rule with default attributes
func NewAwsAPIGatewayStageInvalidCacheClusterSizeRule() *AwsAPIGatewayStageInvalidCacheClusterSizeRule {
	return &AwsAPIGatewayStageInvalidCacheClusterSizeRule{
		resourceType:  "aws_api_gateway_stage",
		attributeName: "cache_cluster_size",
		enum: []string{
			"0.5",
			"1.6",
			"6.1",
			"13.5",
			"28.4",
			"58.2",
			"118",
			"237",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayStageInvalidCacheClusterSizeRule) Name() string {
	return "aws_api_gateway_stage_invalid_cache_cluster_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayStageInvalidCacheClusterSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayStageInvalidCacheClusterSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayStageInvalidCacheClusterSizeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayStageInvalidCacheClusterSizeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as cache_cluster_size`, truncateLongMessage(val)),
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
