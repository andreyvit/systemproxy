package main

import (
	"flag"
	"log"
	"net"

	"github.com/andreyvit/systemproxy"
)

var onoff = map[bool]string{false: "off", true: "on"}

func main() {
	log.SetFlags(0)
	flag.Parse()

	s, err := systemproxy.Get()
	if err != nil {
		log.Fatalf("** %+v", err)
	}

	if flag.NArg() > 1 {
		log.Fatal("** too many arguments")
	} else if flag.NArg() == 1 {
		server := flag.Arg(0)
		if server == "off" {
			s.Enabled = false
		} else if server == "on" {
			s.Enabled = true
		} else {
			_, _, err := net.SplitHostPort(server)
			if err != nil {
				log.Fatalf("** invalid host:port: %q", server)
			}

			s.Enabled = true
			s.DefaultServer = server
		}
		err := systemproxy.Set(s)
		if err != nil {
			log.Fatalf("** %+v", err)
		}
	}

	s, err = systemproxy.Get()
	if err != nil {
		log.Fatalf("** %+v", err)
	}

	log.Printf("Status:   %s", onoff[s.Enabled])
	if s.DefaultServer != "" {
		log.Printf("Default:  %s", s.DefaultServer)
	}
}
