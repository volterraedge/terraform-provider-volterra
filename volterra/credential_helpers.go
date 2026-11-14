// Copyright (c) 2026 F5 Inc. All rights reserved.

package volterra

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/proto"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

type credentialCustomAPIClient interface {
	CustomAPI(context.Context, string, string, string, string, ...vesapi.CallOpt) (proto.Message, error)
}

func isCredentialNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errString := strings.ToLower(err.Error())
	return strings.Contains(errString, "status code 404") || strings.Contains(errString, "object not found")
}
