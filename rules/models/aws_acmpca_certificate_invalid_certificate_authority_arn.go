// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule checks the pattern is valid
type AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAcmpcaCertificateInvalidCertificateAuthorityArnRule returns new rule with default attributes
func NewAwsAcmpcaCertificateInvalidCertificateAuthorityArnRule() *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule {
	return &AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule{
		resourceType:  "aws_acmpca_certificate",
		attributeName: "certificate_authority_arn",
		max:           200,
		min:           5,
		pattern:       regexp.MustCompile(`^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:[\w+=/,.@-]*:[0-9]*:[\w+=,.@-]+(/[\w+=,.@-]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule) Name() string {
	return "aws_acmpca_certificate_invalid_certificate_authority_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAcmpcaCertificateInvalidCertificateAuthorityArnRule) Check(runner tflint.Runner) error {
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
					"certificate_authority_arn must be 200 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"certificate_authority_arn must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:[\w+=/,.@-]*:[0-9]*:[\w+=,.@-]+(/[\w+=,.@-]+)*$`),
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
