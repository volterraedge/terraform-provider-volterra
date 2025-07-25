//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_views_ike_phase1_profile "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/ike_phase1_profile"
	ves_io_schema_views_ike_phase2_profile "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/ike_phase2_profile"
)

// resourceVolterraIkePhase2Profile is implementation of Volterra's IkePhase2Profile resources
func resourceVolterraIkePhase2Profile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraIkePhase2ProfileCreate,
		Read:   resourceVolterraIkePhase2ProfileRead,
		Update: resourceVolterraIkePhase2ProfileUpdate,
		Delete: resourceVolterraIkePhase2ProfileDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"authentication_algos": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"encryption_algos": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"ike_keylifetime_hours": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"duration": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"ike_keylifetime_minutes": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"duration": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"use_default_keylifetime": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dh_group_set": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dh_groups": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"disable_pfs": {

				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// resourceVolterraIkePhase2ProfileCreate creates IkePhase2Profile resource
func resourceVolterraIkePhase2ProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_ike_phase2_profile.CreateSpecType{}
	createReq := &ves_io_schema_views_ike_phase2_profile.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	//authentication_algos
	if v, ok := d.GetOk("authentication_algos"); ok && !isIntfNil(v) {

		authentication_algosList := []ves_io_schema.AuthenticationAlgorithm{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field authentication_algos")
			}
			authentication_algosList = append(authentication_algosList, ves_io_schema.AuthenticationAlgorithm(ves_io_schema.AuthenticationAlgorithm_value[j.(string)]))
		}
		createSpec.AuthenticationAlgos = authentication_algosList

	}

	//encryption_algos
	if v, ok := d.GetOk("encryption_algos"); ok && !isIntfNil(v) {

		encryption_algosList := []ves_io_schema.EncryptionAlgorithm{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field encryption_algos")
			}
			encryption_algosList = append(encryption_algosList, ves_io_schema.EncryptionAlgorithm(ves_io_schema.EncryptionAlgorithm_value[j.(string)]))
		}
		createSpec.EncryptionAlgos = encryption_algosList

	}

	//ike_key_lifetime

	ikeKeyLifetimeTypeFound := false

	if v, ok := d.GetOk("ike_keylifetime_hours"); ok && !isIntfNil(v) && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true
		ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.CreateSpecType_IkeKeylifetimeHours{}
		ikeKeyLifetimeInt.IkeKeylifetimeHours = &ves_io_schema_views_ike_phase1_profile.InputHours{}
		createSpec.IkeKeyLifetime = ikeKeyLifetimeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["duration"]; ok && !isIntfNil(v) {

					ikeKeyLifetimeInt.IkeKeylifetimeHours.Duration = uint32(v.(int))

				}

			}
		}

	}

	if v, ok := d.GetOk("ike_keylifetime_minutes"); ok && !isIntfNil(v) && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true
		ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.CreateSpecType_IkeKeylifetimeMinutes{}
		ikeKeyLifetimeInt.IkeKeylifetimeMinutes = &ves_io_schema_views_ike_phase1_profile.InputMinutes{}
		createSpec.IkeKeyLifetime = ikeKeyLifetimeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["duration"]; ok && !isIntfNil(v) {

					ikeKeyLifetimeInt.IkeKeylifetimeMinutes.Duration = uint32(v.(int))

				}

			}
		}

	}

	if v, ok := d.GetOk("use_default_keylifetime"); ok && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true

		if v.(bool) {
			ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.CreateSpecType_UseDefaultKeylifetime{}
			ikeKeyLifetimeInt.UseDefaultKeylifetime = &ves_io_schema.Empty{}
			createSpec.IkeKeyLifetime = ikeKeyLifetimeInt
		}

	}

	//pfs_mode

	pfsModeTypeFound := false

	if v, ok := d.GetOk("dh_group_set"); ok && !isIntfNil(v) && !pfsModeTypeFound {

		pfsModeTypeFound = true
		pfsModeInt := &ves_io_schema_views_ike_phase2_profile.CreateSpecType_DhGroupSet{}
		pfsModeInt.DhGroupSet = &ves_io_schema_views_ike_phase2_profile.DHGroups{}
		createSpec.PfsMode = pfsModeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["dh_groups"]; ok && !isIntfNil(v) {

					dh_groupsList := []ves_io_schema.DHGroup{}
					for _, j := range v.([]interface{}) {
						if j == nil {
							return fmt.Errorf("please provide valid non-empty enum value of field dh_groups")
						}
						dh_groupsList = append(dh_groupsList, ves_io_schema.DHGroup(ves_io_schema.DHGroup_value[j.(string)]))
					}
					pfsModeInt.DhGroupSet.DhGroups = dh_groupsList

				}

			}
		}

	}

	if v, ok := d.GetOk("disable_pfs"); ok && !pfsModeTypeFound {

		pfsModeTypeFound = true

		if v.(bool) {
			pfsModeInt := &ves_io_schema_views_ike_phase2_profile.CreateSpecType_DisablePfs{}
			pfsModeInt.DisablePfs = &ves_io_schema.Empty{}
			createSpec.PfsMode = pfsModeInt
		}

	}

	log.Printf("[DEBUG] Creating Volterra IkePhase2Profile object with struct: %+v", createReq)

	createIkePhase2ProfileResp, err := client.CreateObject(context.Background(), ves_io_schema_views_ike_phase2_profile.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating IkePhase2Profile: %s", err)
	}
	d.SetId(createIkePhase2ProfileResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraIkePhase2ProfileRead(d, meta)
}

func resourceVolterraIkePhase2ProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_ike_phase2_profile.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] IkePhase2Profile %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra IkePhase2Profile %q: %s", d.Id(), err)
	}
	return setIkePhase2ProfileFields(client, d, resp)
}

func setIkePhase2ProfileFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraIkePhase2ProfileUpdate updates IkePhase2Profile resource
func resourceVolterraIkePhase2ProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_ike_phase2_profile.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("authentication_algos"); ok && !isIntfNil(v) {

		authentication_algosList := []ves_io_schema.AuthenticationAlgorithm{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field authentication_algos")
			}
			authentication_algosList = append(authentication_algosList, ves_io_schema.AuthenticationAlgorithm(ves_io_schema.AuthenticationAlgorithm_value[j.(string)]))
		}
		updateSpec.AuthenticationAlgos = authentication_algosList

	}

	if v, ok := d.GetOk("encryption_algos"); ok && !isIntfNil(v) {

		encryption_algosList := []ves_io_schema.EncryptionAlgorithm{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field encryption_algos")
			}
			encryption_algosList = append(encryption_algosList, ves_io_schema.EncryptionAlgorithm(ves_io_schema.EncryptionAlgorithm_value[j.(string)]))
		}
		updateSpec.EncryptionAlgos = encryption_algosList

	}

	ikeKeyLifetimeTypeFound := false

	if v, ok := d.GetOk("ike_keylifetime_hours"); ok && !isIntfNil(v) && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true
		ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType_IkeKeylifetimeHours{}
		ikeKeyLifetimeInt.IkeKeylifetimeHours = &ves_io_schema_views_ike_phase1_profile.InputHours{}
		updateSpec.IkeKeyLifetime = ikeKeyLifetimeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["duration"]; ok && !isIntfNil(v) {

					ikeKeyLifetimeInt.IkeKeylifetimeHours.Duration = uint32(v.(int))

				}

			}
		}

	}

	if v, ok := d.GetOk("ike_keylifetime_minutes"); ok && !isIntfNil(v) && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true
		ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType_IkeKeylifetimeMinutes{}
		ikeKeyLifetimeInt.IkeKeylifetimeMinutes = &ves_io_schema_views_ike_phase1_profile.InputMinutes{}
		updateSpec.IkeKeyLifetime = ikeKeyLifetimeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["duration"]; ok && !isIntfNil(v) {

					ikeKeyLifetimeInt.IkeKeylifetimeMinutes.Duration = uint32(v.(int))

				}

			}
		}

	}

	if v, ok := d.GetOk("use_default_keylifetime"); ok && !ikeKeyLifetimeTypeFound {

		ikeKeyLifetimeTypeFound = true

		if v.(bool) {
			ikeKeyLifetimeInt := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType_UseDefaultKeylifetime{}
			ikeKeyLifetimeInt.UseDefaultKeylifetime = &ves_io_schema.Empty{}
			updateSpec.IkeKeyLifetime = ikeKeyLifetimeInt
		}

	}

	pfsModeTypeFound := false

	if v, ok := d.GetOk("dh_group_set"); ok && !isIntfNil(v) && !pfsModeTypeFound {

		pfsModeTypeFound = true
		pfsModeInt := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType_DhGroupSet{}
		pfsModeInt.DhGroupSet = &ves_io_schema_views_ike_phase2_profile.DHGroups{}
		updateSpec.PfsMode = pfsModeInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["dh_groups"]; ok && !isIntfNil(v) {

					dh_groupsList := []ves_io_schema.DHGroup{}
					for _, j := range v.([]interface{}) {
						if j == nil {
							return fmt.Errorf("please provide valid non-empty enum value of field dh_groups")
						}
						dh_groupsList = append(dh_groupsList, ves_io_schema.DHGroup(ves_io_schema.DHGroup_value[j.(string)]))
					}
					pfsModeInt.DhGroupSet.DhGroups = dh_groupsList

				}

			}
		}

	}

	if v, ok := d.GetOk("disable_pfs"); ok && !pfsModeTypeFound {

		pfsModeTypeFound = true

		if v.(bool) {
			pfsModeInt := &ves_io_schema_views_ike_phase2_profile.ReplaceSpecType_DisablePfs{}
			pfsModeInt.DisablePfs = &ves_io_schema.Empty{}
			updateSpec.PfsMode = pfsModeInt
		}

	}

	log.Printf("[DEBUG] Updating Volterra IkePhase2Profile obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_ike_phase2_profile.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating IkePhase2Profile: %s", err)
	}

	return resourceVolterraIkePhase2ProfileRead(d, meta)
}

func resourceVolterraIkePhase2ProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_ike_phase2_profile.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] IkePhase2Profile %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra IkePhase2Profile before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra IkePhase2Profile obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_views_ike_phase2_profile.ObjectType, namespace, name, opts...)
}
