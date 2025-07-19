// cmd/insurtechframework/main.go
package main

import (
"flag"
"log"
"os"

"insurtechframework/internal/insurtechframework"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := insurtechframework.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
