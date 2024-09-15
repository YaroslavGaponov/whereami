package geodatadriver

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/YaroslavGaponov/whereami/pkg/geodata"
)

type GeoDataDriver struct {
	fileName string
	file     *zip.ReadCloser
	reader   io.ReadCloser
	scanner  *bufio.Scanner
}

func New(fileName string) *GeoDataDriver {
	return &GeoDataDriver{
		fileName: fileName,
	}
}

func (driver *GeoDataDriver) Open() error {
	fileNames := strings.Split(driver.fileName, "@")
	if len(fileNames) != 2 {
		return errors.New("filename is incorrect")
	}

	var err error

	driver.file, err = zip.OpenReader(fileNames[0])
	if err != nil {
		return err
	}

	var file *zip.File
	for _, f := range driver.file.File {
		if f.Name == fileNames[1] {
			file = f
			break
		}
	}
	if file == nil {
		return fmt.Errorf("csvfile %s is not found", fileNames[1])
	}

	driver.reader, err = file.Open()
	if err != nil {
		return err
	}
	driver.scanner = bufio.NewScanner(driver.reader)
	driver.scanner.Scan()
	return nil
}

func (driver *GeoDataDriver) Read() (*geodata.GeoPoint, error) {
	if driver.scanner.Scan() {
		line := driver.scanner.Text()
		parts := strings.Split(line, ",")
		for idx, part := range parts {
			parts[idx] = strings.Trim(part, "\"")
		}
		p := geodata.GeoPoint{}
		p.City = parts[0]
		p.CityAscii = parts[1]
		lat, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, err
		}
		p.Lat = lat
		lng, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			return nil, err
		}
		p.Lng = lng
		p.Country = parts[4]
		p.Id = parts[10]
		return &p, nil
	}
	return nil, io.EOF
}

func (driver *GeoDataDriver) Close() {
	driver.reader.Close()
	driver.file.Close()
}
