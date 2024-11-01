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
	_ datasource.DataSource = &dataReturn{}
)

func NewDataReturn() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataReturn{}
	}
}

type dataReturn struct{}

func (d *dataReturn) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_return"
}

func (d *dataReturn) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_return` data source creates a return.",
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "The value of the return.",
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the return.",
				Computed:    true,
			},
		},
	}
}

type dataReturnModel struct {
	Value   types.Dynamic `tfsdk:"value"`
	Content types.String  `tfsdk:"content"`
}

func (d *dataReturn) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataReturnModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataReturnModel) bool {
			c := new(strings.Builder)
			c.WriteString("return")
			if !m.Value.IsNull() {
				c.WriteString(" ")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
