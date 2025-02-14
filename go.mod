module shcli

go 1.23.0

toolchain go1.23.4

require (
	github.com/1Password/connect-sdk-go v1.5.3
	github.com/sitehostnz/gosh v0.4.0
	github.com/sitehostnz/terraform-provider-sitehost v1.2.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.6
	github.com/spf13/viper v1.19.0
	github.com/yakmoose/envop v0.1.0
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.36.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.9 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sagikazarmark/locafero v0.7.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20250210185358-939b2ce775ac // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/sitehostnz/gosh => github.com/yakmoose/gosh v0.0.0-20230802000957-ba74e1992d7b
// replace github.com/sitehostnz/terraform-provider-sitehost => github.com/yakmoose/terraform-provider-sitehost v0.0.0-20230726232211-7d012bf89998
replace github.com/sitehostnz/terraform-provider-sitehost => /Users/john/Projects/go/src/terraform-provider-sitehost

replace github.com/sitehostnz/gosh => /Users/john/Projects/go/src/gosh
