// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule checks the pattern is valid
type AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule returns new rule with default attributes
func NewAwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule() *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule {
	return &AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule{
		resourceType:  "aws_chime_voice_connector_termination",
		attributeName: "default_phone_number",
		pattern:       regexp.MustCompile(`^\+?[1-9]\d{1,14}$`),
	}
}

// Name returns the rule name
func (r *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule) Name() string {
	return "aws_chime_voice_connector_termination_invalid_default_phone_number"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsChimeVoiceConnectorTerminationInvalidDefaultPhoneNumberRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\+?[1-9]\d{1,14}$`),
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
