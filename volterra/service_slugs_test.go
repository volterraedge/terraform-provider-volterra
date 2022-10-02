//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSvcSlugForCustomAPI(t *testing.T) {
	ctx, err := addSvcSlugToContextForCustom(context.Background(), "ves.io.schema.views.terraform_parameters.CustomAPI.GetStatus")
	assert.NoError(t, err)
	v := ctx.Value(ctxServiceSlug{})
	assert.NotNil(t, v)
	assert.Equal(t, "config", v)
}
