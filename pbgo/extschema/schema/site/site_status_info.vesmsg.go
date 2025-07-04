// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package site

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CheckSiteExistRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CheckSiteExistRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CheckSiteExistRequest) DeepCopy() *CheckSiteExistRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CheckSiteExistRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CheckSiteExistRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CheckSiteExistRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CheckSiteExistRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCheckSiteExistRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCheckSiteExistRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateCheckSiteExistRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateCheckSiteExistRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateCheckSiteExistRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CheckSiteExistRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CheckSiteExistRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCheckSiteExistRequestValidator = func() *ValidateCheckSiteExistRequest {
	v := &ValidateCheckSiteExistRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CheckSiteExistRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CheckSiteExistRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CheckSiteExistRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	return v
}()

func CheckSiteExistRequestValidator() db.Validator {
	return DefaultCheckSiteExistRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CheckSiteExistResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CheckSiteExistResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CheckSiteExistResponse) DeepCopy() *CheckSiteExistResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CheckSiteExistResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CheckSiteExistResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CheckSiteExistResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CheckSiteExistResponseValidator().Validate(ctx, m, opts...)
}

type ValidateCheckSiteExistResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCheckSiteExistResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CheckSiteExistResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CheckSiteExistResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCheckSiteExistResponseValidator = func() *ValidateCheckSiteExistResponse {
	v := &ValidateCheckSiteExistResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CheckSiteExistResponseValidator() db.Validator {
	return DefaultCheckSiteExistResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SiteStatusMetricsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SiteStatusMetricsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SiteStatusMetricsRequest) DeepCopy() *SiteStatusMetricsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SiteStatusMetricsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SiteStatusMetricsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SiteStatusMetricsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SiteStatusMetricsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSiteStatusMetricsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSiteStatusMetricsRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) SiteValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) FieldSelectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(SiteStatusMetricsField)
		return int32(i)
	}
	// SiteStatusMetricsField_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, SiteStatusMetricsField_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for field_selector")
	}
	itemsValidatorFn := func(ctx context.Context, elems []SiteStatusMetricsField, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for field_selector")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]SiteStatusMetricsField)
		if !ok {
			return fmt.Errorf("Repeated validation expected []SiteStatusMetricsField, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated field_selector")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items field_selector")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SiteStatusMetricsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SiteStatusMetricsRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["field_selector"]; exists {
		vOpts := append(opts, db.WithValidateField("field_selector"))
		if err := fv(ctx, m.GetFieldSelector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		if err := fv(ctx, m.GetSite(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSiteStatusMetricsRequestValidator = func() *ValidateSiteStatusMetricsRequest {
	v := &ValidateSiteStatusMetricsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhSite := v.SiteValidationRuleHandler
	rulesSite := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSite(rulesSite)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.site: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site"] = vFn

	vrhFieldSelector := v.FieldSelectorValidationRuleHandler
	rulesFieldSelector := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhFieldSelector(rulesFieldSelector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.field_selector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field_selector"] = vFn

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func SiteStatusMetricsRequestValidator() db.Validator {
	return DefaultSiteStatusMetricsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SiteStatusMetricsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SiteStatusMetricsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SiteStatusMetricsResponse) DeepCopy() *SiteStatusMetricsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SiteStatusMetricsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SiteStatusMetricsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SiteStatusMetricsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SiteStatusMetricsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSiteStatusMetricsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSiteStatusMetricsResponse) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateSiteStatusMetricsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SiteStatusMetricsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SiteStatusMetricsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSiteStatusMetricsResponseValidator = func() *ValidateSiteStatusMetricsResponse {
	v := &ValidateSiteStatusMetricsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SiteStatusMetricsResponse.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func SiteStatusMetricsResponseValidator() db.Validator {
	return DefaultSiteStatusMetricsResponseValidator
}
