// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package discovered

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

func (m *ApplicationLocation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApplicationLocation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApplicationLocation) DeepCopy() *ApplicationLocation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApplicationLocation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApplicationLocation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApplicationLocation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApplicationLocationValidator().Validate(ctx, m, opts...)
}

type ValidateApplicationLocation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApplicationLocation) IdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for id")
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) InstanceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for instance")
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) HostnameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for hostname")
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) SiteNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_name")
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) IpAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ip_address")
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) VirtualServersValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for virtual_servers")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*VirtualServer, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := VirtualServerValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for virtual_servers")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*VirtualServer)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*VirtualServer, got %T", val)
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
			return errors.Wrap(err, "repeated virtual_servers")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items virtual_servers")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateApplicationLocation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApplicationLocation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApplicationLocation got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["hostname"]; exists {

		vOpts := append(opts, db.WithValidateField("hostname"))
		if err := fv(ctx, m.GetHostname(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance"]; exists {

		vOpts := append(opts, db.WithValidateField("instance"))
		if err := fv(ctx, m.GetInstance(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ip_address"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_address"))
		if err := fv(ctx, m.GetIpAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_servers"]; exists {
		vOpts := append(opts, db.WithValidateField("virtual_servers"))
		if err := fv(ctx, m.GetVirtualServers(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApplicationLocationValidator = func() *ValidateApplicationLocation {
	v := &ValidateApplicationLocation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhId := v.IdValidationRuleHandler
	rulesId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhId(rulesId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["id"] = vFn

	vrhInstance := v.InstanceValidationRuleHandler
	rulesInstance := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhInstance(rulesInstance)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.instance: %s", err)
		panic(errMsg)
	}
	v.FldValidators["instance"] = vFn

	vrhHostname := v.HostnameValidationRuleHandler
	rulesHostname := map[string]string{
		"ves.io.schema.rules.string.hostname": "true",
	}
	vFn, err = vrhHostname(rulesHostname)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.hostname: %s", err)
		panic(errMsg)
	}
	v.FldValidators["hostname"] = vFn

	vrhSiteName := v.SiteNameValidationRuleHandler
	rulesSiteName := map[string]string{
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhSiteName(rulesSiteName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.site_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_name"] = vFn

	vrhIpAddress := v.IpAddressValidationRuleHandler
	rulesIpAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhIpAddress(rulesIpAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.ip_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_address"] = vFn

	vrhVirtualServers := v.VirtualServersValidationRuleHandler
	rulesVirtualServers := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "10",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhVirtualServers(rulesVirtualServers)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApplicationLocation.virtual_servers: %s", err)
		panic(errMsg)
	}
	v.FldValidators["virtual_servers"] = vFn

	return v
}()

func ApplicationLocationValidator() db.Validator {
	return DefaultApplicationLocationValidator
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

func (v *ValidateGetSpecType) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
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

	if fv, exists := v.FldValidators["application"]; exists {

		vOpts := append(opts, db.WithValidateField("application"))
		if err := fv(ctx, m.GetApplication(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance"]; exists {

		vOpts := append(opts, db.WithValidateField("instance"))
		if err := fv(ctx, m.GetInstance(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("private_ip"))
		if err := fv(ctx, m.GetPrivateIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		if err := fv(ctx, m.GetSite(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_server"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_server"))
		if err := fv(ctx, m.GetVirtualServer(), vOpts...); err != nil {
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

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

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

func (v *ValidateGlobalSpecType) CmSiteNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cm_site_name")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) LocationsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for locations")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ApplicationLocation, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ApplicationLocationValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for locations")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ApplicationLocation)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ApplicationLocation, got %T", val)
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
			return errors.Wrap(err, "repeated locations")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items locations")
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

	if fv, exists := v.FldValidators["cm_site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("cm_site_name"))
		if err := fv(ctx, m.GetCmSiteName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locations"]; exists {
		vOpts := append(opts, db.WithValidateField("locations"))
		if err := fv(ctx, m.GetLocations(), vOpts...); err != nil {
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

	vrhCmSiteName := v.CmSiteNameValidationRuleHandler
	rulesCmSiteName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhCmSiteName(rulesCmSiteName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.cm_site_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cm_site_name"] = vFn

	vrhLocations := v.LocationsValidationRuleHandler
	rulesLocations := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "50",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhLocations(rulesLocations)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.locations: %s", err)
		panic(errMsg)
	}
	v.FldValidators["locations"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *VirtualServer) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VirtualServer) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VirtualServer) DeepCopy() *VirtualServer {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VirtualServer{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VirtualServer) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VirtualServer) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VirtualServerValidator().Validate(ctx, m, opts...)
}

type ValidateVirtualServer struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVirtualServer) VsNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for vs_name")
	}

	return validatorFn, nil
}

func (v *ValidateVirtualServer) IpAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ip_address")
	}

	return validatorFn, nil
}

func (v *ValidateVirtualServer) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VirtualServer)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VirtualServer got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ip_address"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_address"))
		if err := fv(ctx, m.GetIpAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vs_name"]; exists {

		vOpts := append(opts, db.WithValidateField("vs_name"))
		if err := fv(ctx, m.GetVsName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVirtualServerValidator = func() *ValidateVirtualServer {
	v := &ValidateVirtualServer{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhVsName := v.VsNameValidationRuleHandler
	rulesVsName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.not_empty": "true",
	}
	vFn, err = vrhVsName(rulesVsName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for VirtualServer.vs_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vs_name"] = vFn

	vrhIpAddress := v.IpAddressValidationRuleHandler
	rulesIpAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhIpAddress(rulesIpAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for VirtualServer.ip_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_address"] = vFn

	return v
}()

func VirtualServerValidator() db.Validator {
	return DefaultVirtualServerValidator
}