// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/svcfw"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
)

func TestAPICredentialReadNamespaces(t *testing.T) {
	tests := []struct {
		name     string
		config   map[string]interface{}
		expected []string
	}{
		{
			name: "kubeconfig uses vk8s namespace with system fallback",
			config: map[string]interface{}{
				"api_credential_type":   "KUBE_CONFIG",
				"virtual_k8s_namespace": "app-ns",
			},
			expected: []string{"app-ns", svcfw.SystemNSVal},
		},
		{
			name: "api token remains system scoped",
			config: map[string]interface{}{
				"api_credential_type":   "API_TOKEN",
				"virtual_k8s_namespace": "app-ns",
			},
			expected: []string{svcfw.SystemNSVal},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := schema.TestResourceDataRaw(t, resourceVolterraAPICredential().Schema, test.config)
			if actual := apiCredentialReadNamespaces(data); !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected namespaces %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestAPICredentialReadClearsStateForBackendNotFound(t *testing.T) {
	data := schema.TestResourceDataRaw(t, resourceVolterraAPICredential().Schema, map[string]interface{}{
		"api_credential_type":   "KUBE_CONFIG",
		"virtual_k8s_namespace": "app-ns",
	})
	data.SetId("credential-1")
	client := &fakeCredentialAPIClient{
		errors: map[string]error{
			"/public/namespaces/app-ns/api_credentials/credential-1": fmt.Errorf("object not found"),
			"/public/namespaces/system/api_credentials/credential-1": fmt.Errorf("object not found"),
		},
		responses: map[string]proto.Message{},
	}

	if err := resourceVolterraAPICredentialRead(data, client); err != nil {
		t.Fatalf("read returned an error: %v", err)
	}
	if data.Id() != "" {
		t.Fatalf("expected missing credential to be removed from state, got ID %q", data.Id())
	}
}

func TestAccAPICredentialBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_api_credential." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")

	timeString := currentTime()
	getCreateResponseAPICredential(name)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testKubeConfigCredential(ves_io_schema_api_credential.ObjectType, name, timeString),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAPICertCredential(ves_io_schema_api_credential.ObjectType, name, timeString),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testKubeConfigCredential(objType, name, timeString string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  api_credential_type = "KUBE_CONFIG"
		  expiry_days = 10
		  created_at = "%s"
		}
		`, tfResourceName, name, name, timeString)
}

func testAPICertCredential(objType, name, timeString string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  api_credential_type = "API_CERTIFICATE"
		  created_at = "%s"
		}
		`, tfResourceName, name, name, timeString)
}

func currentTime() string {
	//code to crearte time.Now and RFC366 return a string value
	return time.Now().UTC().Format(time.RFC3339)
}
