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
	_ resource.Resource = &resourceNew{}
)

func NewResourceNew() func() resource.Resource {
	return func() resource.Resource {
		return &resourceNew{}
	}
}

type resourceNew struct{}

func (r *resourceNew) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_new"
}

func (r *resourceNew) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_new` resource creates a new operation.",
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "The value of the operation.",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the operation.",
				Computed:    true,
			},
		},
	}
}

type resourceNewModel struct {
	Value types.String `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceNew) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceNew) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceNew) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceNew) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceNew) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceNewModel{},
		g,
		s,
		diags,
		func(m *resourceNewModel) bool {
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("new %s", util.StringifyValue(m.Value))))
			return true
		},
	)
}
