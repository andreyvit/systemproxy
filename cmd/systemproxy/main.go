package main

import (
	"flag"
	"log"

	"github.com/andreyvit/systemproxy"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() > 1 {
		log.Fatal("** too many arguments")
	} else if flag.NArg() == 1 {
		server := flag.Arg(0)
		var s systemproxy.Settings
		if server == "off" {
			s.Enabled = false
		} else if server == "on" {
			s.Enabled = true
		} else {
			s.Enabled = true
			s.Server = server
		}
		err := systemproxy.Set(s)
		if err != nil {
			log.Fatalf("** %+v", err)
		}
	}

	s, err := systemproxy.Get()
	if err != nil {
		log.Fatalf("** %+v", err)
	}
	log.Printf("enabled=%v server=%q", s.Enabled, s.Server)
}
