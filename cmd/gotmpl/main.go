package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jtyr/gotmpl/pkg/process"
	"github.com/jtyr/gotmpl/pkg/version"
)

func usage(v int) {
	fmt.Printf("Usage: %s [--help|--version|<params> <tmpl>]\n", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Read parameters and template like a string")
	fmt.Printf("  %s 'name: John Doe' 'Hello {{ .name }}'\n", os.Args[0])
	fmt.Println()
	fmt.Println("  # Read parameters and template from a file")
	fmt.Printf("  %s params.yaml template.gotmpl\n", os.Args[0])
	fmt.Println()
	fmt.Println("  # Read parameters from a file and extract template from another file")
	fmt.Printf("  %s \\\n", os.Args[0])
	fmt.Println("    /tmp/input.yaml \\")
	fmt.Println("    <(yq '.spec.template.spec.source.helm.values' application-set.yaml)")
	fmt.Println()
	fmt.Println("  # Read parameters from a Secret on a Kubernetes cluster and extract template from a file")
	fmt.Printf("  %s \\\n", os.Args[0])
	fmt.Println("    <(kubectl get secret -o yaml -n argocd argocd-cluster-local | yq '. |= pick([\"metadata\"]) | .metadata |= pick([\"annotations\"])') \\")
	fmt.Println("    <(yq 'select(document_index == 1) | .spec.template.spec.source.helm.values' application-set.yaml)")

	os.Exit(v)
}

func main() {
	type Flags struct {
		help bool
		version bool
	}

	flags := Flags{}

	flag.BoolVar(&flags.help, "help", false, "shows this help message")
	flag.BoolVar(&flags.version, "version", false, "shows version details")
	flag.Parse()

	if flags.help {
		usage(0)
	} else if flags.version {
		fmt.Println(version.String())

		os.Exit(0)
	} else if flag.NArg() < 2 {
		usage(1)
	}

	params := os.Args[1]
	tmpl := os.Args[2]

	err := process.ProcessTmpl(params, tmpl)
	if err != nil {
		log.Fatal("failed to execute template processing: ", err)
	}
}
