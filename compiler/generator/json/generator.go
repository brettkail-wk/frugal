/*
 * Copyright 2019 Workiva
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

package json

import (
	"encoding/json"
	"os"

	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/parser"
)

const (
	defaultOutputDir = "gen-json"
)

// Generator generators JSON descriptor files.
type Generator struct {
	outputDir string
	frugals   []*parser.Frugal
	used      map[string]struct{}
}

var gen *Generator

// NewGenerator returns a generator for JSON descriptor files.
func NewGenerator(config *generator.Config) generator.ShareableProgramGenerator {
	return &Generator{
		used: map[string]struct{}{},
	}
}

func (*Generator) GetOutputDir(dir string, frugal *parser.Frugal) string {
	return dir
}

func (*Generator) DefaultOutputDir() string {
	return defaultOutputDir
}

func (*Generator) UseVendor() bool {
	return false
}

// Generate appends the frugal file to the list.
func (g *Generator) Generate(pf *parser.Frugal, outputDir string) error {
	g.outputDir = outputDir
	g.collectFrugals(pf)
	return nil
}

func (g *Generator) collectFrugals(pf *parser.Frugal) {
	if _, ok := g.used[pf.Name]; ok {
		return
	}
	g.used[pf.Name] = struct{}{}
	g.frugals = append(g.frugals, pf)

	for _, include := range pf.OrderedIncludes() {
		// TODO
		//if _, vendored := include.Annotations.Vendor(); vendored && g.UseVendor() {
		//continue
		//}
		inclFrugal := pf.ParsedIncludes[include.Name]
		g.collectFrugals(inclFrugal)
	}
}

// TeardownShared generates the JSON descriptor file.
func (g *Generator) TeardownShared() error {
	f, err := os.OpenFile(g.outputDir+"/frugal.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	ff := g.toFiles()

	enc := json.NewEncoder(f)
	return enc.Encode(ff)
}

type frugalFile struct {
	Services map[string]*service   `json:"s,omitempty"`
	Types    map[string]frugalType `json:"t,omitempty"`
}

type frugalType interface{}

type derivedType struct {
	Key          frugalType       `json:"k,omitempty"`
	Value        frugalType       `json:"v,omitempty"`
	Name         string           `json:"n,omitempty"`
	EnumValues   map[int][]string `json:"e,omitempty"`
	StructFields map[int]*field   `json:"f,omitempty"`
}

type service struct {
	Methods map[string]*method `json:"m"`
}

type method struct {
	Params  map[int]*field `json:"p"`
	Results map[int]*field `json:"r"`
}

type field struct {
	Name string      `json:"n,omitempty"`
	Type interface{} `json:"t"`
}

func (g *Generator) toFiles() map[string]*frugalFile {
	files := map[string]*frugalFile{}
	for _, pf := range g.frugals {
		files[pf.Name] = toFile(pf)
	}
	return files
}

func toFile(pf *parser.Frugal) *frugalFile {
	f := &frugalFile{
		Types:    map[string]frugalType{},
		Services: map[string]*service{},
	}

	for _, ps := range pf.Services {
		f.Services[ps.Name] = toService(ps)
	}

	for _, pt := range pf.Typedefs {
		f.Types[pt.Name] = toType(pt.Type)
	}

	for _, pe := range pf.Enums {
		f.Types[pe.Name] = toEnum(pe)
	}

	for _, ps := range pf.Structs {
		f.Types[ps.Name] = toStructType(ps)
	}
	for _, pu := range pf.Unions {
		f.Types[pu.Name] = toStructType(pu)
	}
	for _, pe := range pf.Exceptions {
		f.Types[pe.Name] = toStructType(pe)
	}

	return f
}

func toService(ps *parser.Service) *service {
	s := &service{
		Methods: map[string]*method{},
	}

	for _, pm := range ps.Methods {
		s.Methods[pm.Name] = toMethod(pm)
	}

	return s
}

func toMethod(pm *parser.Method) *method {
	m := &method{
		Params:  map[int]*field{},
		Results: map[int]*field{},
	}

	for _, pf := range pm.Arguments {
		m.Params[pf.ID] = toField(pf)
	}

	if pm.ReturnType != nil {
		m.Results[0] = &field{Type: toType(pm.ReturnType)}
	}
	for _, pf := range pm.Exceptions {
		m.Results[pf.ID] = toField(pf)
	}

	return m
}

func toField(pf *parser.Field) *field {
	return &field{
		Name: pf.Name,
		Type: toType(pf.Type),
	}
}

var baseTypes = map[string]frugalType{
	"bool":   frugalType("z"),
	"byte":   frugalType("y"),
	"i8":     frugalType("x"),
	"i16":    frugalType("s"),
	"i32":    frugalType("i"),
	"i64":    frugalType("l"),
	"double": frugalType("d"),
	"string": frugalType("s"),
	"binary": frugalType("n"),
}

func toType(pt *parser.Type) frugalType {
	if pt.IsContainer() {
		switch {
		case pt.Name == "list":
			return &derivedType{Value: toType(pt.ValueType)}
		case pt.Name == "set":
			return &derivedType{Key: toType(pt.ValueType)}
		case pt.Name == "map":
			return &derivedType{
				Key:   toType(pt.KeyType),
				Value: toType(pt.ValueType),
			}
		}
		panic("unknown container type " + pt.Name)
	}

	if t := baseTypes[pt.Name]; t != nil {
		return t
	}

	return derivedType{Name: pt.Name}
}

func toStructType(ps *parser.Struct) derivedType {
	fields := map[int]*field{}
	for _, pf := range ps.Fields {
		fields[pf.ID] = &field{
			Name: pf.Name,
			Type: toType(pf.Type),
		}
	}
	return derivedType{StructFields: fields}
}

func toEnum(pe *parser.Enum) derivedType {
	values := map[int][]string{}
	for _, pev := range pe.Values {
		values[pev.Value] = append(values[pev.Value], pev.Name)
	}
	return derivedType{EnumValues: values}
}
