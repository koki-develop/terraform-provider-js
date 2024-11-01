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
	_ datasource.DataSource = &dataNew{}
)

func NewDataNew() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataNew{}
	}
}

type dataNew struct{}

func (d *dataNew) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_new"
}

func (d *dataNew) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_new` data source creates a new operation.",
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "The value of the operation.",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the operation.",
				Computed:    true,
			},
		},
	}
}

type dataNewModel struct {
	Value   types.String `tfsdk:"value"`
	Content types.String `tfsdk:"content"`
}

func (d *dataNew) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataNewModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataNewModel) bool {
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("new %s", util.StringifyValue(m.Value))))
			return true
		},
	)
}
