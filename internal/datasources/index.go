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
	_ datasource.DataSource = (*dataIndex)(nil)
)

type dataIndex struct{}

func NewDataIndex() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataIndex{}
	}
}

func (d *dataIndex) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_index"
}

func (d *dataIndex) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The `js_index` data source allows you to reference a JavaScript object and access its properties or array elements.",
		Attributes: map[string]schema.Attribute{
			"ref": schema.StringAttribute{
				Required:    true,
				Description: "The referenced JavaScript object.",
			},
			"value": schema.DynamicAttribute{
				Required:    true,
				Description: "The index or property name within the referenced object.",
			},

			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The id of the indexed value.",
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Description: "The content of the indexed value.",
			},
		},
	}
}

type dataIndexModel struct {
	Ref   types.String  `tfsdk:"ref"`
	Value types.Dynamic `tfsdk:"value"`

	ID      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

func (d *dataIndex) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataIndexModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataIndexModel) bool {
			c := new(strings.Builder)
			c.WriteString(util.RawString(m.Ref))
			c.WriteRune('[')
			c.WriteString(util.StringifyValue(m.Value))
			c.WriteRune(']')
			m.ID = util.Raw(types.StringValue(c.String()))
			m.Content = m.ID

			return true
		},
	)
}
