// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package ztna_application

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

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProxyAdvertisementDRefInfo()

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetProxyAdvertisementDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProxyAdvertisement() == nil {
		return nil, nil
	}

	drInfos, err := m.GetProxyAdvertisement().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetProxyAdvertisement().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "proxy_advertisement." + dri.DRField
	}
	return drInfos, err

}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) TransportTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(TransportType)
		return int32(i)
	}
	// TransportType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, TransportType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for transport_type")
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

	if fv, exists := v.FldValidators["domain_name"]; exists {

		vOpts := append(opts, db.WithValidateField("domain_name"))
		if err := fv(ctx, m.GetDomainName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["msg"]; exists {

		vOpts := append(opts, db.WithValidateField("msg"))
		if err := fv(ctx, m.GetMsg(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policies"]; exists {

		vOpts := append(opts, db.WithValidateField("policies"))
		if err := fv(ctx, m.GetPolicies(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_advertisement"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_advertisement"))
		if err := fv(ctx, m.GetProxyAdvertisement(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["transport_type"]; exists {

		vOpts := append(opts, db.WithValidateField("transport_type"))
		if err := fv(ctx, m.GetTransportType(), vOpts...); err != nil {
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

	vrhTransportType := v.TransportTypeValidationRuleHandler
	rulesTransportType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTransportType(rulesTransportType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.transport_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["transport_type"] = vFn

	v.FldValidators["msg"] = ServiceDetailsValidator().Validate

	v.FldValidators["proxy_advertisement"] = ProxyAdvertisementTypeValidator().Validate

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

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProxyAdvertisementDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSpecType) GetProxyAdvertisementDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProxyAdvertisement() == nil {
		return nil, nil
	}

	drInfos, err := m.GetProxyAdvertisement().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetProxyAdvertisement().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "proxy_advertisement." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) TransportTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(TransportType)
		return int32(i)
	}
	// TransportType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, TransportType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for transport_type")
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

	if fv, exists := v.FldValidators["domain_name"]; exists {

		vOpts := append(opts, db.WithValidateField("domain_name"))
		if err := fv(ctx, m.GetDomainName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["msg"]; exists {

		vOpts := append(opts, db.WithValidateField("msg"))
		if err := fv(ctx, m.GetMsg(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policies"]; exists {

		vOpts := append(opts, db.WithValidateField("policies"))
		if err := fv(ctx, m.GetPolicies(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_advertisement"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_advertisement"))
		if err := fv(ctx, m.GetProxyAdvertisement(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["transport_type"]; exists {

		vOpts := append(opts, db.WithValidateField("transport_type"))
		if err := fv(ctx, m.GetTransportType(), vOpts...); err != nil {
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

	vrhTransportType := v.TransportTypeValidationRuleHandler
	rulesTransportType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTransportType(rulesTransportType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.transport_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["transport_type"] = vFn

	v.FldValidators["msg"] = ServiceDetailsValidator().Validate

	v.FldValidators["proxy_advertisement"] = ProxyAdvertisementTypeValidator().Validate

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

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetProxyAdvertisementDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetProxyAdvertisementDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetViewInternalDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetViewInternalDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetProxyAdvertisementDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProxyAdvertisement() == nil {
		return nil, nil
	}

	drInfos, err := m.GetProxyAdvertisement().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetProxyAdvertisement().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "proxy_advertisement." + dri.DRField
	}
	return drInfos, err

}

func (m *GlobalSpecType) GetViewInternalDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetViewInternal()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("view_internal.Object")
	dri := db.DRefInfo{
		RefdType:   "view_internal.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "view_internal",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetViewInternalDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetViewInternalDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "view_internal.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: view_internal")
	}

	vref := m.GetViewInternal()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "view_internal.Object",
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

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) TransportTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(TransportType)
		return int32(i)
	}
	// TransportType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, TransportType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for transport_type")
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

	if fv, exists := v.FldValidators["domain_name"]; exists {

		vOpts := append(opts, db.WithValidateField("domain_name"))
		if err := fv(ctx, m.GetDomainName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["msg"]; exists {

		vOpts := append(opts, db.WithValidateField("msg"))
		if err := fv(ctx, m.GetMsg(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policies"]; exists {

		vOpts := append(opts, db.WithValidateField("policies"))
		if err := fv(ctx, m.GetPolicies(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_advertisement"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_advertisement"))
		if err := fv(ctx, m.GetProxyAdvertisement(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["transport_type"]; exists {

		vOpts := append(opts, db.WithValidateField("transport_type"))
		if err := fv(ctx, m.GetTransportType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_internal"]; exists {

		vOpts := append(opts, db.WithValidateField("view_internal"))
		if err := fv(ctx, m.GetViewInternal(), vOpts...); err != nil {
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

	vrhTransportType := v.TransportTypeValidationRuleHandler
	rulesTransportType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTransportType(rulesTransportType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.transport_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["transport_type"] = vFn

	v.FldValidators["msg"] = ServiceDetailsValidator().Validate

	v.FldValidators["proxy_advertisement"] = ProxyAdvertisementTypeValidator().Validate

	v.FldValidators["view_internal"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ProxyAdvertisementType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ProxyAdvertisementType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ProxyAdvertisementType) DeepCopy() *ProxyAdvertisementType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ProxyAdvertisementType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ProxyAdvertisementType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ProxyAdvertisementType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ProxyAdvertisementTypeValidator().Validate(ctx, m, opts...)
}

func (m *ProxyAdvertisementType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetAdvertiseChoiceDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ProxyAdvertisementType) GetAdvertiseChoiceDRefInfo() ([]db.DRefInfo, error) {
	if m.GetAdvertiseChoice() == nil {
		return nil, nil
	}
	switch m.GetAdvertiseChoice().(type) {
	case *ProxyAdvertisementType_DoNotAdvertise:

		return nil, nil

	case *ProxyAdvertisementType_AdvertiseOnPublicDefaultVip:

		return nil, nil

	case *ProxyAdvertisementType_AdvertiseOnPublic:

		drInfos, err := m.GetAdvertiseOnPublic().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAdvertiseOnPublic().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "advertise_on_public." + dri.DRField
		}
		return drInfos, err

	case *ProxyAdvertisementType_AdvertiseCustom:

		drInfos, err := m.GetAdvertiseCustom().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAdvertiseCustom().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "advertise_custom." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateProxyAdvertisementType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateProxyAdvertisementType) AdvertiseChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for advertise_choice")
	}
	return validatorFn, nil
}

func (v *ValidateProxyAdvertisementType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ProxyAdvertisementType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ProxyAdvertisementType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["advertise_choice"]; exists {
		val := m.GetAdvertiseChoice()
		vOpts := append(opts,
			db.WithValidateField("advertise_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetAdvertiseChoice().(type) {
	case *ProxyAdvertisementType_DoNotAdvertise:
		if fv, exists := v.FldValidators["advertise_choice.do_not_advertise"]; exists {
			val := m.GetAdvertiseChoice().(*ProxyAdvertisementType_DoNotAdvertise).DoNotAdvertise
			vOpts := append(opts,
				db.WithValidateField("advertise_choice"),
				db.WithValidateField("do_not_advertise"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProxyAdvertisementType_AdvertiseOnPublicDefaultVip:
		if fv, exists := v.FldValidators["advertise_choice.advertise_on_public_default_vip"]; exists {
			val := m.GetAdvertiseChoice().(*ProxyAdvertisementType_AdvertiseOnPublicDefaultVip).AdvertiseOnPublicDefaultVip
			vOpts := append(opts,
				db.WithValidateField("advertise_choice"),
				db.WithValidateField("advertise_on_public_default_vip"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProxyAdvertisementType_AdvertiseOnPublic:
		if fv, exists := v.FldValidators["advertise_choice.advertise_on_public"]; exists {
			val := m.GetAdvertiseChoice().(*ProxyAdvertisementType_AdvertiseOnPublic).AdvertiseOnPublic
			vOpts := append(opts,
				db.WithValidateField("advertise_choice"),
				db.WithValidateField("advertise_on_public"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProxyAdvertisementType_AdvertiseCustom:
		if fv, exists := v.FldValidators["advertise_choice.advertise_custom"]; exists {
			val := m.GetAdvertiseChoice().(*ProxyAdvertisementType_AdvertiseCustom).AdvertiseCustom
			vOpts := append(opts,
				db.WithValidateField("advertise_choice"),
				db.WithValidateField("advertise_custom"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultProxyAdvertisementTypeValidator = func() *ValidateProxyAdvertisementType {
	v := &ValidateProxyAdvertisementType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAdvertiseChoice := v.AdvertiseChoiceValidationRuleHandler
	rulesAdvertiseChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhAdvertiseChoice(rulesAdvertiseChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ProxyAdvertisementType.advertise_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["advertise_choice"] = vFn

	v.FldValidators["advertise_choice.advertise_on_public"] = ves_io_schema_views.AdvertisePublicValidator().Validate
	v.FldValidators["advertise_choice.advertise_custom"] = ves_io_schema_views.AdvertiseCustomValidator().Validate

	return v
}()

func ProxyAdvertisementTypeValidator() db.Validator {
	return DefaultProxyAdvertisementTypeValidator
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

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProxyAdvertisementDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ReplaceSpecType) GetProxyAdvertisementDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProxyAdvertisement() == nil {
		return nil, nil
	}

	drInfos, err := m.GetProxyAdvertisement().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetProxyAdvertisement().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "proxy_advertisement." + dri.DRField
	}
	return drInfos, err

}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) TransportTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(TransportType)
		return int32(i)
	}
	// TransportType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, TransportType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for transport_type")
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

	if fv, exists := v.FldValidators["domain_name"]; exists {

		vOpts := append(opts, db.WithValidateField("domain_name"))
		if err := fv(ctx, m.GetDomainName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["msg"]; exists {

		vOpts := append(opts, db.WithValidateField("msg"))
		if err := fv(ctx, m.GetMsg(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policies"]; exists {

		vOpts := append(opts, db.WithValidateField("policies"))
		if err := fv(ctx, m.GetPolicies(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_advertisement"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_advertisement"))
		if err := fv(ctx, m.GetProxyAdvertisement(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["transport_type"]; exists {

		vOpts := append(opts, db.WithValidateField("transport_type"))
		if err := fv(ctx, m.GetTransportType(), vOpts...); err != nil {
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

	vrhTransportType := v.TransportTypeValidationRuleHandler
	rulesTransportType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTransportType(rulesTransportType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.transport_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["transport_type"] = vFn

	v.FldValidators["msg"] = ServiceDetailsValidator().Validate

	v.FldValidators["proxy_advertisement"] = ProxyAdvertisementTypeValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ServiceDetails) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceDetails) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceDetails) DeepCopy() *ServiceDetails {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceDetails{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceDetails) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceDetails) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceDetailsValidator().Validate(ctx, m, opts...)
}

type ValidateServiceDetails struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceDetails) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceDetails)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceDetails got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ip"]; exists {

		vOpts := append(opts, db.WithValidateField("ip"))
		if err := fv(ctx, m.GetIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceDetailsValidator = func() *ValidateServiceDetails {
	v := &ValidateServiceDetails{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["ip"] = ves_io_schema.IpAddressTypeValidator().Validate

	return v
}()

func ServiceDetailsValidator() db.Validator {
	return DefaultServiceDetailsValidator
}

// augmented methods on protoc/std generated struct

func (m *TileAccess) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *TileAccess) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *TileAccess) DeepCopy() *TileAccess {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &TileAccess{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *TileAccess) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *TileAccess) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TileAccessValidator().Validate(ctx, m, opts...)
}

type ValidateTileAccess struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTileAccess) DisplayNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for display_name")
	}

	return validatorFn, nil
}

func (v *ValidateTileAccess) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*TileAccess)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *TileAccess got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["display_name"]; exists {

		vOpts := append(opts, db.WithValidateField("display_name"))
		if err := fv(ctx, m.GetDisplayName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTileAccessValidator = func() *ValidateTileAccess {
	v := &ValidateTileAccess{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhDisplayName := v.DisplayNameValidationRuleHandler
	rulesDisplayName := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhDisplayName(rulesDisplayName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for TileAccess.display_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["display_name"] = vFn

	return v
}()

func TileAccessValidator() db.Validator {
	return DefaultTileAccessValidator
}

// augmented methods on protoc/std generated struct

func (m *ZTNApolicies) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ZTNApolicies) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ZTNApolicies) DeepCopy() *ZTNApolicies {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ZTNApolicies{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ZTNApolicies) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ZTNApolicies) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ZTNApoliciesValidator().Validate(ctx, m, opts...)
}

type ValidateZTNApolicies struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateZTNApolicies) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ZTNApolicies)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ZTNApolicies got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["connectivity_policy_name"]; exists {

		vOpts := append(opts, db.WithValidateField("connectivity_policy_name"))
		if err := fv(ctx, m.GetConnectivityPolicyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["message_policy_name"]; exists {

		vOpts := append(opts, db.WithValidateField("message_policy_name"))
		if err := fv(ctx, m.GetMessagePolicyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["session_policy_name"]; exists {

		vOpts := append(opts, db.WithValidateField("session_policy_name"))
		if err := fv(ctx, m.GetSessionPolicyName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultZTNApoliciesValidator = func() *ValidateZTNApolicies {
	v := &ValidateZTNApolicies{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ZTNApoliciesValidator() db.Validator {
	return DefaultZTNApoliciesValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.DomainName = f.GetDomainName()
	m.Msg = f.GetMsg()
	m.Policies = f.GetPolicies()
	m.ProxyAdvertisement = f.GetProxyAdvertisement()
	m.TransportType = f.GetTransportType()
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

	f.DomainName = m1.DomainName
	f.Msg = m1.Msg
	f.Policies = m1.Policies
	f.ProxyAdvertisement = m1.ProxyAdvertisement
	f.TransportType = m1.TransportType
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
	m.DomainName = f.GetDomainName()
	m.Msg = f.GetMsg()
	m.Policies = f.GetPolicies()
	m.ProxyAdvertisement = f.GetProxyAdvertisement()
	m.TransportType = f.GetTransportType()
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

	f.DomainName = m1.DomainName
	f.Msg = m1.Msg
	f.Policies = m1.Policies
	f.ProxyAdvertisement = m1.ProxyAdvertisement
	f.TransportType = m1.TransportType
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
	m.DomainName = f.GetDomainName()
	m.Msg = f.GetMsg()
	m.Policies = f.GetPolicies()
	m.ProxyAdvertisement = f.GetProxyAdvertisement()
	m.TransportType = f.GetTransportType()
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

	f.DomainName = m1.DomainName
	f.Msg = m1.Msg
	f.Policies = m1.Policies
	f.ProxyAdvertisement = m1.ProxyAdvertisement
	f.TransportType = m1.TransportType
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}