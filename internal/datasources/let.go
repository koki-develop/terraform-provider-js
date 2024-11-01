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
	_ datasource.DataSource = &dataLet{}
)

func NewDataLet() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataLet{}
	}
}

type dataLet struct{}

func (d *dataLet) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_let"
}

func (d *dataLet) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_let` data source defines a let statement.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the let statement.",
				Required:    true,
			},
			"value": schema.DynamicAttribute{
				Description: "The value of the let statement.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "The id of the let statement.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the let statement.",
				Computed:    true,
			},
		},
	}
}

type dataLetModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (d *dataLet) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataLetModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataLetModel) bool {
			m.ID = util.Raw(m.Name)

			c := new(strings.Builder)
			c.WriteString("let ")
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
