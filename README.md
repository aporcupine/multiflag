# multiflag

Little helper module for creating flags with multiple names (e.g. -r, -R &amp; --recursive)

https://godoc.org/github.com/aporcupine/multiflag

## Usage Example

```go
import "github.com/aporcupine/multiflag"

var (
  recursive = multiflag.Bool(false, "remove recursively", "r", "R", "recursive")
  config = multiflag.String("/etc/myconfig", "location of config file", "c", "config")
)
```
