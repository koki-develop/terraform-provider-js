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
	_ datasource.DataSource = &dataWhile{}
)

func NewDataWhile() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataWhile{}
	}
}

type dataWhile struct{}

func (d *dataWhile) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_while"
}

func (d *dataWhile) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_while` data source creates a while loop.",
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Description: "The condition expression.",
				Required:    true,
			},
			"body": schema.ListAttribute{
				Description: "The body of the while loop.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the while loop.",
				Computed:    true,
			},
		},
	}
}

type dataWhileModel struct {
	Condition types.String `tfsdk:"condition"`
	Body      types.List   `tfsdk:"body"`
	Content   types.String `tfsdk:"content"`
}

func (d *dataWhile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataWhileModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataWhileModel) bool {
			c := new(strings.Builder)
			c.WriteString("while(")
			c.WriteString(util.StringifyValue(m.Condition))
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
