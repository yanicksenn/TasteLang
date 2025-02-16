# Taste

Taste is a small tool to generate programming language agnostic data structures.

NOTE: This is work in progress and does not fully work yet :)

## Build

```
bazel build //go:main
```

## Run 

```
bazel run //go:main -- \
	--taste=$(pwd)/sample.taste \
	--recipe=$(pwd)/sample.recipe
```

```
# sample.taste

namespace taste.lang;

type MyType {
    string name;
    int counter;
    bool is_active;
}
```

```
# sample.recipe

{{ $namespace := .Namespace }}

{{range .Types}}- {{ $namespace }}.{{.Name}}
{{range .Fields}}  + {{ .Name}}:{{ .Type }}
{{end}}
{{end}}
```

## Output

```
- taste.lang.MyType
  + name:string
  + counter:int
  + is_active:bool
```
