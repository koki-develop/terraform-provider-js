package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ resource.Resource = &resourceFunctionCall{}
)

func NewResourceFunctionCall() func() resource.Resource {
	return func() resource.Resource {
		return &resourceFunctionCall{}
	}
}

type resourceFunctionCall struct{}

func (r *resourceFunctionCall) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function_call"
}

func (r *resourceFunctionCall) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"caller": schema.StringAttribute{
				Optional: true,
			},
			"function": schema.StringAttribute{
				Required: true,
			},
			"args": schema.DynamicAttribute{
				Optional: true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type resourceFunctionCallModel struct {
	Caller   types.String  `tfsdk:"caller"`
	Function types.String  `tfsdk:"function"`
	Args     types.Dynamic `tfsdk:"args"`

	Content types.String `tfsdk:"content"`
}

func (m resourceFunctionCallModel) ContentString(ctx context.Context) (string, error) {
	s := new(strings.Builder)

	if !m.Caller.IsNull() {
		s.WriteString(fmt.Sprintf("%s.", util.RawString(m.Caller)))
	}

	s.WriteString(util.RawString(m.Function))
	s.WriteString("(")

	if !m.Args.IsNull() {
		v, ok := m.Args.UnderlyingValue().(basetypes.TupleValue)
		if !ok {
			return "", fmt.Errorf("args must be a tuple")
		}
		args := util.StringifyValues(v.Elements())
		s.WriteString(strings.Join(args, ","))
	}

	s.WriteString(")")
	return s.String(), nil
}

func (r *resourceFunctionCall) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceFunctionCall) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, resp.Diagnostics)
}

func (r *resourceFunctionCall) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, resp.Diagnostics)
}

func (r *resourceFunctionCall) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resourceFunctionCall) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resourceFunctionCallModel{},
		g,
		s,
		diags,
		func(m *resourceFunctionCallModel) error {
			c, err := m.ContentString(ctx)
			if err != nil {
				return err
			}

			m.Content = util.Raw(types.StringValue(c))
			return nil
		},
	)
}
