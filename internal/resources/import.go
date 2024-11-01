package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

func NewResourceImport() func() resource.Resource {
	return func() resource.Resource {
		return &resourceImport{}
	}
}

type resourceImport struct{}

func (r *resourceImport) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_import"
}

func (r *resourceImport) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_import` resource defines an import.",
		Attributes: map[string]schema.Attribute{
			"from": schema.StringAttribute{
				Description: "The path of the import.",
				Required:    true,
			},
			"as": schema.StringAttribute{
				Description: "The name of the import.",
				Required:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Whether the import is default.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "The ID of the import.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the import.",
				Computed:    true,
			},
		},
	}
}

type resourceImportModel struct {
	From    types.String `tfsdk:"from"`
	As      types.String `tfsdk:"as"`
	Default types.Bool   `tfsdk:"default"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (r *resourceImport) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceImport) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceImport) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceImport) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceImport) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceImportModel{},
		g,
		s,
		diags,
		func(m *resourceImportModel) bool {
			c := new(strings.Builder)
			c.WriteString("import ")

			if m.Default.IsNull() || !m.Default.ValueBool() {
				c.WriteString("* as ")
			}
			c.WriteString(util.RawString(m.As))

			c.WriteString(fmt.Sprintf(" from %q", util.RawString(m.From)))

			m.ID = util.Raw(m.As)
			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
