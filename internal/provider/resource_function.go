package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &resourceFunction{}
)

func NewResourceFunction() resource.Resource {
	return &resourceFunction{}
}

type resourceFunction struct{}

func (r *resourceFunction) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

func (r *resourceFunction) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
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
				CustomType: ID{},
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

	ID      IDValue      `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (m resourceFunctionModel) ContentString(ctx context.Context) (string, error) {
	s := new(strings.Builder)
	s.WriteString(fmt.Sprintf("function %s(", m.Name.ValueString()))

	if !m.Params.IsNull() {
		ps := make([]string, len(m.Params.Elements()))
		for i, p := range m.Params.Elements() {
			ps[i] = p.(types.String).ValueString()
		}
		s.WriteString(strings.Join(ps, ","))
	}
	s.WriteString("){")

	lines := make([]string, len(m.Body.Elements()))
	for i, b := range m.Body.Elements() {
		lines[i] = b.(types.String).ValueString()
	}
	s.WriteString(strings.Join(lines, ";"))

	s.WriteString("}")
	return s.String(), nil
}

func (r *resourceFunction) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceFunctionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := plan.ContentString(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to generate content", err.Error())
		return
	}

	plan.ID = NewIDValue(plan.Name)
	plan.Content = types.StringValue(c)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunction) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *resourceFunction) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceFunctionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := plan.ContentString(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to generate content", err.Error())
		return
	}

	plan.ID = NewIDValue(plan.Name)
	plan.Content = types.StringValue(c)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceFunction) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
