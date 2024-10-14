# Go Templating (gotmpl)

Tool that helps to template string using [Go
template](https://pkg.go.dev/text/template).

## Usage

```text
Usage: gotmpl [--help|--version|<params> <tmpl>]


Examples:
  # Read parameters and template like a string
  gotmpl 'name: John Doe' 'Hello {{ .name }}'

  # Read parameters and template from a file
  gotmpl params.yaml template.gotmpl

  # Read parameters from a file and extract template from another file
  gotmpl \
    /tmp/input.yaml
    <(yq '.spec.template.spec.source.helm.values' application-set.yaml)

  # Read parameters from a Secret on a Kubernetes cluster and extract template from a file
  gotmpl \
    <(kubectl get secret -o yaml -n argocd argocd-cluster-local | yq '. |= pick(["metadata"]) | .metadata |= pick(["annotations"])') \
    <(yq 'select(document_index == 1) | .spec.template.spec.source.helm.values' application-set.yaml)
```

## Author

Jiri Tyr
