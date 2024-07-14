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
	_ resource.Resource = &resourceLet{}
)

func NewResourceLet() func() resource.Resource {
	return func() resource.Resource {
		return &resourceLet{}
	}
}

type resourceLet struct{}

func (r *resourceLet) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_let"
}

func (r *resourceLet) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"value": schema.DynamicAttribute{
				Required: true,
			},

			"id": schema.StringAttribute{
				Computed: true,
			},
			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceLetModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (r *resourceLet) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceLet) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceLet) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceLet) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceLet) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceLetModel{},
		g,
		s,
		diags,
		func(m *resourceLetModel) bool {
			m.ID = util.Raw(m.Name)
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("let %s=%s", util.RawString(m.Name), util.StringifyValue(m.Value))))
			return true
		},
	)
}
