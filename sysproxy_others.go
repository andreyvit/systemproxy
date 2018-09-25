//+build !windows

package systemproxy

func Get() (Settings, error) {
	return Settings{}, ErrNotImpl
}

func Set(s Settings) error {
	return ErrNotImpl
}
