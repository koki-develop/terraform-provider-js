package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceFunctionCall{}
)

func NewResourceFunctionCall() func() resource.Resource {
	return func() resource.Resource {
		return &resourceFunctionCall{}
	}
}

type resourceFunctionCall struct{}

func (r *resourceFunctionCall) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_call"
}

func (r *resourceFunctionCall) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_function_call` resource calls a function.",
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

type resourceFunctionCallModel struct {
	Caller   types.String  `tfsdk:"caller"`
	Function types.String  `tfsdk:"function"`
	Args     types.Dynamic `tfsdk:"args"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceFunctionCall) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionCall) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionCall) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionCall) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFunctionCall) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceFunctionCallModel{},
		g,
		s,
		diags,
		func(m *resourceFunctionCallModel) bool {
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
