// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_leasepool

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

func (m *CreateRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateRequest) DeepCopy() *CreateRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCreateRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateRequest got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateRequestValidator = func() *ValidateCreateRequest {
	v := &ValidateCreateRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectCreateMetaTypeValidator().Validate

	v.FldValidators["spec"] = CreateSpecTypeValidator().Validate

	return v
}()

func CreateRequestValidator() db.Validator {
	return DefaultCreateRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateResponse) DeepCopy() *CreateResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateResponseValidator().Validate(ctx, m, opts...)
}

type ValidateCreateResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateResponse got type %s", t)
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
var DefaultCreateResponseValidator = func() *ValidateCreateResponse {
	v := &ValidateCreateResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	return v
}()

func CreateResponseValidator() db.Validator {
	return DefaultCreateResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *DeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DeleteRequest) DeepCopy() *DeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DeleteRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["fail_if_referred"]; exists {

		vOpts := append(opts, db.WithValidateField("fail_if_referred"))
		if err := fv(ctx, m.GetFailIfReferred(), vOpts...); err != nil {
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
var DefaultDeleteRequestValidator = func() *ValidateDeleteRequest {
	v := &ValidateDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DeleteRequestValidator() db.Validator {
	return DefaultDeleteRequestValidator
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

	if fv, exists := v.FldValidators["response_format"]; exists {

		vOpts := append(opts, db.WithValidateField("response_format"))
		if err := fv(ctx, m.GetResponseFormat(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRequestValidator = func() *ValidateGetRequest {
	v := &ValidateGetRequest{FldValidators: map[string]db.ValidatorFunc{}}

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

func (m *GetResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return nil, nil

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

	if fv, exists := v.FldValidators["create_form"]; exists {

		vOpts := append(opts, db.WithValidateField("create_form"))
		if err := fv(ctx, m.GetCreateForm(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["deleted_referred_objects"]; exists {

		vOpts := append(opts, db.WithValidateField("deleted_referred_objects"))
		for idx, item := range m.GetDeletedReferredObjects() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["disabled_referred_objects"]; exists {

		vOpts := append(opts, db.WithValidateField("disabled_referred_objects"))
		for idx, item := range m.GetDisabledReferredObjects() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["referring_objects"]; exists {

		vOpts := append(opts, db.WithValidateField("referring_objects"))
		for idx, item := range m.GetReferringObjects() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["replace_form"]; exists {

		vOpts := append(opts, db.WithValidateField("replace_form"))
		if err := fv(ctx, m.GetReplaceForm(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["resource_version"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_version"))
		if err := fv(ctx, m.GetResourceVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
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
var DefaultGetResponseValidator = func() *ValidateGetResponse {
	v := &ValidateGetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["create_form"] = CreateRequestValidator().Validate

	v.FldValidators["replace_form"] = ReplaceRequestValidator().Validate

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	v.FldValidators["status"] = StatusObjectValidator().Validate

	return v
}()

func GetResponseValidator() db.Validator {
	return DefaultGetResponseValidator
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

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		if err := fv(ctx, m.GetLabelFilter(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["report_fields"]; exists {

		vOpts := append(opts, db.WithValidateField("report_fields"))
		for idx, item := range m.GetReportFields() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["report_status_fields"]; exists {

		vOpts := append(opts, db.WithValidateField("report_status_fields"))
		for idx, item := range m.GetReportStatusFields() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListRequestValidator = func() *ValidateListRequest {
	v := &ValidateListRequest{FldValidators: map[string]db.ValidatorFunc{}}

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

func (m *ListResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetItemsDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ListResponse) GetItemsDRefInfo() ([]db.DRefInfo, error) {
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

	if fv, exists := v.FldValidators["errors"]; exists {

		vOpts := append(opts, db.WithValidateField("errors"))
		for idx, item := range m.GetErrors() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

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
var DefaultListResponseValidator = func() *ValidateListResponse {
	v := &ValidateListResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["items"] = ListResponseItemValidator().Validate

	return v
}()

func ListResponseValidator() db.Validator {
	return DefaultListResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListResponseItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListResponseItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListResponseItem) DeepCopy() *ListResponseItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListResponseItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListResponseItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListResponseItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListResponseItemValidator().Validate(ctx, m, opts...)
}

func (m *ListResponseItem) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return nil, nil

}

type ValidateListResponseItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListResponseItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListResponseItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListResponseItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["annotations"]; exists {

		vOpts := append(opts, db.WithValidateField("annotations"))
		for key, value := range m.GetAnnotations() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["disabled"]; exists {

		vOpts := append(opts, db.WithValidateField("disabled"))
		if err := fv(ctx, m.GetDisabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["get_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("get_spec"))
		if err := fv(ctx, m.GetGetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["labels"]; exists {

		vOpts := append(opts, db.WithValidateField("labels"))
		for key, value := range m.GetLabels() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["owner_view"]; exists {

		vOpts := append(opts, db.WithValidateField("owner_view"))
		if err := fv(ctx, m.GetOwnerView(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status_set"]; exists {

		vOpts := append(opts, db.WithValidateField("status_set"))
		for idx, item := range m.GetStatusSet() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["uid"]; exists {

		vOpts := append(opts, db.WithValidateField("uid"))
		if err := fv(ctx, m.GetUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListResponseItemValidator = func() *ValidateListResponseItem {
	v := &ValidateListResponseItem{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["get_spec"] = GetSpecTypeValidator().Validate

	v.FldValidators["status_set"] = StatusObjectValidator().Validate

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	return v
}()

func ListResponseItemValidator() db.Validator {
	return DefaultListResponseItemValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceRequest) DeepCopy() *ReplaceRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceRequestValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceRequest got type %s", t)
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

	if fv, exists := v.FldValidators["resource_version"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_version"))
		if err := fv(ctx, m.GetResourceVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceRequestValidator = func() *ValidateReplaceRequest {
	v := &ValidateReplaceRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectReplaceMetaTypeValidator().Validate

	v.FldValidators["spec"] = ReplaceSpecTypeValidator().Validate

	return v
}()

func ReplaceRequestValidator() db.Validator {
	return DefaultReplaceRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceResponse) DeepCopy() *ReplaceResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceResponseValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceResponseValidator = func() *ValidateReplaceResponse {
	v := &ValidateReplaceResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReplaceResponseValidator() db.Validator {
	return DefaultReplaceResponseValidator
}

func (m *CreateRequest) fromObject(e db.Entry, withDeepCopy bool) {
	f := e.(*DBObject)
	if withDeepCopy {
		f = e.DeepCopy().(*DBObject)
	}
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectCreateMetaType{}
	}
	m.Metadata.FromObjectMetaTypeWithoutDeepCopy(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &CreateSpecType{}
	}
	m.Spec.FromGlobalSpecTypeWithoutDeepCopy(f.GetSpec().GetGcSpec())

}

func (m *CreateRequest) FromObject(e db.Entry) {
	m.fromObject(e, true)
}

func (m *CreateRequest) FromObjectWithoutDeepCopy(e db.Entry) {
	m.fromObject(e, false)
}

func (m *CreateRequest) toObject(e db.Entry, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaTypeWithoutDeepCopy(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecTypeWithoutDeepCopy(f.Spec.GcSpec)
	}

}

func (m *CreateRequest) ToObject(e db.Entry) {
	m.toObject(e, true)
}

func (m *CreateRequest) ToObjectWithoutDeepCopy(e db.Entry) {
	m.toObject(e, false)
}

func (m *CreateResponse) fromObject(e db.Entry, withDeepCopy bool) {
	f := e.(*DBObject)
	if withDeepCopy {
		f = e.DeepCopy().(*DBObject)
	}
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectGetMetaType{}
	}
	m.Metadata.FromObjectMetaTypeWithoutDeepCopy(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &GetSpecType{}
	}
	m.Spec.FromGlobalSpecTypeWithoutDeepCopy(f.GetSpec().GetGcSpec())

	if m.SystemMetadata == nil {
		m.SystemMetadata = &ves_io_schema.SystemObjectGetMetaType{}
	}
	m.SystemMetadata.FromSystemObjectMetaTypeWithoutDeepCopy(f.GetSystemMetadata())

}

func (m *CreateResponse) FromObject(e db.Entry) {
	m.fromObject(e, true)
}

func (m *CreateResponse) FromObjectWithoutDeepCopy(e db.Entry) {
	m.fromObject(e, false)
}

func (m *CreateResponse) toObject(e db.Entry, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaTypeWithoutDeepCopy(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecTypeWithoutDeepCopy(f.Spec.GcSpec)
	}

	if m1.SystemMetadata != nil {
		if f.SystemMetadata == nil {
			f.SystemMetadata = &ves_io_schema.SystemObjectMetaType{}
		}
	} else if f.SystemMetadata != nil {
		f.SystemMetadata = nil
	}

	if m1.SystemMetadata != nil {
		m1.SystemMetadata.ToSystemObjectMetaTypeWithoutDeepCopy(f.SystemMetadata)
	}

}

func (m *CreateResponse) ToObject(e db.Entry) {
	m.toObject(e, true)
}

func (m *CreateResponse) ToObjectWithoutDeepCopy(e db.Entry) {
	m.toObject(e, false)
}

func (m *GetResponse) fromObject(e db.Entry, withDeepCopy bool) {
	f := e.(*DBObject)
	if withDeepCopy {
		f = e.DeepCopy().(*DBObject)
	}
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectGetMetaType{}
	}
	m.Metadata.FromObjectMetaTypeWithoutDeepCopy(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &GetSpecType{}
	}
	m.Spec.FromGlobalSpecTypeWithoutDeepCopy(f.GetSpec().GetGcSpec())

	if m.SystemMetadata == nil {
		m.SystemMetadata = &ves_io_schema.SystemObjectGetMetaType{}
	}
	m.SystemMetadata.FromSystemObjectMetaTypeWithoutDeepCopy(f.GetSystemMetadata())

}

func (m *GetResponse) FromObject(e db.Entry) {
	m.fromObject(e, true)
}

func (m *GetResponse) FromObjectWithoutDeepCopy(e db.Entry) {
	m.fromObject(e, false)
}

func (m *GetResponse) toObject(e db.Entry, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaTypeWithoutDeepCopy(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecTypeWithoutDeepCopy(f.Spec.GcSpec)
	}

	if m1.SystemMetadata != nil {
		if f.SystemMetadata == nil {
			f.SystemMetadata = &ves_io_schema.SystemObjectMetaType{}
		}
	} else if f.SystemMetadata != nil {
		f.SystemMetadata = nil
	}

	if m1.SystemMetadata != nil {
		m1.SystemMetadata.ToSystemObjectMetaTypeWithoutDeepCopy(f.SystemMetadata)
	}

}

func (m *GetResponse) ToObject(e db.Entry) {
	m.toObject(e, true)
}

func (m *GetResponse) ToObjectWithoutDeepCopy(e db.Entry) {
	m.toObject(e, false)
}

func (m *ReplaceRequest) fromObject(e db.Entry, withDeepCopy bool) {
	f := e.(*DBObject)
	if withDeepCopy {
		f = e.DeepCopy().(*DBObject)
	}
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectReplaceMetaType{}
	}
	m.Metadata.FromObjectMetaTypeWithoutDeepCopy(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &ReplaceSpecType{}
	}
	m.Spec.FromGlobalSpecTypeWithoutDeepCopy(f.GetSpec().GetGcSpec())

}

func (m *ReplaceRequest) FromObject(e db.Entry) {
	m.fromObject(e, true)
}

func (m *ReplaceRequest) FromObjectWithoutDeepCopy(e db.Entry) {
	m.fromObject(e, false)
}

func (m *ReplaceRequest) toObject(e db.Entry, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaTypeWithoutDeepCopy(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecTypeWithoutDeepCopy(f.Spec.GcSpec)
	}

}

func (m *ReplaceRequest) ToObject(e db.Entry) {
	m.toObject(e, true)
}

func (m *ReplaceRequest) ToObjectWithoutDeepCopy(e db.Entry) {
	m.toObject(e, false)
}