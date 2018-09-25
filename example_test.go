package systemproxy_test

import (
	"fmt"

	"github.com/andreyvit/systemproxy"
)

func ExampleGet() {
	s, err := systemproxy.Get()
	if err != nil {
		fmt.Printf("** %+v", err)
	} else {
		if s.Enabled {
			fmt.Printf("Status:   on\n")
		} else {
			fmt.Printf("Status:   off\n")
		}
		if s.DefaultServer != "" {
			fmt.Printf("Default:  %s\n", s.DefaultServer)
		}
	}
}

func ExampleSet() {
	err := systemproxy.Set(systemproxy.Settings{
		Enabled:       true,
		DefaultServer: "127.0.0.1:8888",
	})
	if err != nil {
		fmt.Printf("** %+v", err)
	} else {
		fmt.Println("OK")
	}
}
