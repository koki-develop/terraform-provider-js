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
	_ resource.Resource = &resourceExport{}
)

func NewResourceExport() func() resource.Resource {
	return func() resource.Resource {
		return &resourceExport{}
	}
}

type resourceExport struct{}

func (r *resourceExport) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_export"
}

func (r *resourceExport) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_export` resource defines an export.",
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "The value of the export.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The content of the export.",
				Computed:    true,
			},
		},
	}
}

type resourceExportModel struct {
	Value types.Dynamic `tfsdk:"value"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceExport) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceExport) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceExport) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceExport) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceExport) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceExportModel{},
		g,
		s,
		diags,
		func(m *resourceExportModel) bool {
			c := new(strings.Builder)
			c.WriteString("export ")
			c.WriteString(util.StringifyValue(m.Value))

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
