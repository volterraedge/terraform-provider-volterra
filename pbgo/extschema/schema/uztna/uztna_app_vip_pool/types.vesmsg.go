// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_app_vip_pool

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) IpVersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip_version")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ip_version"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_version"))
		if err := fv(ctx, m.GetIpVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhIpVersion := v.IpVersionValidationRuleHandler
	rulesIpVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIpVersion(rulesIpVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.ip_version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_version"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
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

func (v *ValidateGetSpecType) IpVersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip_version")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["ip_version"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_version"))
		if err := fv(ctx, m.GetIpVersion(), vOpts...); err != nil {
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

	vrhIpVersion := v.IpVersionValidationRuleHandler
	rulesIpVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIpVersion(rulesIpVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.ip_version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_version"] = vFn

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

func (v *ValidateGlobalSpecType) IpVersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip_version")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["ip_version"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_version"))
		if err := fv(ctx, m.GetIpVersion(), vOpts...); err != nil {
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

	vrhIpVersion := v.IpVersionValidationRuleHandler
	rulesIpVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIpVersion(rulesIpVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.ip_version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_version"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *IPV4ApplicationVIP) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IPV4ApplicationVIP) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IPV4ApplicationVIP) DeepCopy() *IPV4ApplicationVIP {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IPV4ApplicationVIP{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IPV4ApplicationVIP) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IPV4ApplicationVIP) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IPV4ApplicationVIPValidator().Validate(ctx, m, opts...)
}

type ValidateIPV4ApplicationVIP struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIPV4ApplicationVIP) PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for prefix")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for prefix")
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
			return errors.Wrap(err, "repeated prefix")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items prefix")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateIPV4ApplicationVIP) Vip4RangeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for vip4_range")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*VIP4PoolRange, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := VIP4PoolRangeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vip4_range")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*VIP4PoolRange)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*VIP4PoolRange, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated vip4_range")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vip4_range")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateIPV4ApplicationVIP) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IPV4ApplicationVIP)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IPV4ApplicationVIP got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["prefix"]; exists {
		vOpts := append(opts, db.WithValidateField("prefix"))
		if err := fv(ctx, m.GetPrefix(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vip4_range"]; exists {
		vOpts := append(opts, db.WithValidateField("vip4_range"))
		if err := fv(ctx, m.GetVip4Range(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIPV4ApplicationVIPValidator = func() *ValidateIPV4ApplicationVIP {
	v := &ValidateIPV4ApplicationVIP{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrefix := v.PrefixValidationRuleHandler
	rulesPrefix := map[string]string{
		"ves.io.schema.rules.message.required":                  "true",
		"ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
		"ves.io.schema.rules.repeated.items.string.not_empty":   "true",
		"ves.io.schema.rules.repeated.max_items":                "10",
		"ves.io.schema.rules.repeated.min_items":                "1",
		"ves.io.schema.rules.repeated.unique":                   "true",
	}
	vFn, err = vrhPrefix(rulesPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for IPV4ApplicationVIP.prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["prefix"] = vFn

	vrhVip4Range := v.Vip4RangeValidationRuleHandler
	rulesVip4Range := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "10",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhVip4Range(rulesVip4Range)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for IPV4ApplicationVIP.vip4_range: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vip4_range"] = vFn

	return v
}()

func IPV4ApplicationVIPValidator() db.Validator {
	return DefaultIPV4ApplicationVIPValidator
}

// augmented methods on protoc/std generated struct

func (m *IPV6ApplicationVIP) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IPV6ApplicationVIP) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IPV6ApplicationVIP) DeepCopy() *IPV6ApplicationVIP {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IPV6ApplicationVIP{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IPV6ApplicationVIP) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IPV6ApplicationVIP) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IPV6ApplicationVIPValidator().Validate(ctx, m, opts...)
}

type ValidateIPV6ApplicationVIP struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIPV6ApplicationVIP) Ipv6PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for ipv6_prefix")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for ipv6_prefix")
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
			return errors.Wrap(err, "repeated ipv6_prefix")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items ipv6_prefix")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateIPV6ApplicationVIP) Vip6RangeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for vip6_range")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*VIP6PoolRange, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := VIP6PoolRangeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vip6_range")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*VIP6PoolRange)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*VIP6PoolRange, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated vip6_range")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vip6_range")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateIPV6ApplicationVIP) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IPV6ApplicationVIP)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IPV6ApplicationVIP got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ipv6_prefix"]; exists {
		vOpts := append(opts, db.WithValidateField("ipv6_prefix"))
		if err := fv(ctx, m.GetIpv6Prefix(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vip6_range"]; exists {
		vOpts := append(opts, db.WithValidateField("vip6_range"))
		if err := fv(ctx, m.GetVip6Range(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIPV6ApplicationVIPValidator = func() *ValidateIPV6ApplicationVIP {
	v := &ValidateIPV6ApplicationVIP{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhIpv6Prefix := v.Ipv6PrefixValidationRuleHandler
	rulesIpv6Prefix := map[string]string{
		"ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
		"ves.io.schema.rules.repeated.items.string.not_empty":   "true",
		"ves.io.schema.rules.repeated.max_items":                "10",
		"ves.io.schema.rules.repeated.unique":                   "true",
	}
	vFn, err = vrhIpv6Prefix(rulesIpv6Prefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for IPV6ApplicationVIP.ipv6_prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ipv6_prefix"] = vFn

	vrhVip6Range := v.Vip6RangeValidationRuleHandler
	rulesVip6Range := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "10",
		"ves.io.schema.rules.repeated.min_items": "1",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhVip6Range(rulesVip6Range)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for IPV6ApplicationVIP.vip6_range: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vip6_range"] = vFn

	return v
}()

func IPV6ApplicationVIPValidator() db.Validator {
	return DefaultIPV6ApplicationVIPValidator
}

// augmented methods on protoc/std generated struct

func (m *IPVersion) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IPVersion) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IPVersion) DeepCopy() *IPVersion {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IPVersion{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IPVersion) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IPVersion) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IPVersionValidator().Validate(ctx, m, opts...)
}

type ValidateIPVersion struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIPVersion) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IPVersion)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IPVersion got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetIpVip().(type) {
	case *IPVersion_Ipv4Vip:
		if fv, exists := v.FldValidators["ip_vip.ipv4_vip"]; exists {
			val := m.GetIpVip().(*IPVersion_Ipv4Vip).Ipv4Vip
			vOpts := append(opts,
				db.WithValidateField("ip_vip"),
				db.WithValidateField("ipv4_vip"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *IPVersion_Ipv6Vip:
		if fv, exists := v.FldValidators["ip_vip.ipv6_vip"]; exists {
			val := m.GetIpVip().(*IPVersion_Ipv6Vip).Ipv6Vip
			vOpts := append(opts,
				db.WithValidateField("ip_vip"),
				db.WithValidateField("ipv6_vip"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIPVersionValidator = func() *ValidateIPVersion {
	v := &ValidateIPVersion{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["ip_vip.ipv4_vip"] = IPV4ApplicationVIPValidator().Validate
	v.FldValidators["ip_vip.ipv6_vip"] = IPV6ApplicationVIPValidator().Validate

	return v
}()

func IPVersionValidator() db.Validator {
	return DefaultIPVersionValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) IpVersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip_version")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ip_version"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_version"))
		if err := fv(ctx, m.GetIpVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhIpVersion := v.IpVersionValidationRuleHandler
	rulesIpVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIpVersion(rulesIpVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.ip_version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_version"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *VIP4PoolRange) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VIP4PoolRange) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VIP4PoolRange) DeepCopy() *VIP4PoolRange {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VIP4PoolRange{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VIP4PoolRange) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VIP4PoolRange) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VIP4PoolRangeValidator().Validate(ctx, m, opts...)
}

type ValidateVIP4PoolRange struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVIP4PoolRange) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VIP4PoolRange)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VIP4PoolRange got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_address"]; exists {

		vOpts := append(opts, db.WithValidateField("end_address"))
		if err := fv(ctx, m.GetEndAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_address"]; exists {

		vOpts := append(opts, db.WithValidateField("start_address"))
		if err := fv(ctx, m.GetStartAddress(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVIP4PoolRangeValidator = func() *ValidateVIP4PoolRange {
	v := &ValidateVIP4PoolRange{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["start_address"] = ves_io_schema.Ipv4AddressTypeValidator().Validate

	v.FldValidators["end_address"] = ves_io_schema.Ipv4AddressTypeValidator().Validate

	return v
}()

func VIP4PoolRangeValidator() db.Validator {
	return DefaultVIP4PoolRangeValidator
}

// augmented methods on protoc/std generated struct

func (m *VIP6PoolRange) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VIP6PoolRange) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VIP6PoolRange) DeepCopy() *VIP6PoolRange {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VIP6PoolRange{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VIP6PoolRange) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VIP6PoolRange) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VIP6PoolRangeValidator().Validate(ctx, m, opts...)
}

type ValidateVIP6PoolRange struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVIP6PoolRange) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VIP6PoolRange)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VIP6PoolRange got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_address"]; exists {

		vOpts := append(opts, db.WithValidateField("end_address"))
		if err := fv(ctx, m.GetEndAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_address"]; exists {

		vOpts := append(opts, db.WithValidateField("start_address"))
		if err := fv(ctx, m.GetStartAddress(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVIP6PoolRangeValidator = func() *ValidateVIP6PoolRange {
	v := &ValidateVIP6PoolRange{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["start_address"] = ves_io_schema.Ipv6AddressTypeValidator().Validate

	v.FldValidators["end_address"] = ves_io_schema.Ipv6AddressTypeValidator().Validate

	return v
}()

func VIP6PoolRangeValidator() db.Validator {
	return DefaultVIP6PoolRangeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.IpVersion = f.GetIpVersion()
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *CreateSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *CreateSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.IpVersion = m1.IpVersion
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.IpVersion = f.GetIpVersion()
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

	f.IpVersion = m1.IpVersion
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.IpVersion = f.GetIpVersion()
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.IpVersion = m1.IpVersion
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}