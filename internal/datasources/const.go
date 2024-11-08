package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataConst{}
)

func NewDataConst() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataConst{}
	}
}

type dataConst struct{}

func (d *dataConst) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const"
}

func (d *dataConst) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Name of variable to declare.",
				Required:    true,
			},
			"value": schema.DynamicAttribute{
				Description: "Initial value of the variable.",
				Required:    true,
			},

			"id": schema.StringAttribute{
				Computed: true,
			},
			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataConstModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (d *dataConst) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataConstModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataConstModel) bool {
			m.ID = util.Raw(m.Name)
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("const %s=%s", util.RawString(m.Name), util.StringifyValue(m.Value))))
			return true
		},
	)
}
