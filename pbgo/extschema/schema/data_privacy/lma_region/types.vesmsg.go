// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package lma_region

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

func (m *ElasticParams) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ElasticParams) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ElasticParams) DeepCopy() *ElasticParams {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ElasticParams{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ElasticParams) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ElasticParams) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ElasticParamsValidator().Validate(ctx, m, opts...)
}

type ValidateElasticParams struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateElasticParams) UrlsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for urls")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for urls")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated urls")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items urls")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateElasticParams) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ElasticParams)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ElasticParams got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["urls"]; exists {
		vOpts := append(opts, db.WithValidateField("urls"))
		if err := fv(ctx, m.GetUrls(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultElasticParamsValidator = func() *ValidateElasticParams {
	v := &ValidateElasticParams{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhUrls := v.UrlsValidationRuleHandler
	rulesUrls := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "8096",
	}
	vFn, err = vrhUrls(rulesUrls)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ElasticParams.urls: %s", err)
		panic(errMsg)
	}
	v.FldValidators["urls"] = vFn

	return v
}()

func ElasticParamsValidator() db.Validator {
	return DefaultElasticParamsValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) KafkaParamsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for kafka_params")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := KafkaParamsValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) ElasticParamsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for elastic_params")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ElasticParamsValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["elastic_params"]; exists {

		vOpts := append(opts, db.WithValidateField("elastic_params"))
		if err := fv(ctx, m.GetElasticParams(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["is_default"]; exists {

		vOpts := append(opts, db.WithValidateField("is_default"))
		if err := fv(ctx, m.GetIsDefault(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["kafka_params"]; exists {

		vOpts := append(opts, db.WithValidateField("kafka_params"))
		if err := fv(ctx, m.GetKafkaParams(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhKafkaParams := v.KafkaParamsValidationRuleHandler
	rulesKafkaParams := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhKafkaParams(rulesKafkaParams)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.kafka_params: %s", err)
		panic(errMsg)
	}
	v.FldValidators["kafka_params"] = vFn

	vrhElasticParams := v.ElasticParamsValidationRuleHandler
	rulesElasticParams := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhElasticParams(rulesElasticParams)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.elastic_params: %s", err)
		panic(errMsg)
	}
	v.FldValidators["elastic_params"] = vFn

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
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

func (v *ValidateGlobalSpecType) KafkaParamsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for kafka_params")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := KafkaParamsValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ElasticParamsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for elastic_params")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ElasticParamsValidator().Validate(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["elastic_params"]; exists {

		vOpts := append(opts, db.WithValidateField("elastic_params"))
		if err := fv(ctx, m.GetElasticParams(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["is_default"]; exists {

		vOpts := append(opts, db.WithValidateField("is_default"))
		if err := fv(ctx, m.GetIsDefault(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["kafka_params"]; exists {

		vOpts := append(opts, db.WithValidateField("kafka_params"))
		if err := fv(ctx, m.GetKafkaParams(), vOpts...); err != nil {
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

	vrhKafkaParams := v.KafkaParamsValidationRuleHandler
	rulesKafkaParams := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhKafkaParams(rulesKafkaParams)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.kafka_params: %s", err)
		panic(errMsg)
	}
	v.FldValidators["kafka_params"] = vFn

	vrhElasticParams := v.ElasticParamsValidationRuleHandler
	rulesElasticParams := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhElasticParams(rulesElasticParams)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.elastic_params: %s", err)
		panic(errMsg)
	}
	v.FldValidators["elastic_params"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *KafkaParams) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *KafkaParams) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *KafkaParams) DeepCopy() *KafkaParams {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &KafkaParams{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *KafkaParams) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *KafkaParams) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return KafkaParamsValidator().Validate(ctx, m, opts...)
}

type ValidateKafkaParams struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateKafkaParams) BootstrapServersValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for bootstrap_servers")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for bootstrap_servers")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated bootstrap_servers")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items bootstrap_servers")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateKafkaParams) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*KafkaParams)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *KafkaParams got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["bootstrap_servers"]; exists {
		vOpts := append(opts, db.WithValidateField("bootstrap_servers"))
		if err := fv(ctx, m.GetBootstrapServers(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultKafkaParamsValidator = func() *ValidateKafkaParams {
	v := &ValidateKafkaParams{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhBootstrapServers := v.BootstrapServersValidationRuleHandler
	rulesBootstrapServers := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "8096",
	}
	vFn, err = vrhBootstrapServers(rulesBootstrapServers)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for KafkaParams.bootstrap_servers: %s", err)
		panic(errMsg)
	}
	v.FldValidators["bootstrap_servers"] = vFn

	return v
}()

func KafkaParamsValidator() db.Validator {
	return DefaultKafkaParamsValidator
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.ElasticParams = f.GetElasticParams()
	m.IsDefault = f.GetIsDefault()
	m.KafkaParams = f.GetKafkaParams()
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *GetSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *GetSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.ElasticParams = m1.ElasticParams
	f.IsDefault = m1.IsDefault
	f.KafkaParams = m1.KafkaParams
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
