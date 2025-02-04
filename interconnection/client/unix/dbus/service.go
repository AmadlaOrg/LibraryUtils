package dbus

// NewDBusService
func NewDBusService() IDBus {
	return &SDBus{}
}
