//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package namespace

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

func (m *CascadeDeleteItemType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CascadeDeleteItemType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CascadeDeleteItemType) DeepCopy() *CascadeDeleteItemType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CascadeDeleteItemType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CascadeDeleteItemType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CascadeDeleteItemType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CascadeDeleteItemTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCascadeDeleteItemType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCascadeDeleteItemType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CascadeDeleteItemType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CascadeDeleteItemType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["error_message"]; exists {

		vOpts := append(opts, db.WithValidateField("error_message"))
		if err := fv(ctx, m.GetErrorMessage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_name"]; exists {

		vOpts := append(opts, db.WithValidateField("object_name"))
		if err := fv(ctx, m.GetObjectName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_type"]; exists {

		vOpts := append(opts, db.WithValidateField("object_type"))
		if err := fv(ctx, m.GetObjectType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCascadeDeleteItemTypeValidator = func() *ValidateCascadeDeleteItemType {
	v := &ValidateCascadeDeleteItemType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CascadeDeleteItemTypeValidator() db.Validator {
	return DefaultCascadeDeleteItemTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *CascadeDeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CascadeDeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CascadeDeleteRequest) DeepCopy() *CascadeDeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CascadeDeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CascadeDeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CascadeDeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CascadeDeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCascadeDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCascadeDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CascadeDeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CascadeDeleteRequest got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCascadeDeleteRequestValidator = func() *ValidateCascadeDeleteRequest {
	v := &ValidateCascadeDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CascadeDeleteRequestValidator() db.Validator {
	return DefaultCascadeDeleteRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CascadeDeleteResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CascadeDeleteResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CascadeDeleteResponse) DeepCopy() *CascadeDeleteResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CascadeDeleteResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CascadeDeleteResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CascadeDeleteResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CascadeDeleteResponseValidator().Validate(ctx, m, opts...)
}

type ValidateCascadeDeleteResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCascadeDeleteResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CascadeDeleteResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CascadeDeleteResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCascadeDeleteResponseValidator = func() *ValidateCascadeDeleteResponse {
	v := &ValidateCascadeDeleteResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CascadeDeleteResponseValidator() db.Validator {
	return DefaultCascadeDeleteResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

// Redact squashes sensitive info in m (in-place)
func (m *GetSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	for idx, e := range m.GetProxySubCas() {
		if err := e.Redact(ctx); err != nil {
			return errors.Wrapf(err, "Redacting GlobalSpecType.proxy_sub_cas idx %v", idx)
		}
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

	if fv, exists := v.FldValidators["allow_advertise_on_public"]; exists {

		vOpts := append(opts, db.WithValidateField("allow_advertise_on_public"))
		if err := fv(ctx, m.GetAllowAdvertiseOnPublic(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_sub_ca_latest_version"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_sub_ca_latest_version"))
		if err := fv(ctx, m.GetProxySubCaLatestVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_sub_cas"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_sub_cas"))
		for idx, item := range m.GetProxySubCas() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["proxy_sub_cas"] = SubCAValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *ReplaceSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *SubCA) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SubCA) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SubCA) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.PrivateKey = nil

	return copy.string()
}

func (m *SubCA) GoString() string {
	copy := m.DeepCopy()
	copy.PrivateKey = nil

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *SubCA) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.PrivateKey = nil

	return nil
}

func (m *SubCA) DeepCopy() *SubCA {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SubCA{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SubCA) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SubCA) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SubCAValidator().Validate(ctx, m, opts...)
}

type ValidateSubCA struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSubCA) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateSubCA) PemValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for pem")
	}

	return validatorFn, nil
}

func (v *ValidateSubCA) PrivateKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for private_key")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSubCA) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SubCA)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SubCA got type %s", t)
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

	if fv, exists := v.FldValidators["pem"]; exists {

		vOpts := append(opts, db.WithValidateField("pem"))
		if err := fv(ctx, m.GetPem(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_key"]; exists {

		vOpts := append(opts, db.WithValidateField("private_key"))
		if err := fv(ctx, m.GetPrivateKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["version"]; exists {

		vOpts := append(opts, db.WithValidateField("version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSubCAValidator = func() *ValidateSubCA {
	v := &ValidateSubCA{FldValidators: map[string]db.ValidatorFunc{}}

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
		"ves.io.schema.rules.string.max_len":   "514",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SubCA.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhPem := v.PemValidationRuleHandler
	rulesPem := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "10240",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhPem(rulesPem)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SubCA.pem: %s", err)
		panic(errMsg)
	}
	v.FldValidators["pem"] = vFn

	vrhPrivateKey := v.PrivateKeyValidationRuleHandler
	rulesPrivateKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPrivateKey(rulesPrivateKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SubCA.private_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["private_key"] = vFn

	return v
}()

func SubCAValidator() db.Validator {
	return DefaultSubCAValidator
}

// augmented methods on protoc/std generated struct

func (m *SuggestValuesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuggestValuesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuggestValuesReq) DeepCopy() *SuggestValuesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuggestValuesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuggestValuesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuggestValuesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuggestValuesReqValidator().Validate(ctx, m, opts...)
}

type ValidateSuggestValuesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuggestValuesReq) FieldPathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for field_path")
	}

	return validatorFn, nil
}

func (v *ValidateSuggestValuesReq) MatchValueValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for match_value")
	}

	return validatorFn, nil
}

func (v *ValidateSuggestValuesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuggestValuesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuggestValuesReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field_path"]; exists {

		vOpts := append(opts, db.WithValidateField("field_path"))
		if err := fv(ctx, m.GetFieldPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["match_value"]; exists {

		vOpts := append(opts, db.WithValidateField("match_value"))
		if err := fv(ctx, m.GetMatchValue(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["request_body"]; exists {

		vOpts := append(opts, db.WithValidateField("request_body"))
		if err := fv(ctx, m.GetRequestBody(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuggestValuesReqValidator = func() *ValidateSuggestValuesReq {
	v := &ValidateSuggestValuesReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhFieldPath := v.FieldPathValidationRuleHandler
	rulesFieldPath := map[string]string{
		"ves.io.schema.rules.string.max_len": "1024",
	}
	vFn, err = vrhFieldPath(rulesFieldPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SuggestValuesReq.field_path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field_path"] = vFn

	vrhMatchValue := v.MatchValueValidationRuleHandler
	rulesMatchValue := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhMatchValue(rulesMatchValue)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SuggestValuesReq.match_value: %s", err)
		panic(errMsg)
	}
	v.FldValidators["match_value"] = vFn

	return v
}()

func SuggestValuesReqValidator() db.Validator {
	return DefaultSuggestValuesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SuggestValuesResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuggestValuesResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuggestValuesResp) DeepCopy() *SuggestValuesResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuggestValuesResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuggestValuesResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuggestValuesResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuggestValuesRespValidator().Validate(ctx, m, opts...)
}

type ValidateSuggestValuesResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuggestValuesResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuggestValuesResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuggestValuesResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuggestValuesRespValidator = func() *ValidateSuggestValuesResp {
	v := &ValidateSuggestValuesResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuggestValuesRespValidator() db.Validator {
	return DefaultSuggestValuesRespValidator
}

// augmented methods on protoc/std generated struct

func (m *SuggestedItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuggestedItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuggestedItem) DeepCopy() *SuggestedItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuggestedItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuggestedItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuggestedItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuggestedItemValidator().Validate(ctx, m, opts...)
}

type ValidateSuggestedItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuggestedItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuggestedItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuggestedItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuggestedItemValidator = func() *ValidateSuggestedItem {
	v := &ValidateSuggestedItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuggestedItemValidator() db.Validator {
	return DefaultSuggestedItemValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
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

}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
