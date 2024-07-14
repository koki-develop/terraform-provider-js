package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceDecrement{}
)

func NewResourceDecrement() func() resource.Resource {
	return func() resource.Resource {
		return &resourceDecrement{}
	}
}

type resourceDecrement struct{}

func (r *resourceDecrement) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_decrement"
}

func (r *resourceDecrement) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ref": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("prefix", "postfix"),
				},
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceDecrementModel struct {
	Ref  types.String `tfsdk:"ref"`
	Type types.String `tfsdk:"type"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceDecrement) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceDecrement) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceDecrement) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceDecrement) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceDecrement) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceDecrementModel{},
		g,
		s,
		diags,
		func(m *resourceDecrementModel) bool {
			c := new(strings.Builder)
			tp := m.Type.ValueString()

			if tp == "prefix" {
				c.WriteString("--")
			}
			c.WriteString(util.RawString(m.Ref))
			if m.Type.IsNull() || tp == "postfix" {
				c.WriteString("--")
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
