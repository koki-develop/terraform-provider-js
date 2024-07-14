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
	_ resource.Resource = &resourceIf{}
)

func NewResourceIf() func() resource.Resource {
	return func() resource.Resource {
		return &resourceIf{}
	}
}

type resourceIf struct{}

func (r *resourceIf) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_if"
}

func (r *resourceIf) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"condition": schema.StringAttribute{
				Required: true,
			},
			"then": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"else": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceIfModel struct {
	Condition types.String `tfsdk:"condition"`
	Then      types.List   `tfsdk:"then"`
	Else      types.List   `tfsdk:"else"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceIf) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceIf) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceIf) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceIf) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceIf) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceIfModel{},
		g,
		s,
		diags,
		func(m *resourceIfModel) bool {
			c := new(strings.Builder)
			c.WriteString("if(")
			c.WriteString(util.StringifyValue(m.Condition))
			c.WriteString("){")
			if !m.Then.IsNull() {
				c.WriteString(util.StringifyStatements(m.Then.Elements()))
			}
			c.WriteString("}")
			if !m.Else.IsNull() {
				c.WriteString("else{")
				c.WriteString(util.StringifyStatements(m.Else.Elements()))
				c.WriteString("}")
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
