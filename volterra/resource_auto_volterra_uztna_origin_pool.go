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
	ves_io_schema_uztna_uztna_origin_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/uztna/uztna_origin_pool"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraUztnaOriginPool is implementation of Volterra's UztnaOriginPool resources
func resourceVolterraUztnaOriginPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUztnaOriginPoolCreate,
		Read:   resourceVolterraUztnaOriginPoolRead,
		Update: resourceVolterraUztnaOriginPoolUpdate,
		Delete: resourceVolterraUztnaOriginPoolDelete,

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

			"loadbalancer_algorithm": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"least_connections": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"round_robin": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"origin_servers": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"private_ip": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ipv6": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"private_name": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": {
										Type:       schema.TypeString,
										Required:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"refresh_interval": {
										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},
								},
							},
						},

						"labels": {
							Type:     schema.TypeMap,
							Optional: true,
						},

						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},

			"uztna_healthcheck": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
		},
	}
}

// resourceVolterraUztnaOriginPoolCreate creates UztnaOriginPool resource
func resourceVolterraUztnaOriginPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_uztna_uztna_origin_pool.CreateSpecType{}
	createReq := &ves_io_schema_uztna_uztna_origin_pool.CreateRequest{
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

	//loadbalancer_algorithm
	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		loadbalancerAlgorithm := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm{}
		createSpec.LoadbalancerAlgorithm = loadbalancerAlgorithm
		for _, set := range sl {
			if set != nil {
				loadbalancerAlgorithmMapStrToI := set.(map[string]interface{})

				algoChoiceTypeFound := false

				if v, ok := loadbalancerAlgorithmMapStrToI["least_connections"]; ok && !isIntfNil(v) && !algoChoiceTypeFound {

					algoChoiceTypeFound = true

					if v.(bool) {
						algoChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm_LEAST_CONNECTIONS{}
						algoChoiceInt.LEAST_CONNECTIONS = &ves_io_schema.Empty{}
						loadbalancerAlgorithm.AlgoChoice = algoChoiceInt
					}

				}

				if v, ok := loadbalancerAlgorithmMapStrToI["round_robin"]; ok && !isIntfNil(v) && !algoChoiceTypeFound {

					algoChoiceTypeFound = true

					if v.(bool) {
						algoChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm_ROUND_ROBIN{}
						algoChoiceInt.ROUND_ROBIN = &ves_io_schema.Empty{}
						loadbalancerAlgorithm.AlgoChoice = algoChoiceInt
					}

				}

			}
		}

	}

	//origin_servers
	if v, ok := d.GetOk("origin_servers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originServers := make([]*ves_io_schema_uztna_uztna_origin_pool.OriginServerType, len(sl))
		createSpec.OriginServers = originServers
		for i, set := range sl {
			if set != nil {
				originServers[i] = &ves_io_schema_uztna_uztna_origin_pool.OriginServerType{}
				originServersMapStrToI := set.(map[string]interface{})

				choiceTypeFound := false

				if v, ok := originServersMapStrToI["private_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

					choiceTypeFound = true
					choiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerType_PrivateIp{}
					choiceInt.PrivateIp = &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP{}
					originServers[i].Choice = choiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							privateIpChoiceTypeFound := false

							if v, ok := cs["ip"]; ok && !isIntfNil(v) && !privateIpChoiceTypeFound {

								privateIpChoiceTypeFound = true
								privateIpChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP_Ip{}

								choiceInt.PrivateIp.PrivateIpChoice = privateIpChoiceInt

								privateIpChoiceInt.Ip = v.(string)

							}

							if v, ok := cs["ipv6"]; ok && !isIntfNil(v) && !privateIpChoiceTypeFound {

								privateIpChoiceTypeFound = true
								privateIpChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP_Ipv6{}

								choiceInt.PrivateIp.PrivateIpChoice = privateIpChoiceInt

								privateIpChoiceInt.Ipv6 = v.(string)

							}

						}
					}

				}

				if v, ok := originServersMapStrToI["private_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

					choiceTypeFound = true
					choiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerType_PrivateName{}
					choiceInt.PrivateName = &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateName{}
					originServers[i].Choice = choiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["hostname"]; ok && !isIntfNil(v) {

								choiceInt.PrivateName.Hostname = v.(string)

							}

							if v, ok := cs["refresh_interval"]; ok && !isIntfNil(v) {

								choiceInt.PrivateName.RefreshInterval = uint32(v.(int))

							}

						}
					}

				}

				if w, ok := originServersMapStrToI["labels"]; ok && !isIntfNil(w) {
					ms := map[string]string{}
					for k, v := range w.(map[string]interface{}) {
						ms[k] = v.(string)
					}
					originServers[i].Labels = ms
				}

				if w, ok := originServersMapStrToI["port"]; ok && !isIntfNil(w) {
					originServers[i].Port = uint32(w.(int))
				}

			}
		}

	}

	//uztna_healthcheck
	if v, ok := d.GetOk("uztna_healthcheck"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		uztnaHealthcheckInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		createSpec.UztnaHealthcheck = uztnaHealthcheckInt
		for i, ps := range sl {

			uhMapToStrVal := ps.(map[string]interface{})
			uztnaHealthcheckInt[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := uhMapToStrVal["name"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Name = v.(string)
			}

			if v, ok := uhMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Namespace = v.(string)
			}

			if v, ok := uhMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Tenant = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra UztnaOriginPool object with struct: %+v", createReq)

	createUztnaOriginPoolResp, err := client.CreateObject(context.Background(), ves_io_schema_uztna_uztna_origin_pool.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating UztnaOriginPool: %s", err)
	}
	d.SetId(createUztnaOriginPoolResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUztnaOriginPoolRead(d, meta)
}

func resourceVolterraUztnaOriginPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_origin_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaOriginPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaOriginPool %q: %s", d.Id(), err)
	}
	return setUztnaOriginPoolFields(client, d, resp)
}

func setUztnaOriginPoolFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUztnaOriginPoolUpdate updates UztnaOriginPool resource
func resourceVolterraUztnaOriginPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_uztna_uztna_origin_pool.ReplaceSpecType{}
	updateReq := &ves_io_schema_uztna_uztna_origin_pool.ReplaceRequest{
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

	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		loadbalancerAlgorithm := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm{}
		updateSpec.LoadbalancerAlgorithm = loadbalancerAlgorithm
		for _, set := range sl {
			if set != nil {
				loadbalancerAlgorithmMapStrToI := set.(map[string]interface{})

				algoChoiceTypeFound := false

				if v, ok := loadbalancerAlgorithmMapStrToI["least_connections"]; ok && !isIntfNil(v) && !algoChoiceTypeFound {

					algoChoiceTypeFound = true

					if v.(bool) {
						algoChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm_LEAST_CONNECTIONS{}
						algoChoiceInt.LEAST_CONNECTIONS = &ves_io_schema.Empty{}
						loadbalancerAlgorithm.AlgoChoice = algoChoiceInt
					}

				}

				if v, ok := loadbalancerAlgorithmMapStrToI["round_robin"]; ok && !isIntfNil(v) && !algoChoiceTypeFound {

					algoChoiceTypeFound = true

					if v.(bool) {
						algoChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.LoadbalancerAlgorithm_ROUND_ROBIN{}
						algoChoiceInt.ROUND_ROBIN = &ves_io_schema.Empty{}
						loadbalancerAlgorithm.AlgoChoice = algoChoiceInt
					}

				}

			}
		}

	}

	if v, ok := d.GetOk("origin_servers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originServers := make([]*ves_io_schema_uztna_uztna_origin_pool.OriginServerType, len(sl))
		updateSpec.OriginServers = originServers
		for i, set := range sl {
			if set != nil {
				originServers[i] = &ves_io_schema_uztna_uztna_origin_pool.OriginServerType{}
				originServersMapStrToI := set.(map[string]interface{})

				choiceTypeFound := false

				if v, ok := originServersMapStrToI["private_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

					choiceTypeFound = true
					choiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerType_PrivateIp{}
					choiceInt.PrivateIp = &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP{}
					originServers[i].Choice = choiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							privateIpChoiceTypeFound := false

							if v, ok := cs["ip"]; ok && !isIntfNil(v) && !privateIpChoiceTypeFound {

								privateIpChoiceTypeFound = true
								privateIpChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP_Ip{}

								choiceInt.PrivateIp.PrivateIpChoice = privateIpChoiceInt

								privateIpChoiceInt.Ip = v.(string)

							}

							if v, ok := cs["ipv6"]; ok && !isIntfNil(v) && !privateIpChoiceTypeFound {

								privateIpChoiceTypeFound = true
								privateIpChoiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateIP_Ipv6{}

								choiceInt.PrivateIp.PrivateIpChoice = privateIpChoiceInt

								privateIpChoiceInt.Ipv6 = v.(string)

							}

						}
					}

				}

				if v, ok := originServersMapStrToI["private_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

					choiceTypeFound = true
					choiceInt := &ves_io_schema_uztna_uztna_origin_pool.OriginServerType_PrivateName{}
					choiceInt.PrivateName = &ves_io_schema_uztna_uztna_origin_pool.OriginServerPrivateName{}
					originServers[i].Choice = choiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["hostname"]; ok && !isIntfNil(v) {

								choiceInt.PrivateName.Hostname = v.(string)

							}

							if v, ok := cs["refresh_interval"]; ok && !isIntfNil(v) {

								choiceInt.PrivateName.RefreshInterval = uint32(v.(int))

							}

						}
					}

				}

				if w, ok := originServersMapStrToI["labels"]; ok && !isIntfNil(w) {
					ms := map[string]string{}
					for k, v := range w.(map[string]interface{}) {
						ms[k] = v.(string)
					}
					originServers[i].Labels = ms
				}

				if w, ok := originServersMapStrToI["port"]; ok && !isIntfNil(w) {
					originServers[i].Port = uint32(w.(int))
				}

			}
		}

	}

	if v, ok := d.GetOk("uztna_healthcheck"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		uztnaHealthcheckInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		updateSpec.UztnaHealthcheck = uztnaHealthcheckInt
		for i, ps := range sl {

			uhMapToStrVal := ps.(map[string]interface{})
			uztnaHealthcheckInt[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := uhMapToStrVal["name"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Name = v.(string)
			}

			if v, ok := uhMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Namespace = v.(string)
			}

			if v, ok := uhMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				uztnaHealthcheckInt[i].Tenant = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra UztnaOriginPool obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_uztna_uztna_origin_pool.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating UztnaOriginPool: %s", err)
	}

	return resourceVolterraUztnaOriginPoolRead(d, meta)
}

func resourceVolterraUztnaOriginPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_origin_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaOriginPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaOriginPool before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra UztnaOriginPool obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_uztna_uztna_origin_pool.ObjectType, namespace, name)
}
