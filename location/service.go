package location

// NewLocationService to set up the location service
func NewLocationService() ILocation {

	return &SLocation{}
}
