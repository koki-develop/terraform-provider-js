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
	_ resource.Resource = &resourceIncrement{}
)

func NewResourceIncrement() func() resource.Resource {
	return func() resource.Resource {
		return &resourceIncrement{}
	}
}

type resourceIncrement struct{}

func (r *resourceIncrement) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_increment"
}

func (r *resourceIncrement) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_increment` resource increments a reference.",
		Attributes: map[string]schema.Attribute{
			"ref": schema.StringAttribute{
				Description: "The reference to increment.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of increment to perform. (Valid values: `prefix`, `postfix`)",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("prefix", "postfix"),
				},
			},

			"content": schema.StringAttribute{
				Description: "The content of the increment.",
				Computed:    true,
			},
		},
	}
}

type resourceIncrementModel struct {
	Ref  types.String `tfsdk:"ref"`
	Type types.String `tfsdk:"type"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceIncrement) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceIncrement) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceIncrement) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceIncrement) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceIncrement) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceIncrementModel{},
		g,
		s,
		diags,
		func(m *resourceIncrementModel) bool {
			c := new(strings.Builder)
			tp := m.Type.ValueString()

			if tp == "prefix" {
				c.WriteString("++")
			}
			c.WriteString(util.RawString(m.Ref))
			if m.Type.IsNull() || tp == "postfix" {
				c.WriteString("++")
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
