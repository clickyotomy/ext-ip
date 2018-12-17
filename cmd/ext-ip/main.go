// Command `ext-ip' fetches the client's external IP address.
package main

import (
    "os"
    "flag"
    "fmt"

    "github.com/clickyotomy/ext-ip/resolve"
    "github.com/davecgh/go-spew/spew"
)

var spewConf = spew.ConfigState{Indent: "\t"}

func usage() {
    fmt.Fprintf(os.Stderr, "ext-ip: Fetch external IP address.\n")
}

func main() {
    flag.Usage = usage
    flag.Parse()

    addr, _, err := resolve.ExtIP()

    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }

    fmt.Fprintf(os.Stdout, "%s\n", addr)
}
