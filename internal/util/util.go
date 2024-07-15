package util

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const prefix = "@js/"
const RawPrefix = prefix + "raw:"

func Raw(v basetypes.StringValue) basetypes.StringValue {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	if hasRawPrefix(v) {
		return v
	}

	return types.StringValue(RawPrefix + v.ValueString())
}

func RawString(v basetypes.StringValue) string {
	return strings.TrimPrefix(v.ValueString(), RawPrefix)
}

func hasRawPrefix(v basetypes.StringValue) bool {
	return strings.HasPrefix(v.ValueString(), RawPrefix)
}

func StringifyValue(v attr.Value) string {
	if v.IsNull() {
		return "null"
	}

	switch v := v.(type) {
	case basetypes.StringValue:
		if hasRawPrefix(v) {
			return RawString(v)
		}
	case basetypes.DynamicValue:
		return StringifyValue(v.UnderlyingValue())
	}

	return v.String()
}

func StringifyValues(vs []attr.Value) []string {
	strs := make([]string, len(vs))
	for i, v := range vs {
		strs[i] = StringifyValue(v)
	}

	return strs
}

func StringifyStatements(vs []attr.Value) string {
	c := new(strings.Builder)

	for i, v := range vs {
		ch := c.String()[c.Len()-1:]
		if i > 0 && ch != "}" && ch != ";" {
			c.WriteString(";")
		}
		c.WriteString(StringifyValue(v))
	}

	return c.String()
}

type ModelGetter interface {
	Get(ctx context.Context, target any) diag.Diagnostics
}

type ModelSetter interface {
	Set(ctx context.Context, val any) diag.Diagnostics
}

func HandleRequest[T any](ctx context.Context, model T, g ModelGetter, s ModelSetter, diags *diag.Diagnostics, h func(m T) bool) {
	diags.Append(g.Get(ctx, model)...)
	if diags.HasError() {
		return
	}

	if !h(model) {
		return
	}

	diags.Append(s.Set(ctx, model)...)
	if diags.HasError() {
		return
	}
}
