# go-empty
Check empty library

## Usage
```go
import (
    "github.com/m-hosoi/go-empty/empty"
)
type Data struct {
    Name string
}

func (d Data) IsEmpty() bool {
    return d.Name == ""
}

func ExampleIsEmpty() {
    fmt.Printf("%v\n", empty.IsEmpty(true))
    fmt.Printf("%v\n", empty.IsEmpty("a"))
    fmt.Printf("%v\n", empty.IsEmpty(1))
    fmt.Printf("%v\n", empty.IsEmpty(false))
    fmt.Printf("%v\n", empty.IsEmpty(nil))
    fmt.Printf("%v\n", empty.IsEmpty(int64(0)))
    fmt.Printf("%v\n", empty.IsEmpty(0))
    fmt.Printf("%v\n", empty.IsEmpty(float64(0.0))
    fmt.Printf("%v\n", empty.IsEmpty(""))
    // Output:
    // false
    // false
    // false
    // true
    // true
    // true
    // true
    // true
    // true
}
func ExampleIsNotEmpty() {
    fmt.Printf("%v\n", empty.IsNotEmpty(true))
    fmt.Printf("%v\n", empty.IsNotEmpty("a"))
    fmt.Printf("%v\n", empty.IsNotEmpty(1))
    fmt.Printf("%v\n", empty.IsNotEmpty(false))
    fmt.Printf("%v\n", empty.IsNotEmpty(nil))
    fmt.Printf("%v\n", empty.IsNotEmpty(int64(0)))
    fmt.Printf("%v\n", empty.IsNotEmpty(0))
    fmt.Printf("%v\n", empty.IsNotEmpty(float64(0.0)))
    fmt.Printf("%v\n", empty.IsNotEmpty(""))
    // Output:
    // true
    // true
    // true
    // false
    // false
    // false
    // false
    // false
    // false
}

func ExampleIsEmptyStruct() {
    fmt.Printf("%v\n", empty.IsEmpty(Data{Name: ""}))
    fmt.Printf("%v\n", empty.IsEmpty(&Data{Name: ""}))
    fmt.Printf("%v\n", empty.IsEmpty(Data{Name: "foo"}))
    fmt.Printf("%v\n", empty.IsEmpty(&Data{Name: "foo"}))
    // Output:
    // true
    // true
    // false
    // false
}

func ExampleIIf() {
    fmt.Printf("%v\n", empty.IIf(true, "a", "b"))
    fmt.Printf("%v\n", empty.IIf(false, "a", "b"))
    // Output:
    // a
    // b
}
```

## Example with [Pongo2](https://github.com/flosch/pongo2)
### Go
```go
tmpl, err := pongo2.DefaultSet.FromCache("index.html")
checkErr(err)
ctxt := pongo2.Context{
    "iif":          empty.IIf,
    "is_empty":     empty.IsEmpty,
    "is_not_empty": empty.IsNotEmpty,
    "join": func(ar ...string){
        res := ""
        for _, v := range ar {
            res += v
        }
        return res
    }
    ...
}
out, err := tmpl.Execute(ctxt)
```
### HTML (template)
```html
{% set url = iif(item, join("http://example.org/item/", item.Name, ".html"), "http://example.org/") %}
<a href="{{ url }}">Logo</a>
{% if is_empty(item) %}
<h1>Site Title</h1>
{% else %}
<h1>{{item.Name}} |Site Title</h1>
{% endif %}
```
