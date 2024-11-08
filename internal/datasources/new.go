package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/koki-develop/terraform-provider-js/internal/util"
)

var (
	_ datasource.DataSource = &dataNew{}
)

func NewDataNew() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &dataNew{}
	}
}

type dataNew struct{}

func (d *dataNew) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_new"
}

func (d *dataNew) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"constructor": schema.StringAttribute{
				Description: "A class or function that specifies the type of the object instance.",
				Required:    true,
			},
			"args": schema.DynamicAttribute{
				Description: "A list of values that the constructor will be called with.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type dataNewModel struct {
	Constructor types.String  `tfsdk:"constructor"`
	Args        types.Dynamic `tfsdk:"args"`
	Content     types.String  `tfsdk:"content"`
}

func (d *dataNew) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	util.HandleRequest(
		ctx,
		&dataNewModel{},
		&req.Config,
		&resp.State,
		&resp.Diagnostics,
		func(m *dataNewModel) bool {
			c := new(strings.Builder)

			c.WriteString("new ")
			c.WriteString(util.RawString(m.Constructor))

			c.WriteString("(")

			if !m.Args.IsNull() {
				var elms []attr.Value
				switch v := m.Args.UnderlyingValue().(type) {
				case basetypes.ListValue:
					elms = v.Elements()
				case basetypes.TupleValue:
					elms = v.Elements()
				case basetypes.SetValue:
					elms = v.Elements()
				default:
					resp.Diagnostics.AddError("Invalid type of args", "args must be a list, tuple or set")
					return false
				}

				args := util.StringifyValues(elms)
				c.WriteString(strings.Join(args, ","))
			}

			c.WriteString(")")

			m.Content = util.Raw(types.StringValue(c.String()))
			return true
		},
	)
}
