package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			// TODO: args

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceFunctionModel struct {
	Function IDValue      `tfsdk:"function"`
	Content  types.String `tfsdk:"content"`
}

func (r *resourceFunctionCall) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFunctionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Content = types.StringValue(fmt.Sprintf("%s()", plan.Function.ValueString()))
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionCall) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *resourceFunctionCall) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFunctionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Content = types.StringValue(fmt.Sprintf("%s()", plan.Function.ValueString()))
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionCall) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
