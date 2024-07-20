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
	_ resource.Resource = &resourceReturn{}
)

func NewResourceReturn() func() resource.Resource {
	return func() resource.Resource {
		return &resourceReturn{}
	}
}

type resourceReturn struct{}

func (r *resourceReturn) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_return"
}

func (r *resourceReturn) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_return` resource creates a return.",
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "The value of the return.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the return.",
				Computed:    true,
			},
		},
	}
}

type resourceReturnModel struct {
	Value types.Dynamic `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceReturn) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceReturn) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceReturn) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceReturn) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceReturn) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceReturnModel{},
		g,
		s,
		diags,
		func(m *resourceReturnModel) bool {
			c := new(strings.Builder)
			c.WriteString("return")
			if !m.Value.IsNull() {
				c.WriteString(" ")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
