// Copyright (c) 2026 F5 Inc. All rights reserved.

package volterra

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/svcfw"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
)

type credentialAPICall struct {
	method  string
	uri     string
	request string
}

type fakeCredentialAPIClient struct {
	calls     []credentialAPICall
	errors    map[string]error
	responses map[string]proto.Message
}

func (f *fakeCredentialAPIClient) CustomAPI(_ context.Context, method, uri, _, request string, _ ...vesapi.CallOpt) (proto.Message, error) {
	f.calls = append(f.calls, credentialAPICall{method: method, uri: uri, request: request})
	return f.responses[uri], f.errors[uri]
}

func TestIsCredentialNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{name: "http 404", err: errors.New("status code 404"), expected: true},
		{name: "backend object not found", err: errors.New("status code 500: api credential entries: object not found"), expected: true},
		{name: "capitalized backend object not found", err: errors.New("status code 500: Object Not Found"), expected: true},
		{name: "other server error", err: errors.New("status code 500: internal error"), expected: false},
		{name: "nil error", err: nil, expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if actual := isCredentialNotFoundError(test.err); actual != test.expected {
				t.Fatalf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}

func TestServiceCredentialReadNamespaces(t *testing.T) {
	tests := []struct {
		name     string
		config   map[string]interface{}
		expected []string
	}{
		{
			name: "vk8s kubeconfig uses vk8s namespace with system fallback",
			config: map[string]interface{}{
				"vk8s_kubeconfig": []interface{}{
					map[string]interface{}{"vk8s_namespace": "app-ns"},
				},
			},
			expected: []string{"app-ns", svcfw.SystemNSVal},
		},
		{
			name: "site kubeconfig remains system scoped",
			config: map[string]interface{}{
				"site_kubeconfig": []interface{}{
					map[string]interface{}{"site": "site-1"},
				},
			},
			expected: []string{svcfw.SystemNSVal},
		},
		{
			name:     "missing credential choice remains system scoped",
			config:   map[string]interface{}{},
			expected: []string{svcfw.SystemNSVal},
		},
		{
			name: "vk8s kubeconfig without namespace remains system scoped",
			config: map[string]interface{}{
				"vk8s_kubeconfig": []interface{}{
					map[string]interface{}{"vk8s_cluster_name": "cluster-1"},
				},
			},
			expected: []string{svcfw.SystemNSVal},
		},
		{
			name: "vk8s kubeconfig with system namespace does not duplicate fallback",
			config: map[string]interface{}{
				"vk8s_kubeconfig": []interface{}{
					map[string]interface{}{"vk8s_namespace": svcfw.SystemNSVal},
				},
			},
			expected: []string{svcfw.SystemNSVal},
		},
		{
			name: "service api token remains system scoped",
			config: map[string]interface{}{
				"service_api_token": true,
			},
			expected: []string{svcfw.SystemNSVal},
		},
		{
			name: "service api certificate remains system scoped",
			config: map[string]interface{}{
				"service_api_certificate": []interface{}{
					map[string]interface{}{"password": "test-password"},
				},
			},
			expected: []string{svcfw.SystemNSVal},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := schema.TestResourceDataRaw(t, resourceVolterraServiceCredential().Schema, test.config)
			if actual := serviceCredentialReadNamespaces(data); !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected namespaces %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestServiceCredentialVK8sKubeconfigAllowsOneBlock(t *testing.T) {
	if maxItems := resourceVolterraServiceCredential().Schema["vk8s_kubeconfig"].MaxItems; maxItems != 1 {
		t.Fatalf("expected vk8s_kubeconfig MaxItems to be 1, got %d", maxItems)
	}
}

func TestReplaceServiceCredentialUsesVK8sNamespace(t *testing.T) {
	data := schema.TestResourceDataRaw(t, resourceVolterraServiceCredential().Schema, map[string]interface{}{
		"vk8s_kubeconfig": []interface{}{
			map[string]interface{}{"vk8s_namespace": "app-ns"},
		},
	})
	data.SetId("credential-1")
	client := &fakeCredentialAPIClient{
		errors:    map[string]error{},
		responses: map[string]proto.Message{},
	}
	req := &ves_io_schema_api_credential.ReplaceServiceCredentialsRequest{Name: data.Id()}

	if err := replaceServiceCredential(client, data, req); err != nil {
		t.Fatalf("replace returned an error: %v", err)
	}
	if len(client.calls) != 1 {
		t.Fatalf("expected one API call, got %d", len(client.calls))
	}
	call := client.calls[0]
	if call.method != http.MethodPut || call.uri != "/public/namespaces/app-ns/service_credentials/credential-1" {
		t.Fatalf("unexpected replacement call: %s %s", call.method, call.uri)
	}
	if !strings.Contains(call.request, "namespace: app-ns") {
		t.Fatalf("expected request namespace app-ns, got:\n%s", call.request)
	}
}

func TestReplaceServiceCredentialFallsBackToSystem(t *testing.T) {
	data := schema.TestResourceDataRaw(t, resourceVolterraServiceCredential().Schema, map[string]interface{}{
		"vk8s_kubeconfig": []interface{}{
			map[string]interface{}{"vk8s_namespace": "app-ns"},
		},
	})
	data.SetId("credential-1")
	client := &fakeCredentialAPIClient{
		errors: map[string]error{
			"/public/namespaces/app-ns/service_credentials/credential-1": errors.New("object not found"),
		},
		responses: map[string]proto.Message{},
	}
	req := &ves_io_schema_api_credential.ReplaceServiceCredentialsRequest{Name: data.Id()}

	if err := replaceServiceCredential(client, data, req); err != nil {
		t.Fatalf("replace returned an error: %v", err)
	}
	if len(client.calls) != 2 {
		t.Fatalf("expected two API calls, got %d", len(client.calls))
	}
	if uri := client.calls[1].uri; uri != "/public/namespaces/system/service_credentials/credential-1" {
		t.Fatalf("expected system fallback, got %s", uri)
	}
}

func TestReplaceServiceCredentialReturnsUnexpectedError(t *testing.T) {
	data := schema.TestResourceDataRaw(t, resourceVolterraServiceCredential().Schema, map[string]interface{}{
		"vk8s_kubeconfig": []interface{}{
			map[string]interface{}{"vk8s_namespace": "app-ns"},
		},
	})
	data.SetId("credential-1")
	expectedErr := errors.New("internal server error")
	client := &fakeCredentialAPIClient{
		errors: map[string]error{
			"/public/namespaces/app-ns/service_credentials/credential-1": expectedErr,
		},
		responses: map[string]proto.Message{},
	}
	req := &ves_io_schema_api_credential.ReplaceServiceCredentialsRequest{Name: data.Id()}

	if err := replaceServiceCredential(client, data, req); !errors.Is(err, expectedErr) {
		t.Fatalf("expected replacement error %v, got %v", expectedErr, err)
	}
	if len(client.calls) != 1 {
		t.Fatalf("expected replacement to stop after one API call, got %d", len(client.calls))
	}
}

func TestServiceCredentialReadClearsStateForBackendNotFound(t *testing.T) {
	data := schema.TestResourceDataRaw(t, resourceVolterraServiceCredential().Schema, map[string]interface{}{
		"vk8s_kubeconfig": []interface{}{
			map[string]interface{}{"vk8s_namespace": "app-ns"},
		},
	})
	data.SetId("credential-1")
	client := &fakeCredentialAPIClient{
		errors: map[string]error{
			"/public/namespaces/app-ns/service_credentials/credential-1": errors.New("object not found"),
			"/public/namespaces/system/service_credentials/credential-1": errors.New("object not found"),
		},
		responses: map[string]proto.Message{},
	}

	if err := resourceVolterraServiceCredentialRead(data, client); err != nil {
		t.Fatalf("read returned an error: %v", err)
	}
	if data.Id() != "" {
		t.Fatalf("expected missing credential to be removed from state, got ID %q", data.Id())
	}
}
