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
	_ datasource.DataSource = &dataFor{}
)

func NewDataFor() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataFor{}
	}
}

type dataFor struct{}

func (d *dataFor) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_for"
}

func (d *dataFor) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_for` resource creates a for loop.",
		Attributes: map[string]schema.Attribute{
			"init": schema.StringAttribute{
				Description: "The initialization expression.",
				Optional:    true,
			},
			"condition": schema.StringAttribute{
				Description: "The condition expression.",
				Optional:    true,
			},
			"update": schema.StringAttribute{
				Description: "The update expression.",
				Optional:    true,
			},
			"body": schema.ListAttribute{
				Description: "The body of the for loop.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the for loop.",
				Computed:    true,
			},
		},
	}
}

type resourceForModel struct {
	Init      types.String `tfsdk:"init"`
	Condition types.String `tfsdk:"condition"`
	Update    types.String `tfsdk:"update"`
	Body      types.List   `tfsdk:"body"`

	Content types.String `tfsdk:"content"`
}

func (d *dataFor) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&resourceForModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *resourceForModel) bool {
			c := new(strings.Builder)
			c.WriteString("for(")
			if !m.Init.IsNull() {
				c.WriteString(util.StringifyValue(m.Init))
			}
			c.WriteString(";")

			if !m.Condition.IsNull() {
				c.WriteString(util.StringifyValue(m.Condition))
			}
			c.WriteString(";")

			if !m.Update.IsNull() {
				c.WriteString(util.StringifyValue(m.Update))
			}

			c.WriteString("){")

			if !m.Body.IsNull() {
				c.WriteString(util.StringifyStatements(m.Body.Elements()))
			}
			c.WriteString("}")

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
