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
	_ datasource.DataSource = &dataIf{}
)

func NewDataIf() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataIf{}
	}
}

type dataIf struct{}

func (d *dataIf) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_if"
}

func (d *dataIf) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Description: "An expression that is considered to be either truthy or falsy.",
				Required:    true,
			},
			"then": schema.ListAttribute{
				Description: "Statements that are executed if condition is truthy.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"else": schema.ListAttribute{
				Description: "Statements that are executed if condition is falsy.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"statement": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataIfModel struct {
	Condition types.String `tfsdk:"condition"`
	Then      types.List   `tfsdk:"then"`
	Else      types.List   `tfsdk:"else"`

	Statement types.String `tfsdk:"statement"`
}

func (d *dataIf) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataIfModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataIfModel) bool {
			c := new(strings.Builder)
			c.WriteString("if(")
			c.WriteString(util.StringifyValue(m.Condition))
			c.WriteString("){")
			if !m.Then.IsNull() {
				c.WriteString(util.StringifyStatements(m.Then.Elements()))
			}
			c.WriteString("}")
			if !m.Else.IsNull() {
				c.WriteString("else{")
				c.WriteString(util.StringifyStatements(m.Else.Elements()))
				c.WriteString("}")
			}

			m.Statement = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
