// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"context"
	"strings"

	ves_io_schema_combined "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/combined"
)

var (
	// apiGWServiceSlugs is a mapping of public API FQN to the service slug to use on the api gateway
	// this will be replaced by discovery when api gateway supports it. The service slug is the part
	// at https://<xyz>.volterra.us/api/<service-slug>/<resource-uri>
	apiGWServiceSlugs = ves_io_schema_combined.APIGwServiceSlugs
)

// API gateway service slug e.g. "config"
type ctxServiceSlug struct{}

// contextFromServiceSlug sets the api gateway service discriminator in the context
func contextFromServiceSlug(ctx context.Context, slug string) context.Context {
	return context.WithValue(ctx, ctxServiceSlug{}, slug)
}

// addSvcSlugToContextForCRUD get slugs for given objType
func addSvcSlugToContextForCRUD(ctx context.Context, objType string) (context.Context, error) {
	mdReg := ves_io_schema_combined.MDR
	apiFQN := mdReg.GetCRUDAPIName(objType, true)
	if slug, exists := apiGWServiceSlugs[apiFQN]; exists {
		ctx = contextFromServiceSlug(ctx, slug)
	}
	return ctx, nil
}

// addSvcSlugToContextForCustom get slugs for given objType
func addSvcSlugToContextForCustom(ctx context.Context, rpcFQN string) (context.Context, error) {
	var pbSvc string
	rpcFQNList := strings.Split(rpcFQN, ".")
	pbSvc = strings.Join(rpcFQNList[:len(rpcFQNList)-1], ".")
	if slug, exists := apiGWServiceSlugs[pbSvc]; exists {
		ctx = contextFromServiceSlug(ctx, slug)
	}
	return ctx, nil
}
