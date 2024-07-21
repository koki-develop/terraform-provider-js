package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceAwait{}
)

func NewResourceAwait() func() resource.Resource {
	return func() resource.Resource {
		return &resourceAwait{}
	}
}

type resourceAwait struct{}

func (r *resourceAwait) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_await"
}

func (r *resourceAwait) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_await` resource defines an await.",
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "The value of the await.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the await.",
				Computed:    true,
			},
		},
	}
}

type resourceAwaitModel struct {
	Value types.Dynamic `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceAwait) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceAwait) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceAwait) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceAwait) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceAwait) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceAwaitModel{},
		g,
		s,
		diags,
		func(m *resourceAwaitModel) bool {
			c := new(strings.Builder)
			c.WriteString("await ")
			c.WriteString(util.StringifyValue(m.Value))

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
