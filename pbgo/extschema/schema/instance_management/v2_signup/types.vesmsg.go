// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package v2_signup

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_signup "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/signup"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetBillingDetails().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.billing_details")
	}

	return nil
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) SourceChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for source_choice")
	}
	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) CompanyDetailsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for company_details")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema_signup.CompanyMetaValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) UserDetailsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for user_details")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema_signup.UserMetaValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AccountDetailsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for account_details")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema_signup.AccountMetaValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) BillingDetailsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for billing_details")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema_signup.BillingMetaValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["account_details"]; exists {

		vOpts := append(opts, db.WithValidateField("account_details"))
		if err := fv(ctx, m.GetAccountDetails(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["billing_details"]; exists {

		vOpts := append(opts, db.WithValidateField("billing_details"))
		if err := fv(ctx, m.GetBillingDetails(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_details"]; exists {

		vOpts := append(opts, db.WithValidateField("company_details"))
		if err := fv(ctx, m.GetCompanyDetails(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["source_choice"]; exists {
		val := m.GetSourceChoice()
		vOpts := append(opts,
			db.WithValidateField("source_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetSourceChoice().(type) {
	case *GlobalSpecType_SourcePublic:
		if fv, exists := v.FldValidators["source_choice.source_public"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourcePublic).SourcePublic
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_public"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_SourceInternalSre:
		if fv, exists := v.FldValidators["source_choice.source_internal_sre"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourceInternalSre).SourceInternalSre
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_internal_sre"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_SourceInternalScaling:
		if fv, exists := v.FldValidators["source_choice.source_internal_scaling"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourceInternalScaling).SourceInternalScaling
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_internal_scaling"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_SourcePlanTransition:
		if fv, exists := v.FldValidators["source_choice.source_plan_transition"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourcePlanTransition).SourcePlanTransition
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_plan_transition"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_SourceMsp:
		if fv, exists := v.FldValidators["source_choice.source_msp"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourceMsp).SourceMsp
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_msp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_SourceInternalSso:
		if fv, exists := v.FldValidators["source_choice.source_internal_sso"]; exists {
			val := m.GetSourceChoice().(*GlobalSpecType_SourceInternalSso).SourceInternalSso
			vOpts := append(opts,
				db.WithValidateField("source_choice"),
				db.WithValidateField("source_internal_sso"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["user_details"]; exists {

		vOpts := append(opts, db.WithValidateField("user_details"))
		if err := fv(ctx, m.GetUserDetails(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhSourceChoice := v.SourceChoiceValidationRuleHandler
	rulesSourceChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhSourceChoice(rulesSourceChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.source_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["source_choice"] = vFn

	vrhCompanyDetails := v.CompanyDetailsValidationRuleHandler
	rulesCompanyDetails := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhCompanyDetails(rulesCompanyDetails)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.company_details: %s", err)
		panic(errMsg)
	}
	v.FldValidators["company_details"] = vFn

	vrhUserDetails := v.UserDetailsValidationRuleHandler
	rulesUserDetails := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhUserDetails(rulesUserDetails)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.user_details: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_details"] = vFn

	vrhAccountDetails := v.AccountDetailsValidationRuleHandler
	rulesAccountDetails := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAccountDetails(rulesAccountDetails)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.account_details: %s", err)
		panic(errMsg)
	}
	v.FldValidators["account_details"] = vFn

	vrhBillingDetails := v.BillingDetailsValidationRuleHandler
	rulesBillingDetails := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhBillingDetails(rulesBillingDetails)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.billing_details: %s", err)
		panic(errMsg)
	}
	v.FldValidators["billing_details"] = vFn

	v.FldValidators["source_choice.source_internal_sre"] = ves_io_schema_signup.SourceInternalSreValidator().Validate
	v.FldValidators["source_choice.source_msp"] = ves_io_schema_signup.SourceMspValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}