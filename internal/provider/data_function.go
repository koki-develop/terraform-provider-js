package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = (*dataFunction)(nil)
)

type dataFunction struct{}

func NewDataFunction() datasource.DataSource {
	return &dataFunction{}
}

func (d *dataFunction) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_function", req.ProviderTypeName)
}

func (d *dataFunction) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},

			"id": schema.StringAttribute{
				CustomType: ID{},
				Computed:   true,
			},
		},
	}
}

type dataFunctionModel struct {
	ID   IDValue      `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *dataFunction) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var attrs dataFunctionModel
	diags := req.Config.Get(ctx, &attrs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	attrs.ID = IDValue{
		StringValue: types.StringValue(attrs.Name.ValueString()),
	}
	diags = resp.State.Set(ctx, &attrs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
