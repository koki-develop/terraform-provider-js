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
		Attributes: map[string]schema.Attribute{
			"init": schema.StringAttribute{
				Description: "An expression or variable declaration evaluated once before the loop begins.",
				Optional:    true,
			},
			"condition": schema.StringAttribute{
				Description: "An expression to be evaluated before each loop iteration.",
				Optional:    true,
			},
			"update": schema.StringAttribute{
				Description: "An expression to be evaluated at the end of each loop iteration.",
				Optional:    true,
			},
			"body": schema.ListAttribute{
				Description: "Statements that is executed as long as the condition evaluates to true.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataForModel struct {
	Init      types.String `tfsdk:"init"`
	Condition types.String `tfsdk:"condition"`
	Update    types.String `tfsdk:"update"`
	Body      types.List   `tfsdk:"body"`

	Content types.String `tfsdk:"content"`
}

func (d *dataFor) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataForModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataForModel) bool {
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
