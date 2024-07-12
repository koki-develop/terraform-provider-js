package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
)

func stringifyValue(v attr.Value) string {
	if v.IsNull() {
		return "null"
	}

	switch v := v.(type) {
	case IDValue:
		return v.ValueString()
	}

	return v.String()
}

func stringifyValues(vs []attr.Value) []string {
	strs := make([]string, len(vs))
	for i, v := range vs {
		strs[i] = stringifyValue(v)
	}

	return strs
}
