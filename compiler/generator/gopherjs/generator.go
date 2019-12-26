package gopherjs

import (
	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/generator/golang"
)

// Generator implements the LanguageGenerator interface for Go.
type Generator struct {
	*golang.Generator
}

// NewGenerator creates a new Go LanguageGenerator.
func NewGenerator(config *generator.Config) generator.LanguageGenerator {
	options := config.Options
	options["slim"] = "true"
	options["frugal_import"] = "github.com/Workiva/frugal/lib/gopherjs/frugal"
	options["thrift_import"] = "github.com/Workiva/frugal/lib/gopherjs/thrift"
	return &Generator{
		Generator: &golang.Generator{
			BaseGenerator: &generator.BaseGenerator{
				Options: options,
			},
			Delim: config.TopicDelimiter,
		},
	}
}
