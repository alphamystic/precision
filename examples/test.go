package main

import (
  "fmt"
  "github.com/alphamystic/precision"
)

func main(){
  s := "789.908765"
  d := precision.NewDecimal(s)
  fmt.Printf(d)
}
