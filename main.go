package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Workiva/frugal/compiler"
	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/globals"
	"github.com/Workiva/frugal/compiler/parser"
	"github.com/urfave/cli"
)

const defaultTopicDelim = "."

var (
	help    bool
	gen     string
	out     string
	delim   string
	audit   string
	recurse bool
	verbose bool
	version bool
)

func main() {
	app := cli.NewApp()
	app.Name = "frugal"
	app.Usage = "a tool for code generation"
	app.Version = globals.Version
	app.HideVersion = true
	app.HideHelp = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "help, h",
			Usage:       "show help",
			Destination: &help,
		},
		cli.StringFlag{
			Name:        "gen",
			Usage:       genUsage(),
			Destination: &gen,
		},
		cli.StringFlag{
			Name:        "out",
			Usage:       "set the output location for generated files (no gen-* folder will be created)",
			Destination: &out,
		},
		cli.StringFlag{
			Name:        "delim",
			Value:       defaultTopicDelim,
			Usage:       "set the delimiter for pub/sub topic tokens",
			Destination: &delim,
		},
		cli.BoolFlag{
			Name:        "recurse, r",
			Usage:       "generate included files",
			Destination: &recurse,
		},
		cli.BoolFlag{
			Name:        "verbose, v",
			Usage:       "verbose mode",
			Destination: &verbose,
		},
		cli.BoolFlag{
			Name:        "version",
			Usage:       "print the version",
			Destination: &version,
		}, cli.StringFlag{
			Name:        "audit",
			Usage:       "frugal file to run audit against",
			Destination: &audit,
		},
	}

	app.Action = func(c *cli.Context) error {
		if help {
			cli.ShowAppHelp(c)
			os.Exit(0)
		}

		if version {
			cli.ShowVersion(c)
			os.Exit(0)
		}

		if len(c.Args()) == 0 {
			fmt.Printf("Usage: %s [options] file\n\n", app.Name)
			fmt.Printf("Use %s -help for a list of options\n", app.Name)
			os.Exit(1)
		}

		if gen == "" && audit == "" {
			fmt.Println("No output language specified")
			fmt.Printf("Usage: %s [options] file\n\n", app.Name)
			fmt.Printf("Use %s -help for a list of options\n", app.Name)
			os.Exit(1)
		}

		options := compiler.Options{
			Gen:     gen,
			Out:     out,
			Delim:   delim,
			Recurse: recurse,
			Verbose: verbose,
		}

		comp, err := compiler.NewCompiler(options)
		if err != nil {
			fmt.Printf("Failed to parse generation options:\n\t%s\n", err.Error())
			os.Exit(1)
		}

		var file string

		// Handle panics for graceful error messages.
		defer func() {
			if r := recover(); r != nil {
				if file == "" {
					fmt.Printf("Failed to generate:\n\t%s\n", file, r)
				} else {
					fmt.Printf("Failed to generate %s:\n\t%s\n", file, r)
				}
				os.Exit(1)
			}
		}()

		auditor := parser.NewAuditor()
		for _, file = range c.Args() {
			if audit == "" {
				err = comp.Compile(file)
			} else {
				err = auditor.Audit(audit, file)
			}
			if err != nil {
				fmt.Printf("Failed to generate %s:\n\t%s\n", file, err.Error())
				os.Exit(1)
			}
		}
		file = ""

		err = comp.Close()
		if err != nil {
			fmt.Printf("Failed to generate:\n\t%s\n", err.Error())
			os.Exit(1)
		}

		return nil
	}

	app.Run(os.Args)
}

func genUsage() string {
	usage := "generate code with a registered generator and optional parameters " +
		"(lang[:key1=val1[,key2[,key3=val3]]])\n"
	langKeys := make([]string, 0, len(generator.Languages))
	for lang := range generator.Languages {
		langKeys = append(langKeys, lang)
	}
	sort.Strings(langKeys)
	langPrefix := ""
	for _, lang := range langKeys {
		options := generator.Languages[lang]
		optionsStr := ""
		optionKeys := make([]string, 0, len(options))
		for name := range options {
			optionKeys = append(optionKeys, name)
		}
		sort.Strings(optionKeys)
		for _, name := range optionKeys {
			optionsStr += fmt.Sprintf("\n\t        %s:\t%s", name, options[name])
		}
		usage += fmt.Sprintf("%s\t    %s:%s", langPrefix, lang, optionsStr)
		langPrefix = "\n"
	}
	return usage
}
