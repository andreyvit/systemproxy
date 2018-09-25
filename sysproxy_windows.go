package systemproxy

import (
	"golang.org/x/sys/windows/registry"
)

const (
	keyProxyEnable = "ProxyEnable"
	keyProxyServer = "ProxyServer"
)

func Get() (Settings, error) {
	var s Settings

	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		return s, proxyErr{err}
	}
	defer k.Close()

	en, _, err := k.GetIntegerValue(keyProxyEnable)
	if err != nil && err != registry.ErrNotExist {
		return s, proxyErr{err}
	}
	s.Enabled = (en != 0)

	s.Server, _, err = k.GetStringValue(keyProxyServer)
	if err != nil && err != registry.ErrNotExist {
		return s, proxyErr{err}
	}

	return s, nil
}

func Set(s Settings) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return proxyErr{err}
	}
	defer k.Close()

	var en uint32
	if s.Enabled {
		en = 1
	}
	err = k.SetDWordValue(keyProxyEnable, en)
	if err != nil {
		return proxyErr{err}
	}

	if s.Server != "" {
		err = k.SetStringValue(keyProxyServer, s.Server)
		if err != nil {
			return proxyErr{err}
		}
	}

	return nil
}

type proxyErr struct {
	cause error
}

func (e proxyErr) Error() string {
	return "setting proxy settings: " + e.cause.Error()
}
