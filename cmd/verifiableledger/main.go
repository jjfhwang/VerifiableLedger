// cmd/verifiableledger/main.go
package main

import (
"flag"
"log"
"os"

"verifiableledger/internal/verifiableledger"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := verifiableledger.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
