package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataIncrement{}
)

func NewDataIncrement() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataIncrement{}
	}
}

type dataIncrement struct{}

func (d *dataIncrement) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_increment"
}

func (d *dataIncrement) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ref": schema.StringAttribute{
				Description: "Reference to increment.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Type of increment to perform. (Valid values: `prefix`, `postfix`)",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("prefix", "postfix"),
				},
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataIncrementModel struct {
	Ref  types.String `tfsdk:"ref"`
	Type types.String `tfsdk:"type"`

	Content types.String `tfsdk:"content"`
}

func (d *dataIncrement) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataIncrementModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataIncrementModel) bool {
			c := new(strings.Builder)
			tp := m.Type.ValueString()

			if tp == "prefix" {
				c.WriteString("++")
			}
			c.WriteString(util.RawString(m.Ref))
			if m.Type.IsNull() || tp == "postfix" {
				c.WriteString("++")
			}

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
