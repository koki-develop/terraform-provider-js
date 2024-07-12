package util

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	jstypes "github.com/koki-develop/terraform-provider-js/internal/types"
)

func StringifyValue(v attr.Value) string {
	if v.IsNull() {
		return "null"
	}

	switch v := v.(type) {
	case basetypes.StringValue:
		if strings.HasPrefix(v.ValueString(), jstypes.IDPrefix) {
			return strings.TrimPrefix(v.ValueString(), jstypes.IDPrefix)
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

type ModelGetter interface {
	Get(ctx context.Context, target any) diag.Diagnostics
}

type ModelSetter interface {
	Set(ctx context.Context, val any) diag.Diagnostics
}

func HandleRequest[T any](ctx context.Context, model T, g ModelGetter, s ModelSetter, diags diag.Diagnostics, h func(m T) error) {
	diags.Append(g.Get(ctx, model)...)
	if diags.HasError() {
		return
	}

	if err := h(model); err != nil {
		diags.AddError("failed to handle request", err.Error())
		return
	}

	diags.Append(s.Set(ctx, model)...)
	if diags.HasError() {
		return
	}
}
