package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataSourceAwait{}
)

func NewDataAwait() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataSourceAwait{}
	}
}

type dataSourceAwait struct{}

func (d *dataSourceAwait) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_await"
}

func (d *dataSourceAwait) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_await` data source defines an await.",
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "The value of the await.",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the await.",
				Computed:    true,
			},
		},
	}
}

type dataSourceAwaitModel struct {
	Value types.Dynamic `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (d *dataSourceAwait) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataSourceAwaitModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataSourceAwaitModel) bool {
			c := new(strings.Builder)
			c.WriteString("await ")
			c.WriteString(util.StringifyValue(m.Value))

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
