package types

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = ID{}

type ID struct {
	basetypes.StringType
}

func (i ID) Equal(o attr.Type) bool {
	other, ok := o.(ID)
	if !ok {
		return false
	}

	return i.StringType.Equal(other.StringType)
}

func (i ID) String() string {
	return "ID"
}

func (i ID) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return NewIDValue(in), nil
}

func (i ID) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := i.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := i.ValueFromString(ctx, stringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (i ID) ValueType(ctx context.Context) attr.Value {
	return IDValue{}
}

var _ basetypes.StringValuable = IDValue{}
var _ basetypes.StringValuableWithSemanticEquals = IDValue{}

const IDPrefix = Prefix + "id:"

type IDValue struct {
	basetypes.StringValue
}

func NewIDValue(s basetypes.StringValue) IDValue {
	if s.IsNull() || s.IsUnknown() {
		return IDValue{StringValue: s}
	}
	if strings.HasPrefix(s.ValueString(), IDPrefix) {
		return IDValue{StringValue: s}
	}

	return IDValue{StringValue: types.StringValue(fmt.Sprintf("%s%s", IDPrefix, s.ValueString()))}
}

func (v IDValue) ValueString() string {
	return strings.TrimPrefix(v.StringValue.ValueString(), IDPrefix)
}

func (v IDValue) Equal(o attr.Value) bool {
	other, ok := o.(IDValue)
	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v IDValue) Type(ctx context.Context) attr.Type {
	return ID{}
}

func (v IDValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IDValue)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	return strings.TrimPrefix(v.ValueString(), IDPrefix) == strings.TrimPrefix(newValue.ValueString(), IDPrefix), diags
}
