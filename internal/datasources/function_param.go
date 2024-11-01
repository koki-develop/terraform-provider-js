package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataFunctionParam{}
)

func NewDataFunctionParam() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataFunctionParam{}
	}
}

type dataFunctionParam struct{}

func (d *dataFunctionParam) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_param"
}

func (d *dataFunctionParam) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_function_param` data source defines a function parameter.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the function parameter.",
				Required:    true,
			},

			"id": schema.StringAttribute{
				Description: "The id of the function parameter.",
				Computed:    true,
			},
		},
	}
}

type dataFunctionParamModel struct {
	Name types.String `tfsdk:"name"`
	ID   types.String `tfsdk:"id"`
}

func (d *dataFunctionParam) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataFunctionParamModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataFunctionParamModel) bool {
			m.ID = util.Raw(m.Name)
			return true
		},
	)
}
