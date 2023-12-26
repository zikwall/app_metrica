package geolocation

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/oschwald/geoip2-golang"
)

type Wrapper struct {
	name   string
	reader *geoip2.Reader
}

func NewWrapper(reader *geoip2.Reader, name string) *Wrapper {
	w := &Wrapper{reader: reader, name: name}
	return w
}

func (w *Wrapper) CityInfo(ip net.IP) (*geoip2.City, error) {
	data, err := w.Reader().City(ip)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (w *Wrapper) Reader() *geoip2.Reader {
	return w.reader
}

func (w *Wrapper) Drop() error {
	return w.reader.Close()
}

func (w *Wrapper) DropMsg() string {
	return fmt.Sprintf("close %s reader", w.name)
}

func Reader(basePath string) (*Wrapper, error) {
	path := filepath.Join(basePath, "GeoIP2-City.mmdb")
	abs, _ := filepath.Abs(path)

	geoReaderCity, err := geoip2.Open(abs)
	if err != nil {
		return nil, err
	}

	return NewWrapper(geoReaderCity, "city"), nil
}

func ReaderASN(basePath string) (*Wrapper, error) {
	path := filepath.Join(basePath, "GeoLite2-ASN.mmdb")
	abs, _ := filepath.Abs(path)

	geoReaderCity, err := geoip2.Open(abs)
	if err != nil {
		return nil, err
	}

	return NewWrapper(geoReaderCity, "city"), nil
}
