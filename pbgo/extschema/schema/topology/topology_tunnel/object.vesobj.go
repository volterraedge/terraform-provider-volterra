// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package topology_tunnel

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	google_protobuf "github.com/gogo/protobuf/types"
	multierror "github.com/hashicorp/go-multierror"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/store"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	topology "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/topology"

	"github.com/google/uuid"
	"gopkg.volterra.us/stdlib/db/sro"
)

const (
	ObjectDefTblName = "ves.io.schema.topology.topology_tunnel.Object.default"
	ObjectType       = "ves.io.schema.topology.topology_tunnel.Object"
)

// augmented methods on protoc/std generated struct
func (e *Object) Type() string {
	return "ves.io.schema.topology.topology_tunnel.Object"
}

func (e *Object) ToEntry() db.Entry {
	return NewDBObject(e, db.OpWithNoCopy())
}

func LocateObject(ctx context.Context, locator db.EntryLocator, uid, tenant, namespace, name string, opts ...db.FindEntryOpt) (*DBObject, error) {
	timestamp, err := google_protobuf.TimestampProto(time.Now())
	if err != nil {
		return nil, errors.Wrapf(err, "%s: LocateObject", uid)
	}
	if uid != "" {
		obj, exist, err := FindObject(ctx, locator, uid, opts...)
		if err != nil {
			return nil, errors.Wrapf(err, "%s: LocateObject", uid)
		}
		if exist && obj != nil {
			obj.SystemMetadata.ModificationTimestamp = timestamp
			return obj, nil
		}
	} else {
		uid = uuid.New().String()
	}

	sysMD := &ves_io_schema.SystemObjectMetaType{
		Uid:                   uid,
		Tenant:                tenant,
		CreatorClass:          locator.GetCreatorClass(),
		CreatorId:             locator.GetCreatorID(),
		CreationTimestamp:     timestamp,
		ModificationTimestamp: timestamp,
	}
	obj := NewDBObject(nil)
	obj.SetObjUid(uid)
	obj.SetObjName(name)
	obj.SetObjNamespace(namespace)
	obj.SetObjSystemMetadata(sysMD)
	obj.Spec = &SpecType{}
	return obj, nil
}

func FindObject(ctx context.Context, finder db.EntryFinder, key string, opts ...db.FindEntryOpt) (*DBObject, bool, error) {
	e, exist, err := finder.FindEntry(ctx, ObjectDefTblName, key, opts...)
	if !exist || err != nil {
		return nil, exist, err
	}
	obj, ok := e.(*DBObject)
	if !ok {
		return nil, false, fmt.Errorf("Cannot convert entry to object")
	}
	return obj, exist, err
}

func ListObject(ctx context.Context, lister db.EntryLister, opts ...db.ListEntriesOpt) ([]*DBObject, error) {
	var (
		oList []*DBObject
		merr  *multierror.Error
	)
	eList, err := lister.ListEntries(ctx, ObjectDefTblName, opts...)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	for _, e := range eList {
		obj, ok := e.(*DBObject)
		if ok {
			oList = append(oList, obj)
		} else {
			merr = multierror.Append(merr, fmt.Errorf("Cannot convert entry to %s object", ObjectType))
		}
	}
	return oList, errors.ErrOrNil(merr)
}

func (o *Object) DeepCopy() *Object {
	if o == nil {
		return nil
	}
	ser, err := o.Marshal()
	if err != nil {
		return nil
	}
	c := &Object{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (e *Object) ToJSON() (string, error) {
	return codec.ToJSON(e)
}

func (e *Object) ToYAML() (string, error) {
	return codec.ToYAML(e)
}

func (e *Object) GetTraceInfo() string {
	sysMD := e.GetSystemMetadata()
	if sysMD == nil {
		return ""
	}
	return sysMD.GetTraceInfo()
}

// A struct wrapping protoc/std generated struct with additional capabilities
// forming a db.Entry
type DBObject struct {
	// Anonymous embed of standard protobuf generated struct
	*Object

	tbl db.Table
}

// GetObjectIndexers returns the associated store.Indexers for Object
func GetObjectIndexers() store.Indexers {

	return nil

}

func (e *DBObject) GetDB() (*db.DB, error) {
	if e.tbl == nil {
		return nil, fmt.Errorf("Entry has no table")
	}
	return e.tbl.GetDB(), nil
}

// Implement ves.io/stdlib/db.Entry interface
func (e *DBObject) Key(opts ...db.KeyOpt) (string, error) {
	ko := db.NewKeyOpts(opts...)
	if ko.Public {
		md := e.GetMetadata()
		if md == nil {
			return "", fmt.Errorf("Metadata is nil")
		}
		return fmt.Sprintf("%s/%s", md.GetNamespace(), md.GetName()), nil
	} else {
		if e.GetSystemMetadata() == nil {
			return "", fmt.Errorf("SystemMetadata is nil")
		}
		return e.GetSystemMetadata().GetUid(), nil
	}
}

func (e *DBObject) Type() string {
	return "ves.io.schema.topology.topology_tunnel.Object"
}

func (e *DBObject) DeepCopy() db.Entry {
	if e == nil {
		return nil
	}
	n := NewDBObject(e.Object)
	n.tbl = e.tbl
	return n
}

func (e *DBObject) MarshalBytes() ([]byte, error) {
	return e.Marshal()
}

func (e *DBObject) UnmarshalBytes(b []byte) error {
	return e.Unmarshal(b)
}

func (e *DBObject) Sample(r *rand.Rand) (db.Entry, error) {
	uid := uuid.New().String()
	o := &Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Name:      uuid.New().String(),
			Namespace: uuid.New().String(),
			Uid:       uid,
		},
		SystemMetadata: &ves_io_schema.SystemObjectMetaType{
			Uid:    uid,
			Tenant: uuid.New().String(),
		},
		Spec: &SpecType{},
	}

	return &DBObject{o, e.tbl}, nil
}

func (e *DBObject) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectValidator().Validate(ctx, e.Object, opts...)
}

func (e *DBObject) SetBlob(ctx context.Context, bID string, bVal interface{}, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "SetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "SetBlob accessing key")
	}
	err = db.SetEntryBlob(ctx, key, e.Type(), bID, bVal, opts...)
	if err != nil {
		return errors.Wrap(err, "SetBlob setting in db")
	}
	return nil
}

func (e *DBObject) ClrBlob(ctx context.Context, bID string, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "ClrBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "ClrBlob accessing key")
	}
	err = db.ClrEntryBlob(ctx, key, e.Type(), bID, opts...)
	if err != nil {
		return errors.Wrap(err, "ClrBlob clearing in db")
	}
	return nil
}

func (e *DBObject) GetBlob(ctx context.Context, bID string, opts ...db.BlobOpt) (interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob accessing key")
	}
	return db.GetEntryBlob(ctx, key, e.Type(), bID, opts...)
}

func (e *DBObject) GetBlobs(ctx context.Context, opts ...db.BlobOpt) (map[string]interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs accessing key")
	}
	return db.GetEntryBlobs(ctx, key, e.Type(), opts...)
}

func (e *DBObject) IsDeleted() (bool, error) {
	db, err := e.GetDB()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted")
	}
	key, err := e.Key()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing key")
	}
	isDel, err := db.IsEntryDeleted(key, e.Type())
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing db")
	}
	return isDel, nil
}

// Implement ves.io/stdlib/db.EntryPvt interface
func (e *DBObject) SetTable(tbl db.Table) {
	e.tbl = tbl
}

func (e *DBObject) GetDRefInfo() ([]db.DRefInfo, error) {
	if e == nil {
		return nil, nil
	}
	refrUID, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetDRefInfo, error in key")
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := e.GetSpecDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetSpecDRefInfo() FAILED")
	} else {
		for i := range fdrInfos {
			dri := &fdrInfos[i]
			// Convert Spec.LcSpec.vnRefs to ves.io.examplesvc.objectone.Object.Spec.LcSpec.vnRefs
			dri.DRField = "ves.io.schema.topology.topology_tunnel.Object." + dri.DRField
			dri.RefrType = e.Type()
			dri.RefrUID = refrUID

			// convert any ref_to schema annotation specified by kind value to type value
			if !strings.HasPrefix(dri.RefdType, "ves.io") {
				d, err := e.GetDB()
				if err != nil {
					return nil, errors.Wrap(err, "Cannot find db for entry to resolve kind to type")
				}
				refdType, err := d.TypeForEntryKind(dri.RefrType, dri.RefrUID, dri.RefdType)
				if err != nil {
					return nil, errors.Wrap(err, fmt.Sprintf("Cannot convert kind %s to type", dri.RefdType))
				}
				dri.RefdType = refdType
			}
		}
		drInfos = append(drInfos, fdrInfos...)
	}
	if fdrInfos, err := e.GetSystemMetadataDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetSystemMetadataDRefInfo() FAILED")
	} else {
		for i := range fdrInfos {
			dri := &fdrInfos[i]
			// Convert Spec.LcSpec.vnRefs to ves.io.examplesvc.objectone.Object.Spec.LcSpec.vnRefs
			dri.DRField = "ves.io.schema.topology.topology_tunnel.Object." + dri.DRField
			dri.RefrType = e.Type()
			dri.RefrUID = refrUID

			// convert any ref_to schema annotation specified by kind value to type value
			if !strings.HasPrefix(dri.RefdType, "ves.io") {
				d, err := e.GetDB()
				if err != nil {
					return nil, errors.Wrap(err, "Cannot find db for entry to resolve kind to type")
				}
				refdType, err := d.TypeForEntryKind(dri.RefrType, dri.RefrUID, dri.RefdType)
				if err != nil {
					return nil, errors.Wrap(err, fmt.Sprintf("Cannot convert kind %s to type", dri.RefdType))
				}
				dri.RefdType = refdType
			}
		}
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (e *DBObject) ToStore() store.Entry {
	return e.Object
}

func (e *DBObject) ToJSON() (string, error) {
	return e.ToStore().ToJSON()
}

func (e *DBObject) ToYAML() (string, error) {
	return e.ToStore().ToYAML()
}

func (e *DBObject) GetTable() db.Table {
	return e.tbl
}

func NewDBObject(o *Object, opts ...db.OpOption) *DBObject {
	op := db.NewOpFrom(opts...)
	if o == nil {
		return &DBObject{Object: &Object{}}
	}
	obj := o
	if !op.NoCopy() {
		obj = o.DeepCopy()
	}
	return &DBObject{Object: obj}
}

func NewEntryObject(opts ...db.OpOption) db.Entry {
	op := db.NewOpFrom(opts...)
	s := op.StoreEntry()
	switch v := s.(type) {
	case nil:
		return NewDBObject(nil, opts...)
	case *Object:
		return NewDBObject(v, opts...)
	}
	return nil
}

// GetDRefInfo for the field's type
func (e *DBObject) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	if e.GetSpec() == nil {
		return nil, nil
	}

	drInfos, err := e.GetSpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "spec." + dri.DRField
	}
	return drInfos, err

}

// GetDRefInfo for the field's type
func (e *DBObject) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	if e.GetSystemMetadata() == nil {
		return nil, nil
	}

	drInfos, err := e.GetSystemMetadata().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSystemMetadata().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "system_metadata." + dri.DRField
	}
	return drInfos, err

}

// Implement sro.SRO interface
func (o *DBObject) GetObjMetadata() sro.ObjectMetadata {
	if o.GetMetadata() == nil {
		return nil
	}
	return o.GetMetadata()
}

func (o *DBObject) SetObjMetadata(in sro.ObjectMetadata) error {
	if in == nil {
		o.Metadata = nil
		return nil
	}

	m, ok := in.(*ves_io_schema.ObjectMetaType)
	if !ok {
		return fmt.Errorf("Error: SetObjMetadata expected *ObjectMetaType, got %T", in)
	}
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjSystemMetadata() sro.SystemMetadata {
	if o.GetSystemMetadata() == nil {
		return nil
	}
	return o.GetSystemMetadata()
}

func (o *DBObject) SetObjSystemMetadata(in sro.SystemMetadata) error {
	if in == nil {
		o.SystemMetadata = nil
		return nil
	}

	m, ok := in.(*ves_io_schema.SystemObjectMetaType)
	if !ok {
		return fmt.Errorf("Error: SetObjSystemMetadata expected *SystemObjectMetaType, got %T", in)
	}
	o.SystemMetadata = m
	return nil
}

func (e *DBObject) SetDirectRefHash(s string) {
	m := e.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.DirectRefHash = s
}

func (e *DBObject) GetDirectRefHash() string {
	m := e.GetSystemMetadata()
	if m == nil {
		return ""
	}
	return m.DirectRefHash
}

func (e *DBObject) GetModificationTimestamp() *google_protobuf.Timestamp {
	m := e.GetSystemMetadata()
	if m == nil {
		return nil
	}
	return m.ModificationTimestamp
}

func (o *DBObject) GetObjSpec() sro.Spec {
	if o.GetSpec() == nil {
		return nil
	}
	return o.GetSpec()
}

func (o *DBObject) SetObjSpec(in sro.Spec) error {
	if in == nil {
		o.Spec = nil
		return nil
	}

	m, ok := in.(*SpecType)
	if !ok {
		return fmt.Errorf("Error: SetObjSpec expected *SpecType, got %T", in)
	}
	o.Spec = m
	return nil
}

func (o *DBObject) GetObjType() string {
	return o.Type()
}

// GetObjUid returns uuid from source-of-truth, in systemMetadata
func (o *DBObject) GetObjUid() string {
	return o.GetSystemMetadata().GetUid()
}

// GetObjTenant returns tenant from source-of-truth, in systemMetadata
func (o *DBObject) GetObjTenant() string {
	return o.GetSystemMetadata().GetTenant()
}

// GetObjCreatorClass returns creator-class from systemMetadata
func (o *DBObject) GetObjCreatorClass() string {
	return o.GetSystemMetadata().GetCreatorClass()
}

// GetObjectIndex returns object-index from systemMetadata
func (o *DBObject) GetObjectIndex() uint32 {
	return o.GetSystemMetadata().GetObjectIndex()
}

// SetObjUid sets uuid as a hint, in Metadata
func (o *DBObject) SetObjUid(u string) error {
	// TODO: make sure 'u' is of uuid form
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Uid = u
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjName() string {
	return o.GetMetadata().GetName()
}

func (o *DBObject) SetObjName(n string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Name = n
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjNamespace() string {
	return o.GetMetadata().GetNamespace()
}

func (o *DBObject) SetObjNamespace(ns string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Namespace = ns
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjLabels() map[string]string {
	return o.GetMetadata().GetLabels()
}

func (o *DBObject) SetObjLabels(l map[string]string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Labels = l
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjAnnotations() map[string]string {
	return o.GetMetadata().GetAnnotations()
}

func (o *DBObject) SetObjAnnotations(a map[string]string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Annotations = a
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjDescription() string {
	return o.GetMetadata().GetDescription()
}

func (o *DBObject) SetObjDescription(d string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Description = d
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjDisable() bool {
	return o.GetMetadata().GetDisable()
}

func (o *DBObject) SetObjDisable(d bool) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Disable = d
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjSREDisable() bool {
	return o.GetSystemMetadata().GetSreDisable()
}

func (o *DBObject) SetObjSREDisable(d bool) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.SreDisable = d
	return nil
}

func (o *DBObject) SetObjCreator(cls, inst string) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.CreatorClass = cls
	m.CreatorId = inst
	o.SystemMetadata = m
	return nil
}

func (o *DBObject) SetObjectIndex(idx uint32) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.ObjectIndex = idx
	o.SystemMetadata = m
	return nil
}

func (o *DBObject) GetObjFinalizers() []string {
	return o.GetSystemMetadata().GetFinalizers()
}

func (o *DBObject) SetObjFinalizers(values ...string) error {
	m := o.GetSystemMetadata()
	if m == nil {
		return fmt.Errorf("Object has nil system_metadata")
	}
	m.Finalizers = values
	return nil
}

func (o *DBObject) GetObjPendingInitializers() []string {
	initializers := o.GetSystemMetadata().GetInitializers()
	var pending []string
	for _, p := range initializers.GetPending() {
		pending = append(pending, p.GetName())
	}
	return pending
}

func (o *DBObject) SetObjPendingInitializers(pending ...string) {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
		o.SystemMetadata = m
	}
	initializers := m.GetInitializers()
	if initializers == nil {
		initializers = &ves_io_schema.InitializersType{}
		m.Initializers = initializers
	}
	var pendingInitializers []*ves_io_schema.InitializerType
	for _, p := range pending {
		pendingInitializers = append(pendingInitializers, &ves_io_schema.InitializerType{Name: p})
	}
	initializers.Pending = pendingInitializers
}

func (o *DBObject) IsSpecEqual(other sro.SRO) bool {
	otherObjSpec := other.GetObjSpec()
	otherSpec, ok := otherObjSpec.(*SpecType)
	if !ok {
		return false
	}

	return o.GetSpec().Equal(otherSpec)
}

// GetVtrpId returns vtrpId of the object.
func (o *DBObject) GetVtrpId() string {
	return o.GetSystemMetadata().GetVtrpId()
}

// SetVtrpId sets vtrpId of the object.
func (o *DBObject) SetVtrpId(id string) {
	o.GetSystemMetadata().SetVtrpId(id)
}

// GetVtrpStale returns true if the object is stale in Mars
func (o *DBObject) GetVtrpStale() bool {
	return o.GetSystemMetadata().GetVtrpStale()
}

// SetVtrpStale sets vtrpStale on the object
func (o *DBObject) SetVtrpStale(isStale bool) {
	o.GetSystemMetadata().SetVtrpStale(isStale)
}

type ValidateObject struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObject) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	e, ok := pm.(*Object)
	if !ok {
		switch t := pm.(type) {
		default:
			return fmt.Errorf("Expected type *Object got type %s", t)
		}
	}
	if e == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, e.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, e.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, e.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectValidator = func() *ValidateObject {
	v := &ValidateObject{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = SpecTypeValidator().Validate

	return v
}()

func ObjectValidator() db.Validator {
	return DefaultObjectValidator
}

func (o *DBObject) GetTopologyMetadata() *topology.MetaType {
	return o.GetSpec().GetGcSpec().GetTopologyMetadata()
}

func (o *DBObject) SetTopologyMetadata(t *topology.MetaType) {
	o.GetSpec().GetGcSpec().TopologyMetadata = t
}
