package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceThrow{}
)

func NewResourceThrow() func() resource.Resource {
	return func() resource.Resource {
		return &resourceThrow{}
	}
}

type resourceThrow struct{}

func (r *resourceThrow) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throw"
}

func (r *resourceThrow) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Required: true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceThrowModel struct {
	Value types.String `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceThrow) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceThrow) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceThrow) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceThrow) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceThrow) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceThrowModel{},
		g,
		s,
		diags,
		func(m *resourceThrowModel) bool {
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("throw %s", util.StringifyValue(m.Value))))
			return true
		},
	)
}
