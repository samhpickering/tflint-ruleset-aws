// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEfsBackupPolicyInvalidFileSystemIDRule checks the pattern is valid
type AwsEfsBackupPolicyInvalidFileSystemIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsEfsBackupPolicyInvalidFileSystemIDRule returns new rule with default attributes
func NewAwsEfsBackupPolicyInvalidFileSystemIDRule() *AwsEfsBackupPolicyInvalidFileSystemIDRule {
	return &AwsEfsBackupPolicyInvalidFileSystemIDRule{
		resourceType:  "aws_efs_backup_policy",
		attributeName: "file_system_id",
		max:           128,
		pattern:       regexp.MustCompile(`^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:file-system/fs-[0-9a-f]{8,40}|fs-[0-9a-f]{8,40})$`),
	}
}

// Name returns the rule name
func (r *AwsEfsBackupPolicyInvalidFileSystemIDRule) Name() string {
	return "aws_efs_backup_policy_invalid_file_system_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsBackupPolicyInvalidFileSystemIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsBackupPolicyInvalidFileSystemIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsBackupPolicyInvalidFileSystemIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsBackupPolicyInvalidFileSystemIDRule) Check(runner tflint.Runner) error {
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
					"file_system_id must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:file-system/fs-[0-9a-f]{8,40}|fs-[0-9a-f]{8,40})$`),
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
