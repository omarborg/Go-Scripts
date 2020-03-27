package main

import (
  "fmt"
  "os"
  "net"
)

func main() {

    if len(os.Args) != 2 {
      fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n" , os.Args[0])
      os.Exit(1)
    }

    dotAddr := os.Args[1]

    addr := net.ParseIP(dotAddr)

    if addr == nil {
      fmt.Println("Invalid IP address")
      os.Exit(1)
    }

    mask := addr.DefaultMask()
    network := addr.Mask(mask)
    ones, bits := mask.Size()

    fmt.Println("\n Address is:", addr.String(),"\n",
    "Default mask length is: ", ones ,"\n",
    "Leading one count is: ", bits,"\n",
    "hex mask is: ", mask.String(),"\n",
    "Network is: ", network.String(),"\n")

      os.Exit(0)

}
