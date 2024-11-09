package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataReturn{}
)

func NewDataReturn() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataReturn{}
	}
}

type dataReturn struct{}

func (d *dataReturn) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_return"
}

func (d *dataReturn) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"value": schema.DynamicAttribute{
				Description: "Expression whose value is to be returned.",
				Optional:    true,
			},

			"statement": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataReturnModel struct {
	Value types.Dynamic `tfsdk:"value"`

	Statement types.String `tfsdk:"statement"`
}

func (d *dataReturn) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataReturnModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataReturnModel) bool {
			c := new(strings.Builder)
			c.WriteString("return")
			if !m.Value.IsNull() {
				c.WriteString(" ")
				c.WriteString(util.StringifyValue(m.Value))
			}

			m.Statement = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
