// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package common

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

func (m *DashboardLink) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DashboardLink) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DashboardLink) DeepCopy() *DashboardLink {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DashboardLink{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DashboardLink) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DashboardLink) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DashboardLinkValidator().Validate(ctx, m, opts...)
}

type ValidateDashboardLink struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDashboardLink) KeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for key")
	}

	return validatorFn, nil
}

func (v *ValidateDashboardLink) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DashboardLink)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DashboardLink got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["log_filters"]; exists {

		vOpts := append(opts, db.WithValidateField("log_filters"))
		for idx, item := range m.GetLogFilters() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_name"]; exists {

		vOpts := append(opts, db.WithValidateField("object_name"))
		if err := fv(ctx, m.GetObjectName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["time_range"]; exists {

		vOpts := append(opts, db.WithValidateField("time_range"))
		if err := fv(ctx, m.GetTimeRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["title"]; exists {

		vOpts := append(opts, db.WithValidateField("title"))
		if err := fv(ctx, m.GetTitle(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDashboardLinkValidator = func() *ValidateDashboardLink {
	v := &ValidateDashboardLink{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhKey := v.KeyValidationRuleHandler
	rulesKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhKey(rulesKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DashboardLink.key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["key"] = vFn

	v.FldValidators["log_filters"] = LogFilterValidator().Validate

	return v
}()

func DashboardLinkValidator() db.Validator {
	return DefaultDashboardLinkValidator
}

// augmented methods on protoc/std generated struct

func (m *GenericLink) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GenericLink) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GenericLink) DeepCopy() *GenericLink {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GenericLink{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GenericLink) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GenericLink) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GenericLinkValidator().Validate(ctx, m, opts...)
}

type ValidateGenericLink struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGenericLink) KeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for key")
	}

	return validatorFn, nil
}

func (v *ValidateGenericLink) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GenericLink)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GenericLink got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["title"]; exists {

		vOpts := append(opts, db.WithValidateField("title"))
		if err := fv(ctx, m.GetTitle(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["url"]; exists {

		vOpts := append(opts, db.WithValidateField("url"))
		if err := fv(ctx, m.GetUrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGenericLinkValidator = func() *ValidateGenericLink {
	v := &ValidateGenericLink{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhKey := v.KeyValidationRuleHandler
	rulesKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhKey(rulesKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GenericLink.key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["key"] = vFn

	return v
}()

func GenericLinkValidator() db.Validator {
	return DefaultGenericLinkValidator
}

// augmented methods on protoc/std generated struct

func (m *Link) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Link) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Link) DeepCopy() *Link {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Link{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Link) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Link) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LinkValidator().Validate(ctx, m, opts...)
}

type ValidateLink struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLink) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Link)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Link got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetLinkType().(type) {
	case *Link_GenericLink:
		if fv, exists := v.FldValidators["link_type.generic_link"]; exists {
			val := m.GetLinkType().(*Link_GenericLink).GenericLink
			vOpts := append(opts,
				db.WithValidateField("link_type"),
				db.WithValidateField("generic_link"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *Link_DashboardLink:
		if fv, exists := v.FldValidators["link_type.dashboard_link"]; exists {
			val := m.GetLinkType().(*Link_DashboardLink).DashboardLink
			vOpts := append(opts,
				db.WithValidateField("link_type"),
				db.WithValidateField("dashboard_link"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLinkValidator = func() *ValidateLink {
	v := &ValidateLink{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["link_type.generic_link"] = GenericLinkValidator().Validate
	v.FldValidators["link_type.dashboard_link"] = DashboardLinkValidator().Validate

	return v
}()

func LinkValidator() db.Validator {
	return DefaultLinkValidator
}

// augmented methods on protoc/std generated struct

func (m *LogFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LogFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LogFilter) DeepCopy() *LogFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LogFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LogFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LogFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LogFilterValidator().Validate(ctx, m, opts...)
}

type ValidateLogFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLogFilter) ValuesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for values")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for values")
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
			return errors.Wrap(err, "repeated values")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items values")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateLogFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LogFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LogFilter got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["op"]; exists {

		vOpts := append(opts, db.WithValidateField("op"))
		if err := fv(ctx, m.GetOp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["values"]; exists {
		vOpts := append(opts, db.WithValidateField("values"))
		if err := fv(ctx, m.GetValues(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLogFilterValidator = func() *ValidateLogFilter {
	v := &ValidateLogFilter{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhValues := v.ValuesValidationRuleHandler
	rulesValues := map[string]string{
		"ves.io.schema.rules.repeated.unique": "true",
	}
	vFn, err = vrhValues(rulesValues)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for LogFilter.values: %s", err)
		panic(errMsg)
	}
	v.FldValidators["values"] = vFn

	return v
}()

func LogFilterValidator() db.Validator {
	return DefaultLogFilterValidator
}
