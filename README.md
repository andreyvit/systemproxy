systemproxy (Go)
================

[![GoDoc](https://godoc.org/github.com/andreyvit/systemproxy?status.svg)](https://godoc.org/github.com/andreyvit/systemproxy)

Gets and sets systemwide proxy settings. Currently only implemented on Windows.

Get:

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

Set:

    err := systemproxy.Set(systemproxy.Settings{
        Enabled:       true,
        DefaultServer: "127.0.0.1:8888",
    })
