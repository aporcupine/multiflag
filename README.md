# multiflag

Little helper module for creating flags with multiple names (e.g. -r, -R &amp; --recursive)

## Usage Example

```go
var (
  recursive = multiflag.Bool(false, "remove recursively", "r", "R", "recursive")
  config = muliflat.String("/etc/myconfig", "location of config file", "c", "config")
)
```