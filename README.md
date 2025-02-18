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
	--taste_file_path=$(pwd)/sample.taste \
	--recipe_file_path=$(pwd)/sample.recipe \
	--output_file_path=$(pwd)/out.tmp \
	--override_output

```
# sample.taste

namespace taste.lang;

# Agnostic definition of my data.
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
