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
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetInstanceServerResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetInstanceServerResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetInstanceServerResponse) DeepCopy() *GetInstanceServerResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetInstanceServerResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetInstanceServerResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetInstanceServerResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetInstanceServerResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetInstanceServerResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSpecDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetInstanceServerResponse) GetSpecDRefInfo() ([]db.DRefInfo, error) {
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

type ValidateGetInstanceServerResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetInstanceServerResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetInstanceServerResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetInstanceServerResponse got type %s", t)
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
var DefaultGetInstanceServerResponseValidator = func() *ValidateGetInstanceServerResponse {
	v := &ValidateGetInstanceServerResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	return v
}()

func GetInstanceServerResponseValidator() db.Validator {
	return DefaultGetInstanceServerResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetInstanceServersRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetInstanceServersRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetInstanceServersRequest) DeepCopy() *GetInstanceServersRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetInstanceServersRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetInstanceServersRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetInstanceServersRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetInstanceServersRequestValidator().Validate(ctx, m, opts...)
}

func (m *GetInstanceServersRequest) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetInstanceRefDRefInfo()

}

func (m *GetInstanceServersRequest) GetInstanceRefDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetInstanceRef()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("nginx_instance.Object")
	dri := db.DRefInfo{
		RefdType:   "nginx_instance.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "instance_ref",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetInstanceRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetInstanceServersRequest) GetInstanceRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "nginx_instance.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: nginx_instance")
	}

	vref := m.GetInstanceRef()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "nginx_instance.Object",
		Tenant:    vref.Tenant,
		Namespace: vref.Namespace,
		Name:      vref.Name,
	}
	refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
	if err != nil {
		return nil, errors.Wrap(err, "Getting referred entry")
	}
	if refdEnt != nil {
		entries = append(entries, refdEnt)
	}

	return entries, nil
}

type ValidateGetInstanceServersRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetInstanceServersRequest) InstanceRefValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for instance_ref")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema_views.ObjectRefTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetInstanceServersRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetInstanceServersRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetInstanceServersRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["instance_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("instance_ref"))
		if err := fv(ctx, m.GetInstanceRef(), vOpts...); err != nil {
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
var DefaultGetInstanceServersRequestValidator = func() *ValidateGetInstanceServersRequest {
	v := &ValidateGetInstanceServersRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhInstanceRef := v.InstanceRefValidationRuleHandler
	rulesInstanceRef := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInstanceRef(rulesInstanceRef)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetInstanceServersRequest.instance_ref: %s", err)
		panic(errMsg)
	}
	v.FldValidators["instance_ref"] = vFn

	return v
}()

func GetInstanceServersRequestValidator() db.Validator {
	return DefaultGetInstanceServersRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetInstanceServersResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetInstanceServersResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetInstanceServersResponse) DeepCopy() *GetInstanceServersResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetInstanceServersResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetInstanceServersResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetInstanceServersResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetInstanceServersResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetInstanceServersResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetItemsDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetInstanceServersResponse) GetItemsDRefInfo() ([]db.DRefInfo, error) {
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

type ValidateGetInstanceServersResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetInstanceServersResponse) ItemsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for items")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*GetInstanceServerResponse, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := GetInstanceServerResponseValidator().Validate(ctx, el, opts...); err != nil {
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
		elems, ok := val.([]*GetInstanceServerResponse)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*GetInstanceServerResponse, got %T", val)
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

func (v *ValidateGetInstanceServersResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetInstanceServersResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetInstanceServersResponse got type %s", t)
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
var DefaultGetInstanceServersResponseValidator = func() *ValidateGetInstanceServersResponse {
	v := &ValidateGetInstanceServersResponse{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetInstanceServersResponse.items: %s", err)
		panic(errMsg)
	}
	v.FldValidators["items"] = vFn

	return v
}()

func GetInstanceServersResponseValidator() db.Validator {
	return DefaultGetInstanceServersResponseValidator
}