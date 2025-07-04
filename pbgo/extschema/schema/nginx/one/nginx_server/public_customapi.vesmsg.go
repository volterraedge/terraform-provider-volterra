// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package nginx_server

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

func (m *GetDataplaneServerResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDataplaneServerResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDataplaneServerResponse) DeepCopy() *GetDataplaneServerResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDataplaneServerResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDataplaneServerResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDataplaneServerResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDataplaneServerResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetDataplaneServerResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSpecDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetDataplaneServerResponse) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	if m.GetSpec() == nil {
		return nil, nil
	}

	drInfos, err := m.GetSpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "spec." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetDataplaneServerResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDataplaneServerResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDataplaneServerResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDataplaneServerResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDataplaneServerResponseValidator = func() *ValidateGetDataplaneServerResponse {
	v := &ValidateGetDataplaneServerResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	return v
}()

func GetDataplaneServerResponseValidator() db.Validator {
	return DefaultGetDataplaneServerResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetDataplaneServersRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDataplaneServersRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDataplaneServersRequest) DeepCopy() *GetDataplaneServersRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDataplaneServersRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDataplaneServersRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDataplaneServersRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDataplaneServersRequestValidator().Validate(ctx, m, opts...)
}

func (m *GetDataplaneServersRequest) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetDataplaneRefDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetDataplaneServersRequest) GetDataplaneRefDRefInfo() ([]db.DRefInfo, error) {
	if m.GetDataplaneRef() == nil {
		return nil, nil
	}

	drInfos, err := m.GetDataplaneRef().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetDataplaneRef().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "dataplane_ref." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetDataplaneServersRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDataplaneServersRequest) DataplaneRefValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for dataplane_ref")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := DataplaneReferenceValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetDataplaneServersRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDataplaneServersRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDataplaneServersRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dataplane_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("dataplane_ref"))
		if err := fv(ctx, m.GetDataplaneRef(), vOpts...); err != nil {
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
var DefaultGetDataplaneServersRequestValidator = func() *ValidateGetDataplaneServersRequest {
	v := &ValidateGetDataplaneServersRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhDataplaneRef := v.DataplaneRefValidationRuleHandler
	rulesDataplaneRef := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhDataplaneRef(rulesDataplaneRef)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetDataplaneServersRequest.dataplane_ref: %s", err)
		panic(errMsg)
	}
	v.FldValidators["dataplane_ref"] = vFn

	return v
}()

func GetDataplaneServersRequestValidator() db.Validator {
	return DefaultGetDataplaneServersRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetDataplaneServersResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDataplaneServersResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDataplaneServersResponse) DeepCopy() *GetDataplaneServersResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDataplaneServersResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDataplaneServersResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDataplaneServersResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDataplaneServersResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetDataplaneServersResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetItemsDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetDataplaneServersResponse) GetItemsDRefInfo() ([]db.DRefInfo, error) {
	if m.GetItems() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetItems() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetItems() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("items[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateGetDataplaneServersResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDataplaneServersResponse) ItemsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for items")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*GetDataplaneServerResponse, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := GetDataplaneServerResponseValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for items")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*GetDataplaneServerResponse)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*GetDataplaneServerResponse, got %T", val)
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
			return errors.Wrap(err, "repeated items")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items items")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetDataplaneServersResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDataplaneServersResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDataplaneServersResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["items"]; exists {
		vOpts := append(opts, db.WithValidateField("items"))
		if err := fv(ctx, m.GetItems(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDataplaneServersResponseValidator = func() *ValidateGetDataplaneServersResponse {
	v := &ValidateGetDataplaneServersResponse{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhItems := v.ItemsValidationRuleHandler
	rulesItems := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhItems(rulesItems)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetDataplaneServersResponse.items: %s", err)
		panic(errMsg)
	}
	v.FldValidators["items"] = vFn

	return v
}()

func GetDataplaneServersResponseValidator() db.Validator {
	return DefaultGetDataplaneServersResponseValidator
}
