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
	_ resource.Resource = &resourceWhile{}
)

func NewResourceWhile() func() resource.Resource {
	return func() resource.Resource {
		return &resourceWhile{}
	}
}

type resourceWhile struct{}

func (r *resourceWhile) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_while"
}

func (r *resourceWhile) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Required: true,
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

type resourceWhileModel struct {
	Condition types.String `tfsdk:"condition"`
	Body      types.List   `tfsdk:"body"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceWhile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceWhile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceWhile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceWhile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceWhile) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceWhileModel{},
		g,
		s,
		diags,
		func(m *resourceWhileModel) bool {
			c := new(strings.Builder)
			c.WriteString("while(")
			c.WriteString(util.StringifyValue(m.Condition))
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
