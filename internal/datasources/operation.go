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
	_ datasource.DataSource = &dataOperation{}
)

func NewDataOperation() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataOperation{}
	}
}

type dataOperation struct{}

func (d *dataOperation) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operation"
}

func (d *dataOperation) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"operator": schema.StringAttribute{
				Description: "Operator to use in operation.",
				Required:    true,
			},
			"left": schema.DynamicAttribute{
				Description: "Left operand.",
				Required:    true,
			},
			"right": schema.DynamicAttribute{
				Description: "Right operand.",
				Required:    true,
			},

			"expression": schema.StringAttribute{
				Computed: true,
			},
			"statement": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataOperationModel struct {
	Operator types.String  `tfsdk:"operator"`
	Left     types.Dynamic `tfsdk:"left"`
	Right    types.Dynamic `tfsdk:"right"`

	Expression types.String `tfsdk:"expression"`
	Statement  types.String `tfsdk:"statement"`
}

func (d *dataOperation) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataOperationModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataOperationModel) bool {
			c := new(strings.Builder)
			c.WriteString("(")
			c.WriteString(util.StringifyValue(m.Left))
			c.WriteString(util.RawString(m.Operator))
			c.WriteString(util.StringifyValue(m.Right))
			c.WriteString(")")

			m.Expression = util.Raw(types.StringValue(c.String()))
			m.Statement = m.Expression
			return true
		},
	)
}
