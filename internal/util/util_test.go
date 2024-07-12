package util

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func Test_StringifyValue(t *testing.T) {
	tests := []struct {
		v    attr.Value
		json bool
		want string
	}{
		// id
		{v: types.StringValue("@js:id:foo"), want: "foo"},

		// string
		{v: types.StringValue("foo"), want: `"foo"`},
		{v: types.StringValue(""), want: `""`},
		{v: types.StringNull(), want: "null"},
		{v: types.DynamicValue(types.StringValue("foo")), want: `"foo"`},

		// number
		{v: types.NumberValue(big.NewFloat(1)), want: "1"},
		{v: types.NumberValue(big.NewFloat(0)), want: "0"},
		{v: types.NumberNull(), want: "null"},
		{v: types.DynamicValue(types.NumberValue(big.NewFloat(1))), want: "1"},

		// bool
		{v: types.BoolValue(true), want: "true"},
		{v: types.BoolValue(false), want: "false"},
		{v: types.BoolNull(), want: "null"},
		{v: types.DynamicValue(types.BoolValue(true)), want: "true"},

		// list
		{
			v:    types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}),
			want: `["foo","bar"]`,
		},
		{
			v:    types.ListValueMust(types.NumberType, []attr.Value{types.NumberValue(big.NewFloat(1)), types.NumberValue(big.NewFloat(2))}),
			want: `[1,2]`,
		},

		// tuple
		{
			v: types.TupleValueMust(
				[]attr.Type{
					types.StringType,
					types.NumberType,
					types.BoolType,
					types.TupleType{ElemTypes: []attr.Type{types.NumberType, types.NumberType}},
				},
				[]attr.Value{
					types.StringValue("foo"),
					types.NumberValue(big.NewFloat(1)),
					types.BoolValue(true),
					types.TupleValueMust([]attr.Type{types.NumberType, types.NumberType}, []attr.Value{types.NumberValue(big.NewFloat(1)), types.NumberValue(big.NewFloat(2))}),
				},
			),
			want: `["foo",1,true,[1,2]]`,
		},

		// object
		{
			json: true,
			v: types.ObjectValueMust(
				map[string]attr.Type{
					"string": types.StringType,
					"number": types.NumberType,
					"bool":   types.BoolType,
					"tuple":  types.TupleType{ElemTypes: []attr.Type{types.NumberType, types.NumberType}},
					"object": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"string": types.StringType,
						},
					},
				},
				map[string]attr.Value{
					"string": types.StringValue("foo"),
					"number": types.NumberValue(big.NewFloat(1)),
					"bool":   types.BoolValue(true),
					"tuple":  types.TupleValueMust([]attr.Type{types.NumberType, types.NumberType}, []attr.Value{types.NumberValue(big.NewFloat(1)), types.NumberValue(big.NewFloat(2))}),
					"object": types.ObjectValueMust(
						map[string]attr.Type{
							"string": types.StringType,
						},
						map[string]attr.Value{
							"string": types.StringValue("foo"),
						},
					),
				},
			),
			want: `{"string":"foo","number":1,"bool":true,"tuple":[1,2],"object":{"string":"foo"}}`,
		},
		{v: types.ObjectNull(map[string]attr.Type{}), want: "null"},

		// map
		{
			json: true,
			v: types.MapValueMust(
				types.StringType,
				map[string]attr.Value{
					"foo": types.StringValue("bar"),
					"baz": types.StringValue("qux"),
				},
			),
			want: `{"foo":"bar","baz":"qux"}`,
		},
		{v: types.MapNull(types.StringType), want: "null"},

		// other
		{v: types.DynamicNull(), want: "null"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := StringifyValue(tt.v)
			if tt.json {
				assert.JSONEq(t, tt.want, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}