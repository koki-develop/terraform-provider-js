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
	_ datasource.DataSource = &dataConditionalOperation{}
)

func NewDataConditionalOperation() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataConditionalOperation{}
	}
}

type dataConditionalOperation struct{}

func (d *dataConditionalOperation) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_conditional_operation"
}

func (d *dataConditionalOperation) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"condition": schema.DynamicAttribute{
				Description: "An expression whose value is used as a condition. ",
				Required:    true,
			},
			"if_true": schema.DynamicAttribute{
				Description: "An expression which is executed if the condition evaluates to a truthy value.",
				Required:    true,
			},
			"if_false": schema.DynamicAttribute{
				Description: "An expression which is executed if the condition is falsy.",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataConditionalOperationModel struct {
	Condition types.Dynamic `tfsdk:"condition"`
	IfTrue    types.Dynamic `tfsdk:"if_true"`
	IfFalse   types.Dynamic `tfsdk:"if_false"`

	Content types.String `tfsdk:"content"`
}

func (d *dataConditionalOperation) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataConditionalOperationModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataConditionalOperationModel) bool {
			c := new(strings.Builder)
			c.WriteString("(")
			c.WriteString(util.StringifyValue(m.Condition))
			c.WriteString("?")
			c.WriteString(util.StringifyValue(m.IfTrue))
			c.WriteString(":")
			c.WriteString(util.StringifyValue(m.IfFalse))
			c.WriteString(")")

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
