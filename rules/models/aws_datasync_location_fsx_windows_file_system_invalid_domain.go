// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule checks the pattern is valid
type AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule returns new rule with default attributes
func NewAwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule() *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule {
	return &AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule{
		resourceType:  "aws_datasync_location_fsx_windows_file_system",
		attributeName: "domain",
		max:           253,
		pattern:       regexp.MustCompile(`^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule) Name() string {
	return "aws_datasync_location_fsx_windows_file_system_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidDomainRule) Check(runner tflint.Runner) error {
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
					"domain must be 253 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$`),
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
