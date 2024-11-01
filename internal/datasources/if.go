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
		MarkdownDescription: "The `js_if` data source creates an if statement.",
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Description: "The condition expression.",
				Required:    true,
			},
			"then": schema.ListAttribute{
				Description: "The body of the if statement.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"else": schema.ListAttribute{
				Description: "The body of the else statement.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the if statement.",
				Computed:    true,
			},
		},
	}
}

type dataIfModel struct {
	Condition types.String `tfsdk:"condition"`
	Then      types.List   `tfsdk:"then"`
	Else      types.List   `tfsdk:"else"`

	Content types.String `tfsdk:"content"`
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

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
