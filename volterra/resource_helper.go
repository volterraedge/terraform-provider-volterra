//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"log"
	"reflect"
	"strings"
	"time"

	google_protobuf1 "github.com/gogo/protobuf/types"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func getResourceData(d *schema.ResourceData, key string) interface{} {
	if strings.Contains(key, ".") {
		if v, ok := d.GetOk(key); ok {
			return v
		}
		log.Printf("[DEBUG] %s key is not set by user", key)
		return nil
	}
	return d.Get(key)
}

func valueIsZero(v reflect.Value, kind reflect.Kind) bool {
	switch kind {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func isIntfNil(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Ptr:
		foundVal := v.(*schema.Set)
		if foundVal.Len() == 0 {
			return true
		}

	default:
		value := reflect.ValueOf(v)
		if valueIsZero(value, value.Kind()) {
			return true
		}
	}
	return false
}

func parseTime(t string) (*google_protobuf1.Timestamp, error) {
	layout := "0001-01-01T00:00:00Z"
	timeFmt, err := time.Parse(layout, t)
	if err != nil {
		return nil, err
	}
	return google_protobuf1.TimestampProto(timeFmt)
}
