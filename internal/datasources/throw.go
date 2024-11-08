package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataThrow{}
)

func NewDataThrow() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataThrow{}
	}
}

type dataThrow struct{}

func (d *dataThrow) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throw"
}

func (d *dataThrow) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "Expression to throw.",
				Required:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataThrowModel struct {
	Value   types.String `tfsdk:"value"`
	Content types.String `tfsdk:"content"`
}

func (d *dataThrow) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataThrowModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataThrowModel) bool {
			m.Content = util.Raw(types.StringValue(fmt.Sprintf("throw %s", util.StringifyValue(m.Value))))
			return true
		},
	)
}
