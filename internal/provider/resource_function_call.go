package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource = &resourceFunctionCall{}
)

func NewResourceFunctionCall() resource.Resource {
	return &resourceFunctionCall{}
}

type resourceFunctionCall struct{}

func (r *resourceFunctionCall) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_call"
}

func (r *resourceFunctionCall) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"function": schema.StringAttribute{
				CustomType: ID{},
				Required:   true,
			},
			"args": schema.DynamicAttribute{
				Optional: true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceFunctionCallModel struct {
	Function IDValue       `tfsdk:"function"`
	Args     types.Dynamic `tfsdk:"args"`
	Content  types.String  `tfsdk:"content"`
}

func (m resourceFunctionCallModel) ContentString(ctx context.Context) (string, error) {
	if m.Args.IsNull() {
		return fmt.Sprintf("%s()", m.Function.ValueString()), nil
	}

	v, ok := m.Args.UnderlyingValue().(basetypes.TupleValue)
	if !ok {
		return "", fmt.Errorf("args must be a tuple")
	}

	args := stringifyValues(v.Elements())
	return fmt.Sprintf("%s(%s)", m.Function.ValueString(), strings.Join(args, ",")), nil
}

func (r *resourceFunctionCall) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFunctionCallModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := plan.ContentString(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to generate content", err.Error())
		return
	}

	plan.Content = types.StringValue(c)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionCall) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *resourceFunctionCall) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFunctionCallModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := plan.ContentString(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to generate content", err.Error())
		return
	}

	plan.Content = types.StringValue(c)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionCall) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
