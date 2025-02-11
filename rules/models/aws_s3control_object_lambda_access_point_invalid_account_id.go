// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule checks the pattern is valid
type AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsS3controlObjectLambdaAccessPointInvalidAccountIDRule returns new rule with default attributes
func NewAwsS3controlObjectLambdaAccessPointInvalidAccountIDRule() *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule {
	return &AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule{
		resourceType:  "aws_s3control_object_lambda_access_point",
		attributeName: "account_id",
		max:           64,
		pattern:       regexp.MustCompile(`^\d{12}$`),
	}
}

// Name returns the rule name
func (r *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule) Name() string {
	return "aws_s3control_object_lambda_access_point_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3controlObjectLambdaAccessPointInvalidAccountIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"account_id must be 64 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
