package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	return IDValue{StringValue: in}, nil
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

type IDValue struct {
	basetypes.StringValue
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
