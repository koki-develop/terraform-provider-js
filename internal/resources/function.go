package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jstypes "github.com/koki-develop/terraform-provider-js/internal/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceFunction{}
)

func NewResourceFunction() func() resource.Resource {
	return func() resource.Resource {
		return &resourceFunction{}
	}
}

type resourceFunction struct{}

func (r *resourceFunction) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

func (r *resourceFunction) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional: true,
			},
			"params": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"body": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},

			"id": schema.StringAttribute{
				CustomType: jstypes.ID{},
				Computed:   true,
			},
			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceFunctionModel struct {
	Name   types.String `tfsdk:"name"`
	Params types.List   `tfsdk:"params"`
	Body   types.List   `tfsdk:"body"`

	ID      jstypes.IDValue `tfsdk:"id"`
	Content types.String    `tfsdk:"content"`
}

func (m resourceFunctionModel) ContentString(ctx context.Context) (string, error) {
	s := new(strings.Builder)
	s.WriteString(jstypes.ContentPrefix)
	s.WriteString("function")
	if !m.Name.IsNull() {
		s.WriteString(" ")
		s.WriteString(m.Name.ValueString())
	}
	s.WriteString("(")

	if !m.Params.IsNull() {
		ps := make([]string, len(m.Params.Elements()))
		for i, p := range m.Params.Elements() {
			id := jstypes.NewIDValue(p.(types.String))
			ps[i] = id.ValueString()
		}
		s.WriteString(strings.Join(ps, ","))
	}
	s.WriteString("){")

	lines := util.StringifyValues(m.Body.Elements())
	s.WriteString(strings.Join(lines, ";"))

	s.WriteString("}")
	return s.String(), nil
}

func (r *resourceFunction) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceFunction) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, resp.Diagnostics)
}

func (r *resourceFunction) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceFunction) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFunction) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceFunctionModel{},
		g,
		s,
		diags,
		func(m *resourceFunctionModel) error {
			c, err := m.ContentString(ctx)
			if err != nil {
				return err
			}

			m.ID = jstypes.NewIDValue(m.Name)
			m.Content = types.StringValue(c)
			return nil
		},
	)
}
