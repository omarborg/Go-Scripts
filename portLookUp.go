package main

import (
  "fmt"
  "os"
  "net"
)

func main() {

    if len(os.Args) != 3 {
      fmt.Fprintf(os.Stderr, "Usage: %s Network-type(TCP or UDP) service(https or telnet, etc.)\n", os.Args[0])
      os.Exit(1)
    }

    networkType := os.Args[1]
    service := os.Args[2]

    port, err := net.LookupPort(networkType, service)
    if err != nil {
      fmt.Println("Error: ", err.Error())
      os.Exit(2)
    }

    fmt.Println("\nPort for", service, "is" , port, "\n")
    os.Exit(0)
}
