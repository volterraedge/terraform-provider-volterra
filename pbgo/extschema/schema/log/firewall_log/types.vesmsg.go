//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package firewall_log

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

func (m *AggregationRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AggregationRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AggregationRequest) DeepCopy() *AggregationRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AggregationRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AggregationRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AggregationRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AggregationRequestValidator().Validate(ctx, m, opts...)
}

type ValidateAggregationRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAggregationRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AggregationRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AggregationRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetAggregationType().(type) {
	case *AggregationRequest_DateAggregation:
		if fv, exists := v.FldValidators["aggregation_type.date_aggregation"]; exists {
			val := m.GetAggregationType().(*AggregationRequest_DateAggregation).DateAggregation
			vOpts := append(opts,
				db.WithValidateField("aggregation_type"),
				db.WithValidateField("date_aggregation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AggregationRequest_FieldAggregation:
		if fv, exists := v.FldValidators["aggregation_type.field_aggregation"]; exists {
			val := m.GetAggregationType().(*AggregationRequest_FieldAggregation).FieldAggregation
			vOpts := append(opts,
				db.WithValidateField("aggregation_type"),
				db.WithValidateField("field_aggregation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AggregationRequest_CardinalityAggregation:
		if fv, exists := v.FldValidators["aggregation_type.cardinality_aggregation"]; exists {
			val := m.GetAggregationType().(*AggregationRequest_CardinalityAggregation).CardinalityAggregation
			vOpts := append(opts,
				db.WithValidateField("aggregation_type"),
				db.WithValidateField("cardinality_aggregation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAggregationRequestValidator = func() *ValidateAggregationRequest {
	v := &ValidateAggregationRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["aggregation_type.date_aggregation"] = DateAggregationValidator().Validate
	v.FldValidators["aggregation_type.field_aggregation"] = FieldAggregationValidator().Validate
	v.FldValidators["aggregation_type.cardinality_aggregation"] = CardinalityAggregationValidator().Validate

	return v
}()

func AggregationRequestValidator() db.Validator {
	return DefaultAggregationRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CardinalityAggregation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CardinalityAggregation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CardinalityAggregation) DeepCopy() *CardinalityAggregation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CardinalityAggregation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CardinalityAggregation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CardinalityAggregation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CardinalityAggregationValidator().Validate(ctx, m, opts...)
}

type ValidateCardinalityAggregation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCardinalityAggregation) FieldValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(KeyField)
		return int32(i)
	}
	// KeyField_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, KeyField_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for field")
	}

	return validatorFn, nil
}

func (v *ValidateCardinalityAggregation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CardinalityAggregation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CardinalityAggregation got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field"]; exists {

		vOpts := append(opts, db.WithValidateField("field"))
		if err := fv(ctx, m.GetField(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCardinalityAggregationValidator = func() *ValidateCardinalityAggregation {
	v := &ValidateCardinalityAggregation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhField := v.FieldValidationRuleHandler
	rulesField := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhField(rulesField)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CardinalityAggregation.field: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field"] = vFn

	return v
}()

func CardinalityAggregationValidator() db.Validator {
	return DefaultCardinalityAggregationValidator
}

// augmented methods on protoc/std generated struct

func (m *DateAggregation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DateAggregation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DateAggregation) DeepCopy() *DateAggregation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DateAggregation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DateAggregation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DateAggregation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DateAggregationValidator().Validate(ctx, m, opts...)
}

type ValidateDateAggregation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDateAggregation) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateDateAggregation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DateAggregation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DateAggregation got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sub_aggs"]; exists {

		vOpts := append(opts, db.WithValidateField("sub_aggs"))
		for key, value := range m.GetSubAggs() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDateAggregationValidator = func() *ValidateDateAggregation {
	v := &ValidateDateAggregation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DateAggregation.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	v.FldValidators["sub_aggs"] = DateSubAggregationValidator().Validate

	return v
}()

func DateAggregationValidator() db.Validator {
	return DefaultDateAggregationValidator
}

// augmented methods on protoc/std generated struct

func (m *DateSubAggregation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DateSubAggregation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DateSubAggregation) DeepCopy() *DateSubAggregation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DateSubAggregation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DateSubAggregation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DateSubAggregation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DateSubAggregationValidator().Validate(ctx, m, opts...)
}

type ValidateDateSubAggregation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDateSubAggregation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DateSubAggregation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DateSubAggregation got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetAggregationType().(type) {
	case *DateSubAggregation_FieldAggregation:
		if fv, exists := v.FldValidators["aggregation_type.field_aggregation"]; exists {
			val := m.GetAggregationType().(*DateSubAggregation_FieldAggregation).FieldAggregation
			vOpts := append(opts,
				db.WithValidateField("aggregation_type"),
				db.WithValidateField("field_aggregation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDateSubAggregationValidator = func() *ValidateDateSubAggregation {
	v := &ValidateDateSubAggregation{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["aggregation_type.field_aggregation"] = FieldAggregationValidator().Validate

	return v
}()

func DateSubAggregationValidator() db.Validator {
	return DefaultDateSubAggregationValidator
}

// augmented methods on protoc/std generated struct

func (m *FieldAggregation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FieldAggregation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FieldAggregation) DeepCopy() *FieldAggregation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FieldAggregation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FieldAggregation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FieldAggregation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FieldAggregationValidator().Validate(ctx, m, opts...)
}

type ValidateFieldAggregation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFieldAggregation) FieldValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(KeyField)
		return int32(i)
	}
	// KeyField_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, KeyField_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for field")
	}

	return validatorFn, nil
}

func (v *ValidateFieldAggregation) TopkValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for topk")
	}

	return validatorFn, nil
}

func (v *ValidateFieldAggregation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FieldAggregation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FieldAggregation got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field"]; exists {

		vOpts := append(opts, db.WithValidateField("field"))
		if err := fv(ctx, m.GetField(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["topk"]; exists {

		vOpts := append(opts, db.WithValidateField("topk"))
		if err := fv(ctx, m.GetTopk(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFieldAggregationValidator = func() *ValidateFieldAggregation {
	v := &ValidateFieldAggregation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhField := v.FieldValidationRuleHandler
	rulesField := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhField(rulesField)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FieldAggregation.field: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field"] = vFn

	vrhTopk := v.TopkValidationRuleHandler
	rulesTopk := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "100",
	}
	vFn, err = vrhTopk(rulesTopk)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FieldAggregation.topk: %s", err)
		panic(errMsg)
	}
	v.FldValidators["topk"] = vFn

	return v
}()

func FieldAggregationValidator() db.Validator {
	return DefaultFieldAggregationValidator
}
