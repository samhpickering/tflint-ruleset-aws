// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsElastiCacheClusterInvalidSecurityGroupRule checks whether attribute value actually exists
type AwsElastiCacheClusterInvalidSecurityGroupRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheClusterInvalidSecurityGroupRule returns new rule with default attributes
func NewAwsElastiCacheClusterInvalidSecurityGroupRule() *AwsElastiCacheClusterInvalidSecurityGroupRule {
	return &AwsElastiCacheClusterInvalidSecurityGroupRule{
		resourceType:  "aws_elasticache_cluster",
		attributeName: "security_group_ids",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Name() string {
	return "aws_elasticache_cluster_invalid_security_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeSecurityGroups
func (r *AwsElastiCacheClusterInvalidSecurityGroupRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

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

		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeSecurityGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeSecurityGroups()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeSecurityGroups; %w", err)
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid security group.`, val),
					expr.Range(),
				)
			}
		})
	}

	return nil
}
