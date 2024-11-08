package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataFunction{}
)

func NewDataFunction() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataFunction{}
	}
}

type dataFunction struct{}

func (d *dataFunction) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

func (d *dataFunction) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Name of function.",
				Optional:    true,
			},
			"params": schema.ListAttribute{
				Description: "Parameters of function.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"body": schema.ListAttribute{
				Description: "Body of function.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"async": schema.BoolAttribute{
				Description: "Whether function is async.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Computed: true,
			},
			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataFunctionModel struct {
	Name   types.String `tfsdk:"name"`
	Params types.List   `tfsdk:"params"`
	Body   types.List   `tfsdk:"body"`
	Async  types.Bool   `tfsdk:"async"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (d *dataFunction) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.handleRequest(ctx, &req.Config, &resp.State, &resp.Diagnostics)
}

func (d *dataFunction) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&dataFunctionModel{},
		g,
		s,
		diags,
		func(m *dataFunctionModel) bool {
			c := new(strings.Builder)
			if !m.Async.IsNull() && m.Async.ValueBool() {
				c.WriteString("async ")
			}

			c.WriteString("function")
			if !m.Name.IsNull() {
				c.WriteString(" ")
				c.WriteString(util.RawString(m.Name))
			}
			c.WriteString("(")

			if !m.Params.IsNull() {
				ps := make([]string, len(m.Params.Elements()))
				for i, p := range m.Params.Elements() {
					ps[i] = util.RawString(p.(types.String))
				}
				c.WriteString(strings.Join(ps, ","))
			}
			c.WriteString("){")
			c.WriteString(util.StringifyStatements(m.Body.Elements()))
			c.WriteString("}")

			m.ID = util.Raw(m.Name)
			m.Content = util.Raw(types.StringValue(c.String()))

			return true
		},
	)
}
