//+build !windows

package systemproxy

/*
Get returns the current systemwide proxy settings.
*/
func Get() (Settings, error) {
	return Settings{}, ErrNotImpl
}

/*
Set updates systemwide proxy settings.
*/
func Set(s Settings) error {
	return ErrNotImpl
}
