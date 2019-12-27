/*
 * Copyright 2017 Workiva
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/generator/dartlang"
	"github.com/Workiva/frugal/compiler/generator/golang"
	"github.com/Workiva/frugal/compiler/generator/gopherjs"
	"github.com/Workiva/frugal/compiler/generator/html"
	"github.com/Workiva/frugal/compiler/generator/java"
	"github.com/Workiva/frugal/compiler/generator/json"
	"github.com/Workiva/frugal/compiler/generator/python"
	"github.com/Workiva/frugal/compiler/parser"
)

// Options contains compiler options for code generation.
type Options struct {
	Gen     string // Language to generate
	Out     string // Output location for generated code
	Delim   string // Token delimiter for scope topics
	DryRun  bool   // Do not generate code
	Recurse bool   // Generate includes
	Verbose bool   // Verbose mode
}

type Compiler struct {
	lang          string
	langOptions   map[string]string
	options       Options
	now           time.Time
	compiledFiles map[string]*parser.Frugal
	sharedGen     generator.ShareableProgramGenerator
}

func NewCompiler(options Options) (*Compiler, error) {
	lang, langOptions, err := cleanGenParam(options.Gen)
	if err != nil {
		return nil, err
	}

	c := &Compiler{
		lang:          lang,
		langOptions:   langOptions,
		options:       options,
		now:           time.Now(),
		compiledFiles: map[string]*parser.Frugal{},
	}
	return c, nil
}

// Compile parses the Frugal IDL and generates code for it, returning an error
// if something failed.
func (c *Compiler) Compile(file string) error {
	absFile, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	frugal, err := c.parseFrugal(absFile)
	if err != nil {
		return err
	}

	return c.generateFrugal(frugal)
}

// parseFrugal parses a frugal file.
func (c *Compiler) parseFrugal(file string) (*parser.Frugal, error) {
	if !exists(file) {
		return nil, fmt.Errorf("Frugal file not found: %s\n", file)
	}
	c.logv(fmt.Sprintf("Parsing %s", file))
	return parser.ParseFrugal(file)
}

// generateFrugal generates code for a frugal struct.
func (c *Compiler) generateFrugal(f *parser.Frugal) error {
	// Resolve Frugal generator.
	g, err := c.getProgramGenerator()
	if err != nil {
		return err
	}

	// The parsed frugal contains everything needed to generate
	if err := c.generateFrugalRec(f, g, true); err != nil {
		return err
	}

	return nil
}

// generateFrugalRec generates code for a frugal struct, recursively generating
// code for includes
func (c *Compiler) generateFrugalRec(f *parser.Frugal, g generator.ProgramGenerator, generate bool) error {
	if _, ok := c.compiledFiles[f.File]; ok {
		// Already generated this file
		return nil
	}
	c.compiledFiles[f.File] = f

	out := c.options.Out
	if out == "" {
		out = g.DefaultOutputDir()
	}
	fullOut := g.GetOutputDir(out, f)
	if err := os.MkdirAll(out, 0777); err != nil {
		return err
	}

	c.logv(fmt.Sprintf("Generating \"%s\" Frugal code for %s", c.lang, f.File))
	if c.options.DryRun || !generate {
		return nil
	}

	if err := g.Generate(f, fullOut); err != nil {
		return err
	}

	// Iterate through includes in order to ensure determinism in
	// generated code.
	for _, include := range f.OrderedIncludes() {
		// Skip recursive generation if include is marked vendor and use_vendor option is enabled
		if _, vendored := include.Annotations.Vendor(); vendored && g.UseVendor() {
			continue
		}
		inclFrugal := f.ParsedIncludes[include.Name]
		if err := c.generateFrugalRec(inclFrugal, g, c.options.Recurse); err != nil {
			return err
		}
	}

	return nil
}

// getProgramGenerator resolves the ProgramGenerator for the given language. It
// returns an error if the language is not supported.
func (c *Compiler) getProgramGenerator() (generator.ProgramGenerator, error) {
	if c.sharedGen != nil {
		return c.sharedGen, nil
	}

	g, err := c.createProgramGenerator()
	if err == nil {
		if sg, ok := g.(generator.ShareableProgramGenerator); ok {
			c.sharedGen = sg
		}
	}

	return g, err
}

func (c *Compiler) createProgramGenerator() (generator.ProgramGenerator, error) {
	options := c.langOptions
	config := &generator.Config{
		Now:            c.now,
		GlobalOut:      c.options.Out,
		TopicDelimiter: c.options.Delim,
		Options:        options,
	}

	var g generator.ProgramGenerator
	lang := c.lang
	switch lang {
	case "dart":
		g = generator.NewProgramGenerator(dartlang.NewGenerator(config), false)
	case "go":
		// Make sure the package prefix ends with a "/"
		if package_prefix, ok := options["package_prefix"]; ok {
			if package_prefix != "" && !strings.HasSuffix(package_prefix, "/") {
				options["package_prefix"] = package_prefix + "/"
			}
		}

		g = generator.NewProgramGenerator(golang.NewGenerator(config), false)
	case "gopherjs":
		if pkg := options["package_prefix"]; pkg != "" && !strings.HasSuffix(pkg, "/") {
			options["package_prefix"] += "/" // Make sure the package prefix ends with a "/"
		}
		g = generator.NewProgramGenerator(gopherjs.NewGenerator(config), false)
	case "java":
		g = generator.NewProgramGenerator(java.NewGenerator(config), true)
	case "json":
		g = json.NewGenerator(config)
	case "py":
		g = generator.NewProgramGenerator(python.NewGenerator(config), true)
	case "html":
		g = html.NewGenerator(options)
	default:
		return nil, fmt.Errorf("Invalid gen value %s", lang)
	}
	return g, nil
}

func (c *Compiler) Close() error {
	if c.sharedGen != nil {
		return c.sharedGen.TeardownShared()
	}
	return nil
}

// exists determines if the file at the given path exists.
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// cleanGenParam processes a string that includes an optional trailing
// options set.  Format: <language>:<name>=<value>,<name>=<value>,...
func cleanGenParam(gen string) (lang string, options map[string]string, err error) {
	lang = gen
	options = make(map[string]string)
	if strings.Contains(gen, ":") {
		s := strings.Split(gen, ":")
		lang = s[0]
		dirty := s[1]
		var optionArray []string
		if strings.Contains(dirty, ",") {
			optionArray = strings.Split(dirty, ",")
		} else {
			optionArray = append(optionArray, dirty)
		}
		for _, option := range optionArray {
			s := strings.Split(option, "=")
			if !generator.ValidateOption(lang, s[0]) {
				err = fmt.Errorf("Unknown option '%s' for %s", s[0], lang)
			}
			if len(s) == 1 {
				options[s[0]] = ""
			} else {
				options[s[0]] = s[1]
			}
		}
	}
	return
}

// logv prints the message if in verbose mode.
func (c *Compiler) logv(msg string) {
	if c.options.Verbose {
		fmt.Println(msg)
	}
}
