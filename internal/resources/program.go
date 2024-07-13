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
	_ resource.Resource = &resourceProgram{}
)

func NewResourceProgram() resource.Resource {
	return &resourceProgram{}
}

type resourceProgram struct{}

func (r *resourceProgram) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_program"
}

func (r *resourceProgram) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"contents": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceProgramModel struct {
	Contents types.List   `tfsdk:"contents"`
	Content  types.String `tfsdk:"content"`
}

func (r *resourceProgram) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceProgram) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, resp.Diagnostics)
}

func (r *resourceProgram) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceProgram) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceProgram) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceProgramModel{},
		g,
		s,
		diags,
		func(m *resourceProgramModel) error {
			var cs []string
			diags.Append(m.Contents.ElementsAs(ctx, &cs, true)...)
			if diags.HasError() {
				return nil
			}

			s := new(strings.Builder)

			for _, c := range cs {
				if s.Len() > 0 && s.String()[s.Len()-1] != '}' {
					s.WriteString(";")
				}
				s.WriteString(c)
			}

			m.Content = types.StringValue(s.String())
			return nil
		},
	)
}
