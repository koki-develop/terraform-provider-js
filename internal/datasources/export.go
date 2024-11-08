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
	_ datasource.DataSource = &dataExport{}
)

func NewDataExport() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataExport{}
	}
}

type dataExport struct{}

func (d *dataExport) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_export"
}

func (d *dataExport) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "Value to be exported",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the export.",
				Computed:    true,
			},
		},
	}
}

type dataExportModel struct {
	Value   types.Dynamic `tfsdk:"value"`
	Content types.String  `tfsdk:"content"`
}

func (d *dataExport) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataExportModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataExportModel) bool {
			c := new(strings.Builder)
			c.WriteString("export ")
			c.WriteString(util.StringifyValue(m.Value))

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
