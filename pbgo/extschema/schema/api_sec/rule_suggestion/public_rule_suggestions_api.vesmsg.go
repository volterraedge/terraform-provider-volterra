// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package rule_suggestion

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_views_common_waf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_waf"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetSuggestedAPIEndpointProtectionRuleReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedAPIEndpointProtectionRuleReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedAPIEndpointProtectionRuleReq) DeepCopy() *GetSuggestedAPIEndpointProtectionRuleReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedAPIEndpointProtectionRuleReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedAPIEndpointProtectionRuleReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedAPIEndpointProtectionRuleReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedAPIEndpointProtectionRuleReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedAPIEndpointProtectionRuleReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleReq) VirtualHostNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for virtual_host_name")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleReq) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleReq) MethodValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for method")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedAPIEndpointProtectionRuleReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedAPIEndpointProtectionRuleReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host_name"))
		if err := fv(ctx, m.GetVirtualHostName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedAPIEndpointProtectionRuleReqValidator = func() *ValidateGetSuggestedAPIEndpointProtectionRuleReq {
	v := &ValidateGetSuggestedAPIEndpointProtectionRuleReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedAPIEndpointProtectionRuleReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhVirtualHostName := v.VirtualHostNameValidationRuleHandler
	rulesVirtualHostName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVirtualHostName(rulesVirtualHostName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedAPIEndpointProtectionRuleReq.virtual_host_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["virtual_host_name"] = vFn

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.message.required":           "true",
		"ves.io.schema.rules.string.max_len":             "1024",
		"ves.io.schema.rules.string.templated_http_path": "true",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedAPIEndpointProtectionRuleReq.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	vrhMethod := v.MethodValidationRuleHandler
	rulesMethod := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhMethod(rulesMethod)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedAPIEndpointProtectionRuleReq.method: %s", err)
		panic(errMsg)
	}
	v.FldValidators["method"] = vFn

	return v
}()

func GetSuggestedAPIEndpointProtectionRuleReqValidator() db.Validator {
	return DefaultGetSuggestedAPIEndpointProtectionRuleReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) DeepCopy() *GetSuggestedAPIEndpointProtectionRuleRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedAPIEndpointProtectionRuleRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedAPIEndpointProtectionRuleRspValidator().Validate(ctx, m, opts...)
}

func (m *GetSuggestedAPIEndpointProtectionRuleRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetRuleDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSuggestedAPIEndpointProtectionRuleRsp) GetRuleDRefInfo() ([]db.DRefInfo, error) {
	if m.GetRule() == nil {
		return nil, nil
	}

	drInfos, err := m.GetRule().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetRule().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "rule." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetSuggestedAPIEndpointProtectionRuleRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedAPIEndpointProtectionRuleRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedAPIEndpointProtectionRuleRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedAPIEndpointProtectionRuleRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["found_existing_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("found_existing_rule"))
		if err := fv(ctx, m.GetFoundExistingRule(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["loadbalancer_type"]; exists {

		vOpts := append(opts, db.WithValidateField("loadbalancer_type"))
		if err := fv(ctx, m.GetLoadbalancerType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule"]; exists {

		vOpts := append(opts, db.WithValidateField("rule"))
		if err := fv(ctx, m.GetRule(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedAPIEndpointProtectionRuleRspValidator = func() *ValidateGetSuggestedAPIEndpointProtectionRuleRsp {
	v := &ValidateGetSuggestedAPIEndpointProtectionRuleRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["rule"] = ves_io_schema_views_common_waf.APIEndpointProtectionRuleValidator().Validate

	return v
}()

func GetSuggestedAPIEndpointProtectionRuleRspValidator() db.Validator {
	return DefaultGetSuggestedAPIEndpointProtectionRuleRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedOasValidationRuleReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedOasValidationRuleReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedOasValidationRuleReq) DeepCopy() *GetSuggestedOasValidationRuleReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedOasValidationRuleReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedOasValidationRuleReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedOasValidationRuleReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedOasValidationRuleReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedOasValidationRuleReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedOasValidationRuleReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedOasValidationRuleReq) VirtualHostNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for virtual_host_name")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedOasValidationRuleReq) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedOasValidationRuleReq) MethodValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for method")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedOasValidationRuleReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedOasValidationRuleReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedOasValidationRuleReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups"))
		for idx, item := range m.GetApiGroups() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host_name"))
		if err := fv(ctx, m.GetVirtualHostName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedOasValidationRuleReqValidator = func() *ValidateGetSuggestedOasValidationRuleReq {
	v := &ValidateGetSuggestedOasValidationRuleReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedOasValidationRuleReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhVirtualHostName := v.VirtualHostNameValidationRuleHandler
	rulesVirtualHostName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVirtualHostName(rulesVirtualHostName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedOasValidationRuleReq.virtual_host_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["virtual_host_name"] = vFn

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.message.required":           "true",
		"ves.io.schema.rules.string.max_len":             "1024",
		"ves.io.schema.rules.string.templated_http_path": "true",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedOasValidationRuleReq.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	vrhMethod := v.MethodValidationRuleHandler
	rulesMethod := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhMethod(rulesMethod)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedOasValidationRuleReq.method: %s", err)
		panic(errMsg)
	}
	v.FldValidators["method"] = vFn

	return v
}()

func GetSuggestedOasValidationRuleReqValidator() db.Validator {
	return DefaultGetSuggestedOasValidationRuleReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedOasValidationRuleRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedOasValidationRuleRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedOasValidationRuleRsp) DeepCopy() *GetSuggestedOasValidationRuleRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedOasValidationRuleRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedOasValidationRuleRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedOasValidationRuleRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedOasValidationRuleRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedOasValidationRuleRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedOasValidationRuleRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedOasValidationRuleRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedOasValidationRuleRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetApiValidationChoice().(type) {
	case *GetSuggestedOasValidationRuleRsp_CustomOasValidation:
		if fv, exists := v.FldValidators["api_validation_choice.custom_oas_validation"]; exists {
			val := m.GetApiValidationChoice().(*GetSuggestedOasValidationRuleRsp_CustomOasValidation).CustomOasValidation
			vOpts := append(opts,
				db.WithValidateField("api_validation_choice"),
				db.WithValidateField("custom_oas_validation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GetSuggestedOasValidationRuleRsp_AllEndpointsOasValidation:
		if fv, exists := v.FldValidators["api_validation_choice.all_endpoints_oas_validation"]; exists {
			val := m.GetApiValidationChoice().(*GetSuggestedOasValidationRuleRsp_AllEndpointsOasValidation).AllEndpointsOasValidation
			vOpts := append(opts,
				db.WithValidateField("api_validation_choice"),
				db.WithValidateField("all_endpoints_oas_validation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["found_existing_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("found_existing_rule"))
		if err := fv(ctx, m.GetFoundExistingRule(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["loadbalancer_type"]; exists {

		vOpts := append(opts, db.WithValidateField("loadbalancer_type"))
		if err := fv(ctx, m.GetLoadbalancerType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedOasValidationRuleRspValidator = func() *ValidateGetSuggestedOasValidationRuleRsp {
	v := &ValidateGetSuggestedOasValidationRuleRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_validation_choice.custom_oas_validation"] = ves_io_schema_views_common_waf.OpenApiValidationRuleValidator().Validate
	v.FldValidators["api_validation_choice.all_endpoints_oas_validation"] = ves_io_schema_views_common_waf.OpenApiValidationAllSpecEndpointsSettingsValidator().Validate

	return v
}()

func GetSuggestedOasValidationRuleRspValidator() db.Validator {
	return DefaultGetSuggestedOasValidationRuleRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedRateLimitRuleReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedRateLimitRuleReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedRateLimitRuleReq) DeepCopy() *GetSuggestedRateLimitRuleReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedRateLimitRuleReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedRateLimitRuleReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedRateLimitRuleReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedRateLimitRuleReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedRateLimitRuleReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedRateLimitRuleReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedRateLimitRuleReq) VirtualHostNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for virtual_host_name")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedRateLimitRuleReq) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedRateLimitRuleReq) MethodValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for method")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedRateLimitRuleReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedRateLimitRuleReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedRateLimitRuleReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host_name"))
		if err := fv(ctx, m.GetVirtualHostName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedRateLimitRuleReqValidator = func() *ValidateGetSuggestedRateLimitRuleReq {
	v := &ValidateGetSuggestedRateLimitRuleReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedRateLimitRuleReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhVirtualHostName := v.VirtualHostNameValidationRuleHandler
	rulesVirtualHostName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVirtualHostName(rulesVirtualHostName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedRateLimitRuleReq.virtual_host_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["virtual_host_name"] = vFn

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.message.required":           "true",
		"ves.io.schema.rules.string.max_len":             "1024",
		"ves.io.schema.rules.string.templated_http_path": "true",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedRateLimitRuleReq.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	vrhMethod := v.MethodValidationRuleHandler
	rulesMethod := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhMethod(rulesMethod)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedRateLimitRuleReq.method: %s", err)
		panic(errMsg)
	}
	v.FldValidators["method"] = vFn

	return v
}()

func GetSuggestedRateLimitRuleReqValidator() db.Validator {
	return DefaultGetSuggestedRateLimitRuleReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedRateLimitRuleRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedRateLimitRuleRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedRateLimitRuleRsp) DeepCopy() *GetSuggestedRateLimitRuleRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedRateLimitRuleRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedRateLimitRuleRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedRateLimitRuleRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedRateLimitRuleRspValidator().Validate(ctx, m, opts...)
}

func (m *GetSuggestedRateLimitRuleRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetRuleDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSuggestedRateLimitRuleRsp) GetRuleDRefInfo() ([]db.DRefInfo, error) {
	if m.GetRule() == nil {
		return nil, nil
	}

	drInfos, err := m.GetRule().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetRule().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "rule." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetSuggestedRateLimitRuleRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedRateLimitRuleRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedRateLimitRuleRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedRateLimitRuleRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["found_existing_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("found_existing_rule"))
		if err := fv(ctx, m.GetFoundExistingRule(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["loadbalancer_type"]; exists {

		vOpts := append(opts, db.WithValidateField("loadbalancer_type"))
		if err := fv(ctx, m.GetLoadbalancerType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule"]; exists {

		vOpts := append(opts, db.WithValidateField("rule"))
		if err := fv(ctx, m.GetRule(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedRateLimitRuleRspValidator = func() *ValidateGetSuggestedRateLimitRuleRsp {
	v := &ValidateGetSuggestedRateLimitRuleRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["rule"] = ves_io_schema_views_common_waf.ApiEndpointRuleValidator().Validate

	return v
}()

func GetSuggestedRateLimitRuleRspValidator() db.Validator {
	return DefaultGetSuggestedRateLimitRuleRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedSensitiveDataRuleReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedSensitiveDataRuleReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedSensitiveDataRuleReq) DeepCopy() *GetSuggestedSensitiveDataRuleReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedSensitiveDataRuleReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedSensitiveDataRuleReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedSensitiveDataRuleReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedSensitiveDataRuleReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedSensitiveDataRuleReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedSensitiveDataRuleReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedSensitiveDataRuleReq) VirtualHostNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for virtual_host_name")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedSensitiveDataRuleReq) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedSensitiveDataRuleReq) MethodValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for method")
	}

	return validatorFn, nil
}

func (v *ValidateGetSuggestedSensitiveDataRuleReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedSensitiveDataRuleReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedSensitiveDataRuleReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host_name"))
		if err := fv(ctx, m.GetVirtualHostName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedSensitiveDataRuleReqValidator = func() *ValidateGetSuggestedSensitiveDataRuleReq {
	v := &ValidateGetSuggestedSensitiveDataRuleReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedSensitiveDataRuleReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhVirtualHostName := v.VirtualHostNameValidationRuleHandler
	rulesVirtualHostName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVirtualHostName(rulesVirtualHostName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedSensitiveDataRuleReq.virtual_host_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["virtual_host_name"] = vFn

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.message.required":           "true",
		"ves.io.schema.rules.string.max_len":             "1024",
		"ves.io.schema.rules.string.templated_http_path": "true",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedSensitiveDataRuleReq.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	vrhMethod := v.MethodValidationRuleHandler
	rulesMethod := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhMethod(rulesMethod)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSuggestedSensitiveDataRuleReq.method: %s", err)
		panic(errMsg)
	}
	v.FldValidators["method"] = vFn

	return v
}()

func GetSuggestedSensitiveDataRuleReqValidator() db.Validator {
	return DefaultGetSuggestedSensitiveDataRuleReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSuggestedSensitiveDataRuleRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSuggestedSensitiveDataRuleRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSuggestedSensitiveDataRuleRsp) DeepCopy() *GetSuggestedSensitiveDataRuleRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSuggestedSensitiveDataRuleRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSuggestedSensitiveDataRuleRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSuggestedSensitiveDataRuleRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSuggestedSensitiveDataRuleRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetSuggestedSensitiveDataRuleRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSuggestedSensitiveDataRuleRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSuggestedSensitiveDataRuleRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSuggestedSensitiveDataRuleRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["found_existing_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("found_existing_rule"))
		if err := fv(ctx, m.GetFoundExistingRule(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["loadbalancer_type"]; exists {

		vOpts := append(opts, db.WithValidateField("loadbalancer_type"))
		if err := fv(ctx, m.GetLoadbalancerType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule"]; exists {

		vOpts := append(opts, db.WithValidateField("rule"))
		if err := fv(ctx, m.GetRule(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSuggestedSensitiveDataRuleRspValidator = func() *ValidateGetSuggestedSensitiveDataRuleRsp {
	v := &ValidateGetSuggestedSensitiveDataRuleRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["rule"] = ves_io_schema_views_http_loadbalancer.SensitiveDataTypesValidator().Validate

	return v
}()

func GetSuggestedSensitiveDataRuleRspValidator() db.Validator {
	return DefaultGetSuggestedSensitiveDataRuleRspValidator
}