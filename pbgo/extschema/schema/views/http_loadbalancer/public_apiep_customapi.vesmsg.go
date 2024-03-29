// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package http_loadbalancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *APIGroupsApiep) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIGroupsApiep) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIGroupsApiep) DeepCopy() *APIGroupsApiep {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIGroupsApiep{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIGroupsApiep) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIGroupsApiep) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIGroupsApiepValidator().Validate(ctx, m, opts...)
}

type ValidateAPIGroupsApiep struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIGroupsApiep) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIGroupsApiep)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIGroupsApiep got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["category"]; exists {

		vOpts := append(opts, db.WithValidateField("category"))
		for idx, item := range m.GetCategory() {
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

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sensitive_data"]; exists {

		vOpts := append(opts, db.WithValidateField("sensitive_data"))
		for idx, item := range m.GetSensitiveData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["sensitive_data_types"]; exists {

		vOpts := append(opts, db.WithValidateField("sensitive_data_types"))
		for idx, item := range m.GetSensitiveDataTypes() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIGroupsApiepValidator = func() *ValidateAPIGroupsApiep {
	v := &ValidateAPIGroupsApiep{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIGroupsApiepValidator() db.Validator {
	return DefaultAPIGroupsApiepValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAPIEndpointsForGroupsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAPIEndpointsForGroupsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAPIEndpointsForGroupsReq) DeepCopy() *GetAPIEndpointsForGroupsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAPIEndpointsForGroupsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAPIEndpointsForGroupsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAPIEndpointsForGroupsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAPIEndpointsForGroupsReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetAPIEndpointsForGroupsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAPIEndpointsForGroupsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAPIEndpointsForGroupsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAPIEndpointsForGroupsReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAPIEndpointsForGroupsReqValidator = func() *ValidateGetAPIEndpointsForGroupsReq {
	v := &ValidateGetAPIEndpointsForGroupsReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAPIEndpointsForGroupsReqValidator() db.Validator {
	return DefaultGetAPIEndpointsForGroupsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAPIEndpointsForGroupsRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAPIEndpointsForGroupsRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAPIEndpointsForGroupsRsp) DeepCopy() *GetAPIEndpointsForGroupsRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAPIEndpointsForGroupsRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAPIEndpointsForGroupsRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAPIEndpointsForGroupsRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAPIEndpointsForGroupsRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetAPIEndpointsForGroupsRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAPIEndpointsForGroupsRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAPIEndpointsForGroupsRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAPIEndpointsForGroupsRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoints"]; exists {

		vOpts := append(opts, db.WithValidateField("api_endpoints"))
		for idx, item := range m.GetApiEndpoints() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["apieps_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("apieps_timestamp"))
		if err := fv(ctx, m.GetApiepsTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAPIEndpointsForGroupsRspValidator = func() *ValidateGetAPIEndpointsForGroupsRsp {
	v := &ValidateGetAPIEndpointsForGroupsRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAPIEndpointsForGroupsRspValidator() db.Validator {
	return DefaultGetAPIEndpointsForGroupsRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAPIEndpointsSchemaUpdatesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAPIEndpointsSchemaUpdatesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAPIEndpointsSchemaUpdatesReq) DeepCopy() *GetAPIEndpointsSchemaUpdatesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAPIEndpointsSchemaUpdatesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAPIEndpointsSchemaUpdatesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAPIEndpointsSchemaUpdatesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAPIEndpointsSchemaUpdatesReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetAPIEndpointsSchemaUpdatesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAPIEndpointsSchemaUpdatesReq) ApiEndpointsFilterValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for api_endpoints_filter")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.ApiOperation, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_views.ApiOperationValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for api_endpoints_filter")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.ApiOperation)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.ApiOperation, got %T", val)
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
			return errors.Wrap(err, "repeated api_endpoints_filter")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items api_endpoints_filter")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetAPIEndpointsSchemaUpdatesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAPIEndpointsSchemaUpdatesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAPIEndpointsSchemaUpdatesReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoints_filter"]; exists {
		vOpts := append(opts, db.WithValidateField("api_endpoints_filter"))
		if err := fv(ctx, m.GetApiEndpointsFilter(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_type"]; exists {

		vOpts := append(opts, db.WithValidateField("query_type"))
		if err := fv(ctx, m.GetQueryType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAPIEndpointsSchemaUpdatesReqValidator = func() *ValidateGetAPIEndpointsSchemaUpdatesReq {
	v := &ValidateGetAPIEndpointsSchemaUpdatesReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhApiEndpointsFilter := v.ApiEndpointsFilterValidationRuleHandler
	rulesApiEndpointsFilter := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "100",
	}
	vFn, err = vrhApiEndpointsFilter(rulesApiEndpointsFilter)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetAPIEndpointsSchemaUpdatesReq.api_endpoints_filter: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_endpoints_filter"] = vFn

	return v
}()

func GetAPIEndpointsSchemaUpdatesReqValidator() db.Validator {
	return DefaultGetAPIEndpointsSchemaUpdatesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAPIEndpointsSchemaUpdatesResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAPIEndpointsSchemaUpdatesResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAPIEndpointsSchemaUpdatesResp) DeepCopy() *GetAPIEndpointsSchemaUpdatesResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAPIEndpointsSchemaUpdatesResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAPIEndpointsSchemaUpdatesResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAPIEndpointsSchemaUpdatesResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAPIEndpointsSchemaUpdatesRespValidator().Validate(ctx, m, opts...)
}

type ValidateGetAPIEndpointsSchemaUpdatesResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAPIEndpointsSchemaUpdatesResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAPIEndpointsSchemaUpdatesResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAPIEndpointsSchemaUpdatesResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoints_current_schemas"]; exists {

		vOpts := append(opts, db.WithValidateField("api_endpoints_current_schemas"))
		for idx, item := range m.GetApiEndpointsCurrentSchemas() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["api_endpoints_updated_schemas"]; exists {

		vOpts := append(opts, db.WithValidateField("api_endpoints_updated_schemas"))
		for idx, item := range m.GetApiEndpointsUpdatedSchemas() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAPIEndpointsSchemaUpdatesRespValidator = func() *ValidateGetAPIEndpointsSchemaUpdatesResp {
	v := &ValidateGetAPIEndpointsSchemaUpdatesResp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_endpoints_current_schemas"] = ves_io_schema_views.ApiEndpointWithSchemaValidator().Validate

	v.FldValidators["api_endpoints_updated_schemas"] = ves_io_schema_views.ApiEndpointWithSchemaValidator().Validate

	return v
}()

func GetAPIEndpointsSchemaUpdatesRespValidator() db.Validator {
	return DefaultGetAPIEndpointsSchemaUpdatesRespValidator
}

// augmented methods on protoc/std generated struct

func (m *SwaggerSpecReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SwaggerSpecReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SwaggerSpecReq) DeepCopy() *SwaggerSpecReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SwaggerSpecReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SwaggerSpecReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SwaggerSpecReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SwaggerSpecReqValidator().Validate(ctx, m, opts...)
}

type ValidateSwaggerSpecReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSwaggerSpecReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SwaggerSpecReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SwaggerSpecReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSwaggerSpecReqValidator = func() *ValidateSwaggerSpecReq {
	v := &ValidateSwaggerSpecReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SwaggerSpecReqValidator() db.Validator {
	return DefaultSwaggerSpecReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SwaggerSpecRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SwaggerSpecRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SwaggerSpecRsp) DeepCopy() *SwaggerSpecRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SwaggerSpecRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SwaggerSpecRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SwaggerSpecRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SwaggerSpecRspValidator().Validate(ctx, m, opts...)
}

type ValidateSwaggerSpecRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSwaggerSpecRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SwaggerSpecRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SwaggerSpecRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["swagger_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("swagger_spec"))
		if err := fv(ctx, m.GetSwaggerSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSwaggerSpecRspValidator = func() *ValidateSwaggerSpecRsp {
	v := &ValidateSwaggerSpecRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SwaggerSpecRspValidator() db.Validator {
	return DefaultSwaggerSpecRspValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateAPIEndpointsSchemasReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateAPIEndpointsSchemasReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateAPIEndpointsSchemasReq) DeepCopy() *UpdateAPIEndpointsSchemasReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateAPIEndpointsSchemasReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateAPIEndpointsSchemasReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateAPIEndpointsSchemasReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateAPIEndpointsSchemasReqValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateAPIEndpointsSchemasReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateAPIEndpointsSchemasReq) ApiEndpointsSchemaUpdatesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for api_endpoints_schema_updates")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.ApiEndpointWithSchema, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_views.ApiEndpointWithSchemaValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for api_endpoints_schema_updates")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.ApiEndpointWithSchema)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.ApiEndpointWithSchema, got %T", val)
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
			return errors.Wrap(err, "repeated api_endpoints_schema_updates")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items api_endpoints_schema_updates")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateUpdateAPIEndpointsSchemasReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateAPIEndpointsSchemasReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateAPIEndpointsSchemasReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoints_schema_updates"]; exists {
		vOpts := append(opts, db.WithValidateField("api_endpoints_schema_updates"))
		if err := fv(ctx, m.GetApiEndpointsSchemaUpdates(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateAPIEndpointsSchemasReqValidator = func() *ValidateUpdateAPIEndpointsSchemasReq {
	v := &ValidateUpdateAPIEndpointsSchemasReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhApiEndpointsSchemaUpdates := v.ApiEndpointsSchemaUpdatesValidationRuleHandler
	rulesApiEndpointsSchemaUpdates := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "100",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhApiEndpointsSchemaUpdates(rulesApiEndpointsSchemaUpdates)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpdateAPIEndpointsSchemasReq.api_endpoints_schema_updates: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_endpoints_schema_updates"] = vFn

	return v
}()

func UpdateAPIEndpointsSchemasReqValidator() db.Validator {
	return DefaultUpdateAPIEndpointsSchemasReqValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateAPIEndpointsSchemasResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateAPIEndpointsSchemasResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateAPIEndpointsSchemasResp) DeepCopy() *UpdateAPIEndpointsSchemasResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateAPIEndpointsSchemasResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateAPIEndpointsSchemasResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateAPIEndpointsSchemasResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateAPIEndpointsSchemasRespValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateAPIEndpointsSchemasResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateAPIEndpointsSchemasResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateAPIEndpointsSchemasResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateAPIEndpointsSchemasResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["updated_api_endpoints"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_api_endpoints"))
		for idx, item := range m.GetUpdatedApiEndpoints() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateAPIEndpointsSchemasRespValidator = func() *ValidateUpdateAPIEndpointsSchemasResp {
	v := &ValidateUpdateAPIEndpointsSchemasResp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["updated_api_endpoints"] = ves_io_schema_views.ApiOperationValidator().Validate

	return v
}()

func UpdateAPIEndpointsSchemasRespValidator() db.Validator {
	return DefaultUpdateAPIEndpointsSchemasRespValidator
}
