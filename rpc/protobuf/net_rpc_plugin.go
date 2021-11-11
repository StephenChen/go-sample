package main

import "github.com/golang/protobuf/protoc-gen-go/generator"

type netRpcPlugin struct {
	*generator.Generator
}

func (p *netRpcPlugin) Name() string {
	return "net_rpc"
}

func (p *netRpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *netRpcPlugin) Generate(file *generator.FileDescriptor) {
	//for _, svc := range file.Service {
	//	p.genServiceCode(svc)
	//}
}

func (p *netRpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		//p.genImportCode(file)
	}
}
