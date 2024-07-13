package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceFunctionParam{}
)

func NewResourceFunctionParam() func() resource.Resource {
	return func() resource.Resource {
		return &resourceFunctionParam{}
	}
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
				Computed: true,
			},
		},
	}
}

type resourceFunctionParamModel struct {
	Name types.String `tfsdk:"name"`
	ID   types.String `tfsdk:"id"`
}

func (r *resourceFunctionParam) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionParam) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionParam) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFunctionParam) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFunctionParam) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceFunctionParamModel{},
		g,
		s,
		diags,
		func(m *resourceFunctionParamModel) bool {
			m.ID = util.Raw(m.Name)
			return true
		},
	)
}
