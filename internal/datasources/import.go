package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

func NewDataImport() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataImport{}
	}
}

type dataImport struct{}

func (d *dataImport) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_import"
}

func (d *dataImport) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"from": schema.StringAttribute{
				Description: "The module to import from.",
				Required:    true,
			},
			"as": schema.StringAttribute{
				Description: "Name of the module object that will be used as a kind of namespace when referring to the imports.",
				Required:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Whether the import is default.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Computed: true,
			},
			"statement": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataImportModel struct {
	From    types.String `tfsdk:"from"`
	As      types.String `tfsdk:"as"`
	Default types.Bool   `tfsdk:"default"`

	ID        types.String `tfsdk:"id"`
	Statement types.String `tfsdk:"statement"`
}

func (d *dataImport) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataImportModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataImportModel) bool {
			c := new(strings.Builder)
			c.WriteString("import ")

			if m.Default.IsNull() || !m.Default.ValueBool() {
				c.WriteString("* as ")
			}
			c.WriteString(util.RawString(m.As))

			c.WriteString(fmt.Sprintf(" from %q", util.RawString(m.From)))

			m.ID = util.Raw(m.As)
			m.Statement = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
