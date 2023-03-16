// Copyright (c) 2023 F5 Inc. All rights reserved.
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
