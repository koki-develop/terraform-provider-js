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
	_ resource.Resource = &resourceVar{}
)

func NewResourceVar() func() resource.Resource {
	return func() resource.Resource {
		return &resourceVar{}
	}
}

type resourceVar struct{}

func (r *resourceVar) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_var"
}

func (r *resourceVar) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_var` resource defines a variable.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the variable.",
				Required:    true,
			},
			"value": schema.DynamicAttribute{
				Description: "The value of the variable.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "The id of the variable.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the variable.",
				Computed:    true,
			},
		},
	}
}

type resourceVarModel struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (r *resourceVar) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceVar) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceVar) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceVar) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceVar) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceVarModel{},
		g,
		s,
		diags,
		func(m *resourceVarModel) bool {
			m.ID = util.Raw(m.Name)

			c := new(strings.Builder)
			c.WriteString("var ")
			c.WriteString(util.RawString(m.Name))
			if !m.Value.IsNull() {
				c.WriteString("=")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
