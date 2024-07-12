package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jstypes "github.com/koki-develop/terraform-provider-js/internal/types"
)

var (
	_ resource.Resource = &resourceFunctionParam{}
)

func NewResourceFunctionParam() resource.Resource {
	return &resourceFunctionParam{}
}

type resourceFunctionParam struct{}

func (r *resourceFunctionParam) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_param"
}

func (r *resourceFunctionParam) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},

			"id": schema.StringAttribute{
				CustomType: jstypes.ID{},
				Computed:   true,
			},
		},
	}
}

type resourceFunctionParamModel struct {
	Name types.String    `tfsdk:"name"`
	ID   jstypes.IDValue `tfsdk:"id"`
}

func (r *resourceFunctionParam) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFunctionParamModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = jstypes.NewIDValue(plan.Name)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionParam) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *resourceFunctionParam) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFunctionParamModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = jstypes.NewIDValue(plan.Name)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunctionParam) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
