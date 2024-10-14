// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package geo_config

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

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetSpecDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

// GetDRefInfo for the field's type
func (m *GetResponse) GetSpecDRefInfo() ([]db.DRefInfo, error) {
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

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	v.FldValidators["status"] = StatusObjectValidator().Validate

	return v
}()

func GetResponseValidator() db.Validator {
	return DefaultGetResponseValidator
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
