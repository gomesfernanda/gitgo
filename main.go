package main

import (
	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	flag.Parse()
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
