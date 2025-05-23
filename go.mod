module shcli

go 1.24.2

require (
	github.com/1Password/connect-sdk-go v1.5.3
	github.com/sitehostnz/gosh v0.4.0
	github.com/sitehostnz/terraform-provider-sitehost v1.2.0
	github.com/spf13/cobra v1.9.1
	github.com/spf13/pflag v1.0.6
	github.com/spf13/viper v1.20.1
	github.com/yakmoose/envop v0.2.0
)

require (
	github.com/1password/onepassword-sdk-go v0.2.1 // indirect
	github.com/dylibso/observe-sdk/go v0.0.0-20240828172851-9145d8ad07e1 // indirect
	github.com/extism/go-sdk v1.7.1 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-envparse v0.1.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.36.0 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20240912202439-0a2b6291aafd // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tetratelabs/wabin v0.0.0-20230304001439-f6f874872834 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/sitehostnz/gosh => github.com/yakmoose/gosh v0.0.0-20230802000957-ba74e1992d7b
// replace github.com/sitehostnz/terraform-provider-sitehost => github.com/yakmoose/terraform-provider-sitehost v0.0.0-20230726232211-7d012bf89998

replace github.com/sitehostnz/terraform-provider-sitehost => /Users/john/Projects/go/src/terraform-provider-sitehost
replace github.com/sitehostnz/gosh => /Users/john/Projects/go/src/gosh
replace github.com/yakmoose/envop => /Users/john/Projects/go/src/envop
