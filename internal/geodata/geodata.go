package geodata

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type GeoData struct {
	fileName string
	file     *zip.ReadCloser
	reader   io.ReadCloser
	scanner  *bufio.Scanner
}

func New(fileName string) GeoData {
	return GeoData{
		fileName: fileName,
	}
}

func (geodata *GeoData) Open() error {
	fileNames := strings.Split(geodata.fileName, "@")
	if len(fileNames) != 2 {
		return errors.New("filename is incorrect")
	}

	var err error

	geodata.file, err = zip.OpenReader(fileNames[0])
	if err != nil {
		return err
	}

	var file *zip.File
	for _, f := range geodata.file.File {
		if f.Name == fileNames[1] {
			file = f
			break
		}
	}
	if file == nil {
		return fmt.Errorf("csvfile %s is not found", fileNames[1])
	}

	geodata.reader, err = file.Open()
	if err != nil {
		return err
	}
	geodata.scanner = bufio.NewScanner(geodata.reader)
	geodata.scanner.Scan()
	return nil
}

func (geodata *GeoData) Read() (*GeoPoint, error) {
	if geodata.scanner.Scan() {
		line := geodata.scanner.Text()
		parts := strings.Split(line, ",")
		for idx, part := range parts {
			parts[idx] = strings.Trim(part, "\"")
		}
		p := GeoPoint{}
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

func (geodata *GeoData) Close() {
	geodata.reader.Close()
	geodata.file.Close()
}
