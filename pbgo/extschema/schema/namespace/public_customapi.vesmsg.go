// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package namespace

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

func (m *APIItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIItem) DeepCopy() *APIItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIItemValidator().Validate(ctx, m, opts...)
}

type ValidateAPIItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIItem got type %s", t)
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

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["result"]; exists {

		vOpts := append(opts, db.WithValidateField("result"))
		if err := fv(ctx, m.GetResult(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIItemValidator = func() *ValidateAPIItem {
	v := &ValidateAPIItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIItemValidator() db.Validator {
	return DefaultAPIItemValidator
}

// augmented methods on protoc/std generated struct

func (m *APIItemList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIItemList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIItemList) DeepCopy() *APIItemList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIItemList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIItemList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIItemList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIItemListValidator().Validate(ctx, m, opts...)
}

type ValidateAPIItemList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIItemList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIItemList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIItemList got type %s", t)
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

	if fv, exists := v.FldValidators["list_id"]; exists {

		vOpts := append(opts, db.WithValidateField("list_id"))
		if err := fv(ctx, m.GetListId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["result"]; exists {

		vOpts := append(opts, db.WithValidateField("result"))
		if err := fv(ctx, m.GetResult(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIItemListValidator = func() *ValidateAPIItemList {
	v := &ValidateAPIItemList{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIItemListValidator() db.Validator {
	return DefaultAPIItemListValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateAPIAccessReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateAPIAccessReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateAPIAccessReq) DeepCopy() *EvaluateAPIAccessReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateAPIAccessReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateAPIAccessReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateAPIAccessReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateAPIAccessReqValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateAPIAccessReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateAPIAccessReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateAPIAccessReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateAPIAccessReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_lists"]; exists {

		vOpts := append(opts, db.WithValidateField("item_lists"))
		for idx, item := range m.GetItemLists() {
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateAPIAccessReqValidator = func() *ValidateEvaluateAPIAccessReq {
	v := &ValidateEvaluateAPIAccessReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EvaluateAPIAccessReqValidator() db.Validator {
	return DefaultEvaluateAPIAccessReqValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateAPIAccessResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateAPIAccessResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateAPIAccessResp) DeepCopy() *EvaluateAPIAccessResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateAPIAccessResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateAPIAccessResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateAPIAccessResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateAPIAccessRespValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateAPIAccessResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateAPIAccessResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateAPIAccessResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateAPIAccessResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_lists"]; exists {

		vOpts := append(opts, db.WithValidateField("item_lists"))
		for idx, item := range m.GetItemLists() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateAPIAccessRespValidator = func() *ValidateEvaluateAPIAccessResp {
	v := &ValidateEvaluateAPIAccessResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EvaluateAPIAccessRespValidator() db.Validator {
	return DefaultEvaluateAPIAccessRespValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateBatchAPIAccessReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateBatchAPIAccessReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateBatchAPIAccessReq) DeepCopy() *EvaluateBatchAPIAccessReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateBatchAPIAccessReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateBatchAPIAccessReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateBatchAPIAccessReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateBatchAPIAccessReqValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateBatchAPIAccessReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateBatchAPIAccessReq) BatchNamespaceApiListValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for batch_namespace_api_list")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*NamespaceAPIList, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := NamespaceAPIListValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for batch_namespace_api_list")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*NamespaceAPIList)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*NamespaceAPIList, got %T", val)
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
			return errors.Wrap(err, "repeated batch_namespace_api_list")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items batch_namespace_api_list")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateEvaluateBatchAPIAccessReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateBatchAPIAccessReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateBatchAPIAccessReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["batch_namespace_api_list"]; exists {
		vOpts := append(opts, db.WithValidateField("batch_namespace_api_list"))
		if err := fv(ctx, m.GetBatchNamespaceApiList(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateBatchAPIAccessReqValidator = func() *ValidateEvaluateBatchAPIAccessReq {
	v := &ValidateEvaluateBatchAPIAccessReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhBatchNamespaceApiList := v.BatchNamespaceApiListValidationRuleHandler
	rulesBatchNamespaceApiList := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "75",
	}
	vFn, err = vrhBatchNamespaceApiList(rulesBatchNamespaceApiList)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EvaluateBatchAPIAccessReq.batch_namespace_api_list: %s", err)
		panic(errMsg)
	}
	v.FldValidators["batch_namespace_api_list"] = vFn

	return v
}()

func EvaluateBatchAPIAccessReqValidator() db.Validator {
	return DefaultEvaluateBatchAPIAccessReqValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateBatchAPIAccessResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateBatchAPIAccessResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateBatchAPIAccessResp) DeepCopy() *EvaluateBatchAPIAccessResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateBatchAPIAccessResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateBatchAPIAccessResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateBatchAPIAccessResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateBatchAPIAccessRespValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateBatchAPIAccessResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateBatchAPIAccessResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateBatchAPIAccessResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateBatchAPIAccessResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["batch_namespace_api_list"]; exists {

		vOpts := append(opts, db.WithValidateField("batch_namespace_api_list"))
		for idx, item := range m.GetBatchNamespaceApiList() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateBatchAPIAccessRespValidator = func() *ValidateEvaluateBatchAPIAccessResp {
	v := &ValidateEvaluateBatchAPIAccessResp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["batch_namespace_api_list"] = NamespaceAPIListValidator().Validate

	return v
}()

func EvaluateBatchAPIAccessRespValidator() db.Validator {
	return DefaultEvaluateBatchAPIAccessRespValidator
}

// augmented methods on protoc/std generated struct

func (m *LookupUserRolesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LookupUserRolesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LookupUserRolesReq) DeepCopy() *LookupUserRolesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LookupUserRolesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LookupUserRolesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LookupUserRolesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LookupUserRolesReqValidator().Validate(ctx, m, opts...)
}

type ValidateLookupUserRolesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLookupUserRolesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LookupUserRolesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LookupUserRolesReq got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultLookupUserRolesReqValidator = func() *ValidateLookupUserRolesReq {
	v := &ValidateLookupUserRolesReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LookupUserRolesReqValidator() db.Validator {
	return DefaultLookupUserRolesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *LookupUserRolesResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LookupUserRolesResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LookupUserRolesResp) DeepCopy() *LookupUserRolesResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LookupUserRolesResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LookupUserRolesResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LookupUserRolesResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LookupUserRolesRespValidator().Validate(ctx, m, opts...)
}

type ValidateLookupUserRolesResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLookupUserRolesResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LookupUserRolesResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LookupUserRolesResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["roles"]; exists {

		vOpts := append(opts, db.WithValidateField("roles"))
		for idx, item := range m.GetRoles() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLookupUserRolesRespValidator = func() *ValidateLookupUserRolesResp {
	v := &ValidateLookupUserRolesResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LookupUserRolesRespValidator() db.Validator {
	return DefaultLookupUserRolesRespValidator
}

// augmented methods on protoc/std generated struct

func (m *NamespaceAPIList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NamespaceAPIList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NamespaceAPIList) DeepCopy() *NamespaceAPIList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NamespaceAPIList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NamespaceAPIList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NamespaceAPIList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NamespaceAPIListValidator().Validate(ctx, m, opts...)
}

type ValidateNamespaceAPIList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNamespaceAPIList) ItemListsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for item_lists")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*APIItemList, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := APIItemListValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for item_lists")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*APIItemList)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*APIItemList, got %T", val)
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
			return errors.Wrap(err, "repeated item_lists")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items item_lists")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateNamespaceAPIList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NamespaceAPIList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NamespaceAPIList got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_lists"]; exists {
		vOpts := append(opts, db.WithValidateField("item_lists"))
		if err := fv(ctx, m.GetItemLists(), vOpts...); err != nil {
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
var DefaultNamespaceAPIListValidator = func() *ValidateNamespaceAPIList {
	v := &ValidateNamespaceAPIList{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhItemLists := v.ItemListsValidationRuleHandler
	rulesItemLists := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "15",
	}
	vFn, err = vrhItemLists(rulesItemLists)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NamespaceAPIList.item_lists: %s", err)
		panic(errMsg)
	}
	v.FldValidators["item_lists"] = vFn

	return v
}()

func NamespaceAPIListValidator() db.Validator {
	return DefaultNamespaceAPIListValidator
}
