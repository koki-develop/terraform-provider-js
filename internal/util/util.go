package util

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	jstypes "github.com/koki-develop/terraform-provider-js/internal/types"
)

func StringifyValue(v attr.Value) string {
	if v.IsNull() {
		return "null"
	}

	switch v := v.(type) {
	case jstypes.IDValue:
		return v.ValueString()
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
