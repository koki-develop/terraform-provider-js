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
	_ resource.Resource = &resourceOperation{}
)

func NewResourceOperation() func() resource.Resource {
	return func() resource.Resource {
		return &resourceOperation{}
	}
}

type resourceOperation struct{}

func (r *resourceOperation) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operation"
}

func (r *resourceOperation) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"operator": schema.StringAttribute{
				Required: true,
			},
			"left": schema.DynamicAttribute{
				Required: true,
			},
			"right": schema.DynamicAttribute{
				Required: true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceOperationModel struct {
	Operator types.String  `tfsdk:"operator"`
	Left     types.Dynamic `tfsdk:"left"`
	Right    types.Dynamic `tfsdk:"right"`

	Content types.String `tfsdk:"content"`
}

func (r *resourceOperation) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceOperation) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resourceOperation) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resourceOperation) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceOperation) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceOperationModel{},
		g,
		s,
		diags,
		func(m *resourceOperationModel) bool {
			c := new(strings.Builder)
			c.WriteString(util.StringifyValue(m.Left))
			c.WriteString(util.RawString(m.Operator))
			c.WriteString(util.StringifyValue(m.Right))

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
