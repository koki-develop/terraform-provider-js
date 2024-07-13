package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = (*dataRaw)(nil)
)

type dataRaw struct{}

func NewDataRaw() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataRaw{}
	}
}

func (d *dataRaw) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_raw"
}

func (d *dataRaw) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Required: true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataFunctionModel struct {
	Value types.String `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (d *dataRaw) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataFunctionModel{},
		&req.Config,
		&resp.State,
		resp.Diagnostics,
		func(m *dataFunctionModel) error {
			m.Content = util.Raw(m.Value)
			return nil
		},
	)
}
