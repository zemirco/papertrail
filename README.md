
# papertrail

Send logs from your Golang apps to Papertrail.

## Example

```go
package main

import (
  "log"
  "github.com/zemirco/papertrail"
)

func main() {
  writer := papertrail.Writer{
    Port: 12345,
    Network: papertrail.UDP,
  }

  // use writer directly
  n, err := writer.Write([]byte("writer\n"))
  if err != nil {
    panic(err)
  }
  fmt.Printf("number of bytes written: %d\n", n)

  // or create a new logger
  logger := log.New(&writer, "", log.LstdFlags)
  logger.Print("logger")
}
```

## License

MIT
