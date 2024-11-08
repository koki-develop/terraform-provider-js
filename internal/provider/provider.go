package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/koki-develop/terraform-provider-js/internal/datasources"
)

var (
	_ provider.Provider = (*jsProvider)(nil)
)

type jsProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &jsProvider{
			version: version,
		}
	}
}

func (p *jsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "js"
	resp.Version = p.version
}

func (p *jsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Next Generation AltJS.",
	}
}

func (p *jsProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *jsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewDataAwait(),
		datasources.NewDataConditionalOperation(),
		datasources.NewDataConst(),
		datasources.NewDataDecrement(),
		datasources.NewDataExport(),
		datasources.NewDataFor(),
		datasources.NewDataFunction(),
		datasources.NewDataFunctionCall(),
		datasources.NewDataFunctionParam(),
		datasources.NewDataIf(),
		datasources.NewDataImport(),
		datasources.NewDataIncrement(),
		datasources.NewDataIndex(),
		datasources.NewDataLet(),
		datasources.NewDataNew(),
		datasources.NewDataOperation(),
		datasources.NewDataProgram(p.version),
		datasources.NewDataRaw(),
		datasources.NewDataReturn(),
		datasources.NewDataThrow(),
		datasources.NewDataVar(),
		datasources.NewDataWhile(),
	}
}

func (p *jsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
