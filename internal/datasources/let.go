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
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Name of variable to declare.",
				Required:    true,
			},
			"value": schema.DynamicAttribute{
				Description: "Initial value of the variable.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Computed: true,
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

type dataLetModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID         types.String `tfsdk:"id"`
	Expression types.String `tfsdk:"expression"`
	Statement  types.String `tfsdk:"statement"`
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
			m.Expression = m.ID

			c := new(strings.Builder)
			c.WriteString("let ")
			c.WriteString(util.RawString(m.Name))
			if !m.Value.IsNull() {
				c.WriteString("=")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Statement = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
