package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataFunctionCall{}
)

func NewDataFunctionCall() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataFunctionCall{}
	}
}

type dataFunctionCall struct{}

func (d *dataFunctionCall) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_call"
}

func (d *dataFunctionCall) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_function_call` data source calls a function.",
		Attributes: map[string]schema.Attribute{
			"caller": schema.StringAttribute{
				Description: "The caller of the function.",
				Optional:    true,
			},
			"function": schema.StringAttribute{
				Description: "The function to call.",
				Required:    true,
			},
			"args": schema.DynamicAttribute{
				Description: "The arguments of the function.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the function call.",
				Computed:    true,
			},
		},
	}
}

type dataFunctionCallModel struct {
	Caller   types.String  `tfsdk:"caller"`
	Function types.String  `tfsdk:"function"`
	Args     types.Dynamic `tfsdk:"args"`

	Content types.String `tfsdk:"content"`
}

func (d *dataFunctionCall) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.handleRequest(ctx, &req.Config, &resp.State, &resp.Diagnostics)
}

func (d *dataFunctionCall) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&dataFunctionCallModel{},
		g,
		s,
		diags,
		func(m *dataFunctionCallModel) bool {
			c := new(strings.Builder)

			if !m.Caller.IsNull() {
				c.WriteString(fmt.Sprintf("%s.", util.RawString(m.Caller)))
			}

			c.WriteString(util.RawString(m.Function))
			c.WriteString("(")

			if !m.Args.IsNull() {
				var elms []attr.Value
				switch v := m.Args.UnderlyingValue().(type) {
				case basetypes.ListValue:
					elms = v.Elements()
				case basetypes.TupleValue:
					elms = v.Elements()
				case basetypes.SetValue:
					elms = v.Elements()
				default:
					diags.AddError("Invalid type of args", "args must be a list, tuple or set")
					return false
				}

				args := util.StringifyValues(elms)
				c.WriteString(strings.Join(args, ","))
			}

			c.WriteString(")")

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
