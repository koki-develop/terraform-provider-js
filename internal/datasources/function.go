package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jstypes "github.com/koki-develop/terraform-provider-js/internal/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = (*dataFunction)(nil)
)

type dataFunction struct{}

func NewDataFunction() datasource.DataSource {
	return &dataFunction{}
}

func (d *dataFunction) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_function"
}

func (d *dataFunction) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},

			"id": schema.StringAttribute{
				CustomType: jstypes.ID{},
				Computed:   true,
			},
		},
	}
}

type dataFunctionModel struct {
	Name types.String `tfsdk:"name"`

	ID jstypes.IDValue `tfsdk:"id"`
}

func (d *dataFunction) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataFunctionModel{},
		&req.Config,
		&resp.State,
		resp.Diagnostics,
		func(m *dataFunctionModel) error {
			m.ID = jstypes.NewIDValue(m.Name)
			return nil
		},
	)
}
