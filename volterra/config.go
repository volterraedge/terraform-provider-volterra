//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/moul/http2curl"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"

	ves_io_schema_combined "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/combined"
)

// Config data for volterra provider
type Config struct {
	apiP12File     string
	apiP12Password string
	apiCert        string
	apiKey         string
	apiCACert      string
	url            string
	test           bool
	vesenv         bool
	tenant         string
	timeout        string
	// mapping of protobuf service FQN to service slug(e.g. 'config') on apiGw
	serviceSlugs map[string]string
}

// APIClient embeds vesapi.APIClient
type APIClient struct {
	*vesapi.APIClient
}

// CreateObject populates the correct slug and calls volt's CreateObject
func (a *APIClient) CreateObject(ctx context.Context, objType string, req vesapi.CreateObjectRequest, opts ...vesapi.CallOpt) (vesapi.CreateObjectResponse, error) {
	ctx, err := addSvcSlugToContextForCRUD(ctx, objType)
	if err != nil {
		return nil, err
	}
	return a.APIClient.CreateObject(ctx, objType, req, opts...)
}

// GetObject populates the correct slug and calls volt's GetObject
func (a *APIClient) GetObject(ctx context.Context, objType, namespace, name string, opts ...vesapi.CallOpt) (vesapi.GetObjectResponse, error) {
	ctx, err := addSvcSlugToContextForCRUD(ctx, objType)
	if err != nil {
		return nil, err
	}
	return a.APIClient.GetObject(ctx, objType, namespace, name, opts...)
}

// ReplaceObject populates the correct slug and calls volt's ReplaceObject
func (a *APIClient) ReplaceObject(ctx context.Context, objType string, req vesapi.ReplaceObjectRequest, opts ...vesapi.CallOpt) error {
	ctx, err := addSvcSlugToContextForCRUD(ctx, objType)
	if err != nil {
		return err
	}
	return a.APIClient.ReplaceObject(ctx, objType, req, opts...)

}

// DeleteObject populates the correct slug and calls volt's DeleteObject
func (a *APIClient) DeleteObject(ctx context.Context, objType string, namespace, name string, opts ...vesapi.CallOpt) error {
	ctx, err := addSvcSlugToContextForCRUD(ctx, objType)
	if err != nil {
		return err
	}
	return a.APIClient.DeleteObject(ctx, objType, namespace, name, opts...)
}

// ListObjects populates the correct slug and calls volt's ListObjects
func (a *APIClient) ListObjects(ctx context.Context, objType, namespace string, opts ...vesapi.CallOpt) ([]vesapi.ListObjectItem, error) {
	ctx, err := addSvcSlugToContextForCRUD(ctx, objType)
	if err != nil {
		return nil, err
	}
	return a.APIClient.ListObjects(ctx, objType, namespace, opts...)
}

// CustomAPI populates the correct slug and calls volt's custom API
func (a *APIClient) CustomAPI(ctx context.Context, method, uri, rpcFQN, reqYml string, opts ...vesapi.CallOpt) (proto.Message, error) {
	ctx, err := addSvcSlugToContextForCustom(ctx, rpcFQN)
	if err != nil {
		return nil, err
	}
	return a.APIClient.CustomAPI(ctx, method, uri, rpcFQN, reqYml, opts...)

}

// serviceSlugFromContext gets the api gateway service discriminator from the context
func serviceSlugFromContext(ctx context.Context) string {
	slug, ok := ctx.Value(ctxServiceSlug{}).(string)
	if !ok {
		return ""
	}
	return slug
}

// PrismURIRemapperFn is used as a http.RoundTripper to munge URIs as expected by
// Prism API gateway (for public APIs)
func prismURIRemapperFn(apiGwURL string) (client.RoundTripperFunc, error) {
	u, err := url.Parse(apiGwURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Parsing apigw url %s", apiGwURL)
	}
	rtFn := func(req *http.Request, rt http.RoundTripper) (*http.Response, error) {
		// convert https://x.y.volterra.us/api + /public/some/thing/here to
		//         https://x.y.volterra.us/api/config/some/thing/here
		srvURI := u.Path
		ctx := req.Context()
		slug := serviceSlugFromContext(ctx)
		if slug == "" {
			return nil, fmt.Errorf("No service slug defined for URI %s", srvURI)
		}
		toStrip := "/public"
		if len(srvURI) > 0 && srvURI[len(srvURI)-1] != '/' {
			// "/public" should be replaced by "/<slug>"
			slug = "/" + slug
		}
		uriPfx := fmt.Sprintf("%s%s", srvURI, toStrip)
		if strings.HasPrefix(req.URL.Path, uriPfx) {
			// replace "public" with "config"
			newURL := strings.Replace(req.URL.String(), toStrip, slug, 1)
			u, err := url.Parse(newURL)
			if err != nil {
				return nil, errors.Wrapf(err, "Parsing remapped url %s", newURL)
			}
			req.URL = u
		}

		curlOut, err := http2curl.GetCurlCommand(req)
		if err == nil {
			log.Printf("[TRACE] Request Sent: \n%s\n", curlOut)
		}

		rsp, err := rt.RoundTrip(req)
		if err != nil {
			log.Printf("[TRACE] Response Error: \n%s\n", err.Error())
		} else {
			if out, err := httputil.DumpResponse(rsp, true); err == nil {
				log.Printf("[TRACE] Response Ouptut:\n%s\n", out)
			}
		}
		return rsp, err
	}
	return rtFn, nil
}

// uriRemapperFn is used as a http.RoundTripper to munge URIs as expected by regular services
func uriRemapperFn(apiGwURL, tenant string) client.RoundTripperFunc {
	rtFn := func(req *http.Request, rt http.RoundTripper) (*http.Response, error) {
		req.Header.Set(server.TenantHdrName, tenant)
		curlOut, err := http2curl.GetCurlCommand(req)
		if err == nil {
			log.Printf("[TRACE] Request Sent: \n%s\n", curlOut)
		}

		rsp, err := rt.RoundTrip(req)
		if err != nil {
			log.Printf("[TRACE] Response Error: \n%s\n", err.Error())
		} else {
			if out, err := httputil.DumpResponse(rsp, true); err == nil {
				log.Printf("[TRACE] Response Ouptut:\n%s\n", out)
			}
		}
		return rsp, err
	}
	return rtFn
}

func getclientOpts(c *Config) ([]vesapi.ConfigOpt, error) {
	remapperFn, err := prismURIRemapperFn(c.url)
	if err != nil {
		return nil, err
	}
	if c.vesenv {
		remapperFn = uriRemapperFn(c.url, c.tenant)
	}

	dur, err := time.ParseDuration(c.timeout)
	if err != nil {
		return nil, err
	}
	clOpts := []vesapi.ConfigOpt{
		vesapi.WithCfgRoundTripperFn(remapperFn),
		vesapi.WithCfgRPCTimeout(dur),
	}

	if c.apiP12File != "" {
		if _, err := os.Stat(c.apiP12File); err != nil {
			clOpts = append(clOpts, vesapi.WithCfgP12Bundle(c.apiP12File, c.apiP12Password))
		}
	}

	if c.apiCert != "" && c.apiKey != "" {
		clOpts = append(clOpts, vesapi.WithCfgCert(c.apiCert))
		clOpts = append(clOpts, vesapi.WithCfgKey(c.apiKey))
	}

	if c.apiCACert != "" {
		clOpts = append(clOpts, vesapi.WithCfgCACert(c.apiCACert))
	}

	return clOpts, nil
}

// Client returns a new volt api client, which will be used to access public API's
func (c *Config) Client() (*APIClient, error) {

	var clOpts []vesapi.ConfigOpt
	clOpts, err := getclientOpts(c)
	if err != nil {
		return nil, errors.Wrap(err, "Building client options")
	}
	if c.test {
		clOpts, err = getTestClientOpts(c.url, c.tenant)
		if err != nil {
			return nil, errors.Wrap(err, "Building client options")
		}
	}

	voltAPICl, err := vesapi.NewAPIClient(c.url, ves_io_schema_combined.MDR, clOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new Volterra APIClient")
	}

	apiCl := &APIClient{
		APIClient: voltAPICl,
	}
	return apiCl, nil
}
