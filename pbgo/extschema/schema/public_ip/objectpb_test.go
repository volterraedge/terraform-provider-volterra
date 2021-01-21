// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/public_ip/object.proto

/*
Package public_ip is a generated protocol buffer package.

It is generated from these files:
	ves.io/schema/public_ip/object.proto
	ves.io/schema/public_ip/public_crudapi.proto
	ves.io/schema/public_ip/types.proto

It has these top-level messages:
	Object
	SpecType
	StatusObject
	ReplaceRequest
	ReplaceResponse
	GetRequest
	GetResponse
	ListRequest
	ListResponseItem
	ListResponse
	GlobalSpecType
	ReplaceSpecType
	GetSpecType
*/
package public_ip

import testing "testing"
import rand "math/rand"
import time "time"
import proto "github.com/gogo/protobuf/proto"
import jsonpb "github.com/gogo/protobuf/jsonpb"
import fmt "fmt"
import parser "go/parser"
import golang_proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestObjectProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Object{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestObjectMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Object{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkObjectProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*Object, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedObject(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkObjectProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedObject(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &Object{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestSpecTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &SpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestSpecTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &SpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkSpecTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*SpecType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkSpecTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedSpecType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &SpecType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestStatusObjectProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StatusObject{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestStatusObjectMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StatusObject{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkStatusObjectProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*StatusObject, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedStatusObject(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkStatusObjectProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedStatusObject(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &StatusObject{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestObjectJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Object{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestSpecTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &SpecType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestStatusObjectJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StatusObject{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestObjectProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &Object{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestObjectProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &Object{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestSpecTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &SpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestSpecTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &SpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestStatusObjectProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &StatusObject{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestStatusObjectProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &StatusObject{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestObjectGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedObject(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestSpecTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSpecType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestStatusObjectGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedStatusObject(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestObjectSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedObject(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkObjectSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*Object, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedObject(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestSpecTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedSpecType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkSpecTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*SpecType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestStatusObjectSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStatusObject(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkStatusObjectSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*StatusObject, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedStatusObject(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestObjectStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedObject(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestSpecTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedSpecType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestStatusObjectStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedStatusObject(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
