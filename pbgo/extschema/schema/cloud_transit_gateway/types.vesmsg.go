// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_transit_gateway

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

func (m *AWSType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AWSType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AWSType) DeepCopy() *AWSType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AWSType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AWSType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AWSType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AWSTypeValidator().Validate(ctx, m, opts...)
}

func (m *AWSType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetCloudCredentialsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCloudCredentialsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetPrimarySecureMeshSiteDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetPrimarySecureMeshSiteDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *AWSType) GetCloudCredentialsDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetCloudCredentials()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("cloud_credentials.Object")
	dri := db.DRefInfo{
		RefdType:   "cloud_credentials.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "cloud_credentials",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetCloudCredentialsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *AWSType) GetCloudCredentialsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "cloud_credentials.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: cloud_credentials")
	}

	vref := m.GetCloudCredentials()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "cloud_credentials.Object",
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

func (m *AWSType) GetPrimarySecureMeshSiteDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetPrimarySecureMeshSite()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("securemesh_site_v2.Object")
	dri := db.DRefInfo{
		RefdType:   "securemesh_site_v2.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "primary_secure_mesh_site",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetPrimarySecureMeshSiteDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *AWSType) GetPrimarySecureMeshSiteDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "securemesh_site_v2.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: securemesh_site_v2")
	}

	vref := m.GetPrimarySecureMeshSite()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "securemesh_site_v2.Object",
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

type ValidateAWSType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAWSType) TgwChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tgw_choice")
	}
	return validatorFn, nil
}

func (v *ValidateAWSType) TgwCidrChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tgw_cidr_choice")
	}
	return validatorFn, nil
}

func (v *ValidateAWSType) AwsRegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for aws_region")
	}

	return validatorFn, nil
}

func (v *ValidateAWSType) ServiceVpcValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for service_vpc")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ServiceVPCTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateAWSType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AWSType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AWSType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["aws_region"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_region"))
		if err := fv(ctx, m.GetAwsRegion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["cloud_credentials"]; exists {

		vOpts := append(opts, db.WithValidateField("cloud_credentials"))
		if err := fv(ctx, m.GetCloudCredentials(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["primary_secure_mesh_site"]; exists {

		vOpts := append(opts, db.WithValidateField("primary_secure_mesh_site"))
		if err := fv(ctx, m.GetPrimarySecureMeshSite(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_vpc"]; exists {

		vOpts := append(opts, db.WithValidateField("service_vpc"))
		if err := fv(ctx, m.GetServiceVpc(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tgw_asn"]; exists {

		vOpts := append(opts, db.WithValidateField("tgw_asn"))
		if err := fv(ctx, m.GetTgwAsn(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tgw_choice"]; exists {
		val := m.GetTgwChoice()
		vOpts := append(opts,
			db.WithValidateField("tgw_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetTgwChoice().(type) {
	case *AWSType_NewTgw:
		if fv, exists := v.FldValidators["tgw_choice.new_tgw"]; exists {
			val := m.GetTgwChoice().(*AWSType_NewTgw).NewTgw
			vOpts := append(opts,
				db.WithValidateField("tgw_choice"),
				db.WithValidateField("new_tgw"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AWSType_ExistingTgw:
		if fv, exists := v.FldValidators["tgw_choice.existing_tgw"]; exists {
			val := m.GetTgwChoice().(*AWSType_ExistingTgw).ExistingTgw
			vOpts := append(opts,
				db.WithValidateField("tgw_choice"),
				db.WithValidateField("existing_tgw"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["tgw_cidr_choice"]; exists {
		val := m.GetTgwCidrChoice()
		vOpts := append(opts,
			db.WithValidateField("tgw_cidr_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetTgwCidrChoice().(type) {
	case *AWSType_ReservedTgwCidr:
		if fv, exists := v.FldValidators["tgw_cidr_choice.reserved_tgw_cidr"]; exists {
			val := m.GetTgwCidrChoice().(*AWSType_ReservedTgwCidr).ReservedTgwCidr
			vOpts := append(opts,
				db.WithValidateField("tgw_cidr_choice"),
				db.WithValidateField("reserved_tgw_cidr"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AWSType_TgwCidr:
		if fv, exists := v.FldValidators["tgw_cidr_choice.tgw_cidr"]; exists {
			val := m.GetTgwCidrChoice().(*AWSType_TgwCidr).TgwCidr
			vOpts := append(opts,
				db.WithValidateField("tgw_cidr_choice"),
				db.WithValidateField("tgw_cidr"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["volterra_site_asn"]; exists {

		vOpts := append(opts, db.WithValidateField("volterra_site_asn"))
		if err := fv(ctx, m.GetVolterraSiteAsn(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAWSTypeValidator = func() *ValidateAWSType {
	v := &ValidateAWSType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhTgwChoice := v.TgwChoiceValidationRuleHandler
	rulesTgwChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhTgwChoice(rulesTgwChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AWSType.tgw_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tgw_choice"] = vFn

	vrhTgwCidrChoice := v.TgwCidrChoiceValidationRuleHandler
	rulesTgwCidrChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhTgwCidrChoice(rulesTgwCidrChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AWSType.tgw_cidr_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tgw_cidr_choice"] = vFn

	vrhAwsRegion := v.AwsRegionValidationRuleHandler
	rulesAwsRegion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAwsRegion(rulesAwsRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AWSType.aws_region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["aws_region"] = vFn

	vrhServiceVpc := v.ServiceVpcValidationRuleHandler
	rulesServiceVpc := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhServiceVpc(rulesServiceVpc)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AWSType.service_vpc: %s", err)
		panic(errMsg)
	}
	v.FldValidators["service_vpc"] = vFn

	v.FldValidators["tgw_choice.new_tgw"] = ves_io_schema_views.TGWParamsTypeValidator().Validate
	v.FldValidators["tgw_choice.existing_tgw"] = ves_io_schema_views.ExistingTGWTypeValidator().Validate

	v.FldValidators["tgw_cidr_choice.tgw_cidr"] = ves_io_schema_views.CloudSubnetParamTypeValidator().Validate

	v.FldValidators["primary_secure_mesh_site"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	v.FldValidators["cloud_credentials"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func AWSTypeValidator() db.Validator {
	return DefaultAWSTypeValidator
}

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

	return m.GetCloudChoiceDRefInfo()

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetCloudChoiceDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCloudChoice() == nil {
		return nil, nil
	}
	switch m.GetCloudChoice().(type) {
	case *CreateSpecType_Aws:

		drInfos, err := m.GetAws().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAws().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "aws." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) CloudChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cloud_choice")
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

	if fv, exists := v.FldValidators["cloud_choice"]; exists {
		val := m.GetCloudChoice()
		vOpts := append(opts,
			db.WithValidateField("cloud_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCloudChoice().(type) {
	case *CreateSpecType_Aws:
		if fv, exists := v.FldValidators["cloud_choice.aws"]; exists {
			val := m.GetCloudChoice().(*CreateSpecType_Aws).Aws
			vOpts := append(opts,
				db.WithValidateField("cloud_choice"),
				db.WithValidateField("aws"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
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

	vrhCloudChoice := v.CloudChoiceValidationRuleHandler
	rulesCloudChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCloudChoice(rulesCloudChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.cloud_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cloud_choice"] = vFn

	v.FldValidators["cloud_choice.aws"] = AWSTypeValidator().Validate

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

	return m.GetCloudChoiceDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSpecType) GetCloudChoiceDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCloudChoice() == nil {
		return nil, nil
	}
	switch m.GetCloudChoice().(type) {
	case *GetSpecType_Aws:

		drInfos, err := m.GetAws().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAws().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "aws." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) CloudChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cloud_choice")
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

	if fv, exists := v.FldValidators["cloud_choice"]; exists {
		val := m.GetCloudChoice()
		vOpts := append(opts,
			db.WithValidateField("cloud_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCloudChoice().(type) {
	case *GetSpecType_Aws:
		if fv, exists := v.FldValidators["cloud_choice.aws"]; exists {
			val := m.GetCloudChoice().(*GetSpecType_Aws).Aws
			vOpts := append(opts,
				db.WithValidateField("cloud_choice"),
				db.WithValidateField("aws"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
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

	vrhCloudChoice := v.CloudChoiceValidationRuleHandler
	rulesCloudChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCloudChoice(rulesCloudChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.cloud_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cloud_choice"] = vFn

	v.FldValidators["cloud_choice.aws"] = AWSTypeValidator().Validate

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

	return m.GetCloudChoiceDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetCloudChoiceDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCloudChoice() == nil {
		return nil, nil
	}
	switch m.GetCloudChoice().(type) {
	case *GlobalSpecType_Aws:

		drInfos, err := m.GetAws().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAws().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "aws." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) CloudChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cloud_choice")
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

	if fv, exists := v.FldValidators["cloud_choice"]; exists {
		val := m.GetCloudChoice()
		vOpts := append(opts,
			db.WithValidateField("cloud_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCloudChoice().(type) {
	case *GlobalSpecType_Aws:
		if fv, exists := v.FldValidators["cloud_choice.aws"]; exists {
			val := m.GetCloudChoice().(*GlobalSpecType_Aws).Aws
			vOpts := append(opts,
				db.WithValidateField("cloud_choice"),
				db.WithValidateField("aws"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
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

	vrhCloudChoice := v.CloudChoiceValidationRuleHandler
	rulesCloudChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCloudChoice(rulesCloudChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.cloud_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cloud_choice"] = vFn

	v.FldValidators["cloud_choice.aws"] = AWSTypeValidator().Validate

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

	return m.GetCloudChoiceDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ReplaceSpecType) GetCloudChoiceDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCloudChoice() == nil {
		return nil, nil
	}
	switch m.GetCloudChoice().(type) {
	case *ReplaceSpecType_Aws:

		drInfos, err := m.GetAws().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAws().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "aws." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) CloudChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cloud_choice")
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

	if fv, exists := v.FldValidators["cloud_choice"]; exists {
		val := m.GetCloudChoice()
		vOpts := append(opts,
			db.WithValidateField("cloud_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCloudChoice().(type) {
	case *ReplaceSpecType_Aws:
		if fv, exists := v.FldValidators["cloud_choice.aws"]; exists {
			val := m.GetCloudChoice().(*ReplaceSpecType_Aws).Aws
			vOpts := append(opts,
				db.WithValidateField("cloud_choice"),
				db.WithValidateField("aws"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
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

	vrhCloudChoice := v.CloudChoiceValidationRuleHandler
	rulesCloudChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCloudChoice(rulesCloudChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.cloud_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cloud_choice"] = vFn

	v.FldValidators["cloud_choice.aws"] = AWSTypeValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ServiceVPCType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceVPCType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceVPCType) DeepCopy() *ServiceVPCType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceVPCType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceVPCType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceVPCType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceVPCTypeValidator().Validate(ctx, m, opts...)
}

type ValidateServiceVPCType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceVPCType) VpcIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for vpc_id")
	}

	return validatorFn, nil
}

func (v *ValidateServiceVPCType) SubnetIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for subnet_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for subnet_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated subnet_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items subnet_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateServiceVPCType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceVPCType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceVPCType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["subnet_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("subnet_ids"))
		if err := fv(ctx, m.GetSubnetIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vpc_id"]; exists {

		vOpts := append(opts, db.WithValidateField("vpc_id"))
		if err := fv(ctx, m.GetVpcId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceVPCTypeValidator = func() *ValidateServiceVPCType {
	v := &ValidateServiceVPCType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhVpcId := v.VpcIdValidationRuleHandler
	rulesVpcId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "64",
		"ves.io.schema.rules.string.pattern":   "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
	}
	vFn, err = vrhVpcId(rulesVpcId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ServiceVPCType.vpc_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vpc_id"] = vFn

	vrhSubnetIds := v.SubnetIdsValidationRuleHandler
	rulesSubnetIds := map[string]string{
		"ves.io.schema.rules.message.required":              "true",
		"ves.io.schema.rules.repeated.items.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$",
		"ves.io.schema.rules.repeated.max_items":            "4",
		"ves.io.schema.rules.repeated.min_items":            "1",
		"ves.io.schema.rules.repeated.unique":               "true",
	}
	vFn, err = vrhSubnetIds(rulesSubnetIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ServiceVPCType.subnet_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["subnet_ids"] = vFn

	return v
}()

func ServiceVPCTypeValidator() db.Validator {
	return DefaultServiceVPCTypeValidator
}

// create setters in CreateSpecType from GlobalSpecType for oneof fields
func (r *CreateSpecType) SetCloudChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.CloudChoice.(type) {
	case nil:
		o.CloudChoice = nil

	case *CreateSpecType_Aws:
		o.CloudChoice = &GlobalSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *CreateSpecType) GetCloudChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.CloudChoice.(type) {
	case nil:
		r.CloudChoice = nil

	case *GlobalSpecType_Aws:
		r.CloudChoice = &CreateSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetCloudChoiceFromGlobalSpecType(f)
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

	m1.SetCloudChoiceToGlobalSpecType(f)
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

// create setters in GetSpecType from GlobalSpecType for oneof fields
func (r *GetSpecType) SetCloudChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.CloudChoice.(type) {
	case nil:
		o.CloudChoice = nil

	case *GetSpecType_Aws:
		o.CloudChoice = &GlobalSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *GetSpecType) GetCloudChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.CloudChoice.(type) {
	case nil:
		r.CloudChoice = nil

	case *GlobalSpecType_Aws:
		r.CloudChoice = &GetSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetCloudChoiceFromGlobalSpecType(f)
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

	m1.SetCloudChoiceToGlobalSpecType(f)
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

// create setters in ReplaceSpecType from GlobalSpecType for oneof fields
func (r *ReplaceSpecType) SetCloudChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.CloudChoice.(type) {
	case nil:
		o.CloudChoice = nil

	case *ReplaceSpecType_Aws:
		o.CloudChoice = &GlobalSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *ReplaceSpecType) GetCloudChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.CloudChoice.(type) {
	case nil:
		r.CloudChoice = nil

	case *GlobalSpecType_Aws:
		r.CloudChoice = &ReplaceSpecType_Aws{Aws: of.Aws}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetCloudChoiceFromGlobalSpecType(f)
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

	m1.SetCloudChoiceToGlobalSpecType(f)
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
