//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_service_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
)

// resourceVolterraServicePolicy is implementation of Volterra's ServicePolicy resources
func resourceVolterraServicePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraServicePolicyCreate,
		Read:   resourceVolterraServicePolicyRead,
		Update: resourceVolterraServicePolicyUpdate,
		Delete: resourceVolterraServicePolicyDelete,

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
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"algo": {
				Type:     schema.TypeString,
				Required: true,
			},

			"port_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"ports": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"any_server": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"server_name": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"server_name_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"exact_values": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"regex_values": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"server_selector": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expressions": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraServicePolicyCreate creates ServicePolicy resource
func resourceVolterraServicePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_service_policy.CreateSpecType{}
	createReq := &ves_io_schema_service_policy.CreateRequest{
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

	if v, ok := d.GetOk("algo"); ok && !isIntfNil(v) {

		createSpec.Algo = ves_io_schema_policy.RuleCombiningAlgorithm(ves_io_schema_policy.RuleCombiningAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("port_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		portMatcher := &ves_io_schema_policy.PortMatcherType{}
		createSpec.PortMatcher = portMatcher
		for _, set := range sl {

			portMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := portMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				portMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := portMatcherMapStrToI["ports"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				portMatcher.Ports = ls
			}

		}

	}

	if v, ok := d.GetOk("rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Rules = rulesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			rulesInt[i] = &ves_io_schema.ObjectRefType{}

			rulesInt[i].Kind = "service_policy_rule"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				rulesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rulesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rulesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rulesInt[i].Uid = v.(string)
			}

		}

	}

	serverChoiceTypeFound := false

	if v, ok := d.GetOk("any_server"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true

		if v.(bool) {
			serverChoiceInt := &ves_io_schema_service_policy.CreateSpecType_AnyServer{}
			serverChoiceInt.AnyServer = &ves_io_schema.Empty{}
			createSpec.ServerChoice = serverChoiceInt
		}

	}

	if v, ok := d.GetOk("server_name"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.CreateSpecType_ServerName{}

		createSpec.ServerChoice = serverChoiceInt

		serverChoiceInt.ServerName = v.(string)

	}

	if v, ok := d.GetOk("server_name_matcher"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.CreateSpecType_ServerNameMatcher{}
		serverChoiceInt.ServerNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		createSpec.ServerChoice = serverChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("server_selector"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.CreateSpecType_ServerSelector{}
		serverChoiceInt.ServerSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.ServerChoice = serverChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerSelector.Expressions = ls

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra ServicePolicy object with struct: %+v", createReq)

	createServicePolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_service_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ServicePolicy: %s", err)
	}
	d.SetId(createServicePolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraServicePolicyRead(d, meta)
}

func resourceVolterraServicePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_service_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ServicePolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ServicePolicy %q: %s", d.Id(), err)
	}
	return setServicePolicyFields(client, d, resp)
}

func setServicePolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraServicePolicyUpdate updates ServicePolicy resource
func resourceVolterraServicePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_service_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_service_policy.ReplaceRequest{
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

	if v, ok := d.GetOk("algo"); ok && !isIntfNil(v) {

		updateSpec.Algo = ves_io_schema_policy.RuleCombiningAlgorithm(ves_io_schema_policy.RuleCombiningAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("port_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		portMatcher := &ves_io_schema_policy.PortMatcherType{}
		updateSpec.PortMatcher = portMatcher
		for _, set := range sl {

			portMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := portMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				portMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := portMatcherMapStrToI["ports"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				portMatcher.Ports = ls
			}

		}

	}

	if v, ok := d.GetOk("rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Rules = rulesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			rulesInt[i] = &ves_io_schema.ObjectRefType{}

			rulesInt[i].Kind = "service_policy_rule"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				rulesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rulesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rulesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rulesInt[i].Uid = v.(string)
			}

		}

	}

	serverChoiceTypeFound := false

	if v, ok := d.GetOk("any_server"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true

		if v.(bool) {
			serverChoiceInt := &ves_io_schema_service_policy.ReplaceSpecType_AnyServer{}
			serverChoiceInt.AnyServer = &ves_io_schema.Empty{}
			updateSpec.ServerChoice = serverChoiceInt
		}

	}

	if v, ok := d.GetOk("server_name"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.ReplaceSpecType_ServerName{}

		updateSpec.ServerChoice = serverChoiceInt

		serverChoiceInt.ServerName = v.(string)

	}

	if v, ok := d.GetOk("server_name_matcher"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.ReplaceSpecType_ServerNameMatcher{}
		serverChoiceInt.ServerNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		updateSpec.ServerChoice = serverChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("server_selector"); ok && !serverChoiceTypeFound {

		serverChoiceTypeFound = true
		serverChoiceInt := &ves_io_schema_service_policy.ReplaceSpecType_ServerSelector{}
		serverChoiceInt.ServerSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.ServerChoice = serverChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				serverChoiceInt.ServerSelector.Expressions = ls

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra ServicePolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_service_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ServicePolicy: %s", err)
	}

	return resourceVolterraServicePolicyRead(d, meta)
}

func resourceVolterraServicePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_service_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ServicePolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ServicePolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ServicePolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_service_policy.ObjectType, namespace, name)
}
