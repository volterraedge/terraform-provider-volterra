// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package application

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

func (m *Application) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Application) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Application) DeepCopy() *Application {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Application{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Application) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Application) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApplicationValidator().Validate(ctx, m, opts...)
}

type ValidateApplication struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApplication) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Application)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Application got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
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
var DefaultApplicationValidator = func() *ValidateApplication {
	v := &ValidateApplication{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ApplicationValidator() db.Validator {
	return DefaultApplicationValidator
}

// augmented methods on protoc/std generated struct

func (m *GetRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetRequest) DeepCopy() *GetRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetRequest) CmIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cm_id")
	}

	return validatorFn, nil
}

func (v *ValidateGetRequest) ApplicationIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for application_id")
	}

	return validatorFn, nil
}

func (v *ValidateGetRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["application_id"]; exists {

		vOpts := append(opts, db.WithValidateField("application_id"))
		if err := fv(ctx, m.GetApplicationId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["cm_id"]; exists {

		vOpts := append(opts, db.WithValidateField("cm_id"))
		if err := fv(ctx, m.GetCmId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRequestValidator = func() *ValidateGetRequest {
	v := &ValidateGetRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhCmId := v.CmIdValidationRuleHandler
	rulesCmId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhCmId(rulesCmId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetRequest.cm_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cm_id"] = vFn

	vrhApplicationId := v.ApplicationIdValidationRuleHandler
	rulesApplicationId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhApplicationId(rulesApplicationId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetRequest.application_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["application_id"] = vFn

	return v
}()

func GetRequestValidator() db.Validator {
	return DefaultGetRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetResponse) DeepCopy() *GetResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locations"]; exists {

		vOpts := append(opts, db.WithValidateField("locations"))
		for idx, item := range m.GetLocations() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
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
var DefaultGetResponseValidator = func() *ValidateGetResponse {
	v := &ValidateGetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["locations"] = LocationValidator().Validate

	return v
}()

func GetResponseValidator() db.Validator {
	return DefaultGetResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *Links) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Links) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Links) DeepCopy() *Links {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Links{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Links) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Links) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LinksValidator().Validate(ctx, m, opts...)
}

type ValidateLinks struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLinks) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Links)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Links got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["first"]; exists {

		vOpts := append(opts, db.WithValidateField("first"))
		if err := fv(ctx, m.GetFirst(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last"]; exists {

		vOpts := append(opts, db.WithValidateField("last"))
		if err := fv(ctx, m.GetLast(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["next"]; exists {

		vOpts := append(opts, db.WithValidateField("next"))
		if err := fv(ctx, m.GetNext(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["prev"]; exists {

		vOpts := append(opts, db.WithValidateField("prev"))
		if err := fv(ctx, m.GetPrev(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["self"]; exists {

		vOpts := append(opts, db.WithValidateField("self"))
		if err := fv(ctx, m.GetSelf(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLinksValidator = func() *ValidateLinks {
	v := &ValidateLinks{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LinksValidator() db.Validator {
	return DefaultLinksValidator
}

// augmented methods on protoc/std generated struct

func (m *ListRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListRequest) DeepCopy() *ListRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListRequestValidator().Validate(ctx, m, opts...)
}

type ValidateListRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListRequest) CmIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cm_id")
	}

	return validatorFn, nil
}

func (v *ValidateListRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cm_id"]; exists {

		vOpts := append(opts, db.WithValidateField("cm_id"))
		if err := fv(ctx, m.GetCmId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["limit"]; exists {

		vOpts := append(opts, db.WithValidateField("limit"))
		if err := fv(ctx, m.GetLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["page"]; exists {

		vOpts := append(opts, db.WithValidateField("page"))
		if err := fv(ctx, m.GetPage(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListRequestValidator = func() *ValidateListRequest {
	v := &ValidateListRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhCmId := v.CmIdValidationRuleHandler
	rulesCmId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhCmId(rulesCmId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ListRequest.cm_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cm_id"] = vFn

	return v
}()

func ListRequestValidator() db.Validator {
	return DefaultListRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ListResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListResponse) DeepCopy() *ListResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListResponseValidator().Validate(ctx, m, opts...)
}

type ValidateListResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["applications"]; exists {

		vOpts := append(opts, db.WithValidateField("applications"))
		for idx, item := range m.GetApplications() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["count"]; exists {

		vOpts := append(opts, db.WithValidateField("count"))
		if err := fv(ctx, m.GetCount(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["link"]; exists {

		vOpts := append(opts, db.WithValidateField("link"))
		if err := fv(ctx, m.GetLink(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["total"]; exists {

		vOpts := append(opts, db.WithValidateField("total"))
		if err := fv(ctx, m.GetTotal(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListResponseValidator = func() *ValidateListResponse {
	v := &ValidateListResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListResponseValidator() db.Validator {
	return DefaultListResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *Location) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Location) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Location) DeepCopy() *Location {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Location{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Location) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Location) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LocationValidator().Validate(ctx, m, opts...)
}

type ValidateLocation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLocation) IpAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ip_address")
	}

	return validatorFn, nil
}

func (v *ValidateLocation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Location)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Location got type %s", t)
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

	if fv, exists := v.FldValidators["ip_address"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_address"))
		if err := fv(ctx, m.GetIpAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_servers"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_servers"))
		for idx, item := range m.GetVirtualServers() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLocationValidator = func() *ValidateLocation {
	v := &ValidateLocation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhIpAddress := v.IpAddressValidationRuleHandler
	rulesIpAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhIpAddress(rulesIpAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Location.ip_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_address"] = vFn

	v.FldValidators["virtual_servers"] = VirtualServerValidator().Validate

	return v
}()

func LocationValidator() db.Validator {
	return DefaultLocationValidator
}

// augmented methods on protoc/std generated struct

func (m *PoolMember) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PoolMember) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PoolMember) DeepCopy() *PoolMember {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PoolMember{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PoolMember) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PoolMember) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PoolMemberValidator().Validate(ctx, m, opts...)
}

type ValidatePoolMember struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePoolMember) IpAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ip_address")
	}

	return validatorFn, nil
}

func (v *ValidatePoolMember) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PoolMember)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PoolMember got type %s", t)
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

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPoolMemberValidator = func() *ValidatePoolMember {
	v := &ValidatePoolMember{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhIpAddress := v.IpAddressValidationRuleHandler
	rulesIpAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhIpAddress(rulesIpAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PoolMember.ip_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip_address"] = vFn

	return v
}()

func PoolMemberValidator() db.Validator {
	return DefaultPoolMemberValidator
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

	if fv, exists := v.FldValidators["pool_members"]; exists {

		vOpts := append(opts, db.WithValidateField("pool_members"))
		for idx, item := range m.GetPoolMembers() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
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

	v.FldValidators["pool_members"] = PoolMemberValidator().Validate

	return v
}()

func VirtualServerValidator() db.Validator {
	return DefaultVirtualServerValidator
}