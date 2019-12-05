package generator

import (
	"bytes"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Oneof", func() {

	DescribeTable("processOneofFields",
		func(dataList []*Data, expected string) {
			buf := []byte{}
			w := bytes.NewBuffer(buf)
			err := processOneofFields(w, dataList)
			Expect(err).NotTo(HaveOccurred())

			Expect(w.String()).To(Equal(expected))
		},

		Entry("Empty data", nil, ""),

		Entry("Empty field list", []*Data{
			{
				SrcPref:       "src_pref",
				Src:           "src",
				SrcFn:         "src_fn",
				SrcPointer:    "src_pointer",
				DstPref:       "dst_pref",
				Dst:           "dst",
				DstFn:         "dst_fn",
				DstPointer:    "dst_pointer",
				Swapped:       false,
				HelperPackage: "hp",
				Ptr:           false,
				Fields:        nil,
			},
		}, ""),

		Entry("Single field", []*Data{
			&Data{
				SrcPref:       "src_pref",
				Src:           "src",
				SrcFn:         "src_fn",
				SrcPointer:    "src_pointer",
				DstPref:       "dst_pref",
				Dst:           "dst",
				DstFn:         "dst_fn",
				DstPointer:    "dst_pointer",
				Swapped:       false,
				HelperPackage: "hp",
				Ptr:           false,
				Fields: []Field{
					{
						Name:          "GoField",
						ProtoType:     "pt",
						GoToProtoType: "gtTopt",
						OneofDecl:     "decl_name",
					},
				},
			},
		}, singleField),

		Entry("Two fields, 2nd is oneof", []*Data{
			&Data{
				SrcPref:       "src_pref",
				Src:           "src",
				SrcFn:         "src_fn",
				SrcPointer:    "src_pointer",
				DstPref:       "dst_pref",
				Dst:           "dst",
				DstFn:         "dst_fn",
				DstPointer:    "dst_pointer",
				Swapped:       false,
				HelperPackage: "hp",
				Ptr:           false,
				Fields: []Field{
					{
						Name:          "FirstField",
						ProtoType:     "pt",
						GoToProtoType: "gtTopt",
					},
					{
						Name:          "SecondField",
						ProtoType:     "pt",
						GoToProtoType: "gtTopt",
						OneofDecl:     "decl_name",
					},
				},
			},
		}, twoFields),
	)

	BeforeEach(func() {
		version = "v1.1.1"
		buildTime = time.Date(2019, time.March, 1, 5, 34, 19, 0, time.UTC).Format(time.RFC3339)
	})

	DescribeTable("OptHelpers",
		func(name, expected string) {
			r := OptHelpers(name)
			Expect(r).To(Equal(expected))
		},
		Entry("Package One", "one", headerOne),
	)

})

var (
	singleField = `
type OneofDeclName interface {
	GetStringValue() string
	GetInt64Value() int64
}

func ptTogt(src OneofDeclName) string {
	if s := src.GetStringValue(); s != "" {
		return s
	}

	if i := src.GetInt64Value(); i != 0 {
		return strconv.FormatInt(i, 10)
	}

	return "<nil>"
}

func gtTopt(s string, dst *dst_pref.pt, v string) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil  || v == "v2"{
		dst.DeclName = &dst_pref.pt_StringValue{StringValue: s}
		return
	}

	dst.DeclName = &dst_pref.pt_Int64Value{Int64Value: i}
	return
}

`

	twoFields = `
type OneofDeclName interface {
	GetStringValue() string
	GetInt64Value() int64
}

func ptTogt(src OneofDeclName) string {
	if s := src.GetStringValue(); s != "" {
		return s
	}

	if i := src.GetInt64Value(); i != 0 {
		return strconv.FormatInt(i, 10)
	}

	return "<nil>"
}

func gtTopt(s string, dst *dst_pref.pt, v string) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil  || v == "v2"{
		dst.DeclName = &dst_pref.pt_StringValue{StringValue: s}
		return
	}

	dst.DeclName = &dst_pref.pt_Int64Value{Int64Value: i}
	return
}

`

	headerOne = `// Code generated by protoc-gen-struct-transformer, version: v1.1.1. DO NOT EDIT.

package one
var version string

// TransformParam is a function option type.
type TransformParam func()

// WithVersion sets global version variable.
func WithVersion(v string) TransformParam {
	return func() {
		version = v
	}
}

func applyOptions(opts ...TransformParam) {
	for _, o := range opts {
		o()
	}
}


`
)
