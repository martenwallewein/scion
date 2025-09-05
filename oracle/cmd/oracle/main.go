package main

import (
	"log"

	"github.com/scionproto/scion/oracle"
)

func main() {

	o := oracle.New(":8281")
	log.Fatal(o.Run())
}
