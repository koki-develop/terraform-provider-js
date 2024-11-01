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
	_ datasource.DataSource = &dataVar{}
)

func NewDataVar() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataVar{}
	}
}

type dataVar struct{}

func (d *dataVar) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_var"
}

func (d *dataVar) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_var` data source defines a variable.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the variable.",
				Required:    true,
			},
			"value": schema.DynamicAttribute{
				Description: "The value of the variable.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "The id of the variable.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the variable.",
				Computed:    true,
			},
		},
	}
}

type dataVarModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (d *dataVar) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataVarModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataVarModel) bool {
			m.ID = util.Raw(m.Name)

			c := new(strings.Builder)
			c.WriteString("var ")
			c.WriteString(util.RawString(m.Name))
			if !m.Value.IsNull() {
				c.WriteString("=")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
