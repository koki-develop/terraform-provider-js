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
	_ resource.Resource = &resourceFor{}
)

func NewResourceFor() func() resource.Resource {
	return func() resource.Resource {
		return &resourceFor{}
	}
}

type resourceFor struct{}

func (r *resourceFor) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_for"
}

func (r *resourceFor) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"init": schema.StringAttribute{
				Optional: true,
			},
			"condition": schema.StringAttribute{
				Optional: true,
			},
			"update": schema.StringAttribute{
				Optional: true,
			},
			"body": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceForModel struct {
	Init      types.String `tfsdk:"init"`
	Condition types.String `tfsdk:"condition"`
	Update    types.String `tfsdk:"update"`
	Body      types.List   `tfsdk:"body"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceFor) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFor) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceFor) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceFor) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFor) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceForModel{},
		g,
		s,
		diags,
		func(m *resourceForModel) bool {
			c := new(strings.Builder)
			c.WriteString("for(")
			if !m.Init.IsNull() {
				c.WriteString(util.StringifyValue(m.Init))
			}
			c.WriteString(";")

			if !m.Condition.IsNull() {
				c.WriteString(util.StringifyValue(m.Condition))
			}
			c.WriteString(";")

			if !m.Update.IsNull() {
				c.WriteString(util.StringifyValue(m.Update))
			}

			c.WriteString("){")

			if !m.Body.IsNull() {
				c.WriteString(util.StringifyStatements(m.Body.Elements()))
			}
			c.WriteString("}")

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
