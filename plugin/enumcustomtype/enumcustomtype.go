package enumcustomtype

import "github.com/gogo/protobuf/protoc-gen-gogo/generator"

type plugin struct {
	*generator.Generator
	generator.PluginImports
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "enumcustomtype"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	p.P("// It is alive!")
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}
