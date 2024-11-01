package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/koki-develop/terraform-provider-js/internal/datasources"
	"github.com/koki-develop/terraform-provider-js/internal/resources"
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
		datasources.NewDataIndex(),
		datasources.NewDataRaw(),
	}
}

func (p *jsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewResourceAwait(),
		resources.NewResourceConst(),
		resources.NewResourceDecrement(),
		resources.NewResourceExport(),
		resources.NewResourceFor(),
		resources.NewResourceFunction(),
		resources.NewResourceFunctionCall(),
		resources.NewResourceFunctionParam(),
		resources.NewResourceIf(),
		resources.NewResourceImport(),
		resources.NewResourceIncrement(),
		resources.NewResourceLet(),
		resources.NewResourceNew(),
		resources.NewResourceOperation(),
		resources.NewResourceProgram(p.version),
		resources.NewResourceReturn(),
		resources.NewResourceThrow(),
		resources.NewResourceVar(),
		resources.NewResourceWhile(),
	}
}
