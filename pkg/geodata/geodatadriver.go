package geodata

type IGeoDataDriver interface {
	Open() error
	Read() (*GeoPoint, error)
	Close()
}
