// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule checks the pattern is valid
type AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule returns new rule with default attributes
func NewAwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule() *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule {
	return &AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule{
		resourceType:  "aws_servicecatalog_principal_portfolio_association",
		attributeName: "accept_language",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule) Name() string {
	return "aws_servicecatalog_principal_portfolio_association_invalid_accept_language"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidAcceptLanguageRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"accept_language must be 100 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
