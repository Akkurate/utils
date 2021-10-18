package compress

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/Akkurate/utils/logging"
)

// MakeGzip MakeGzip
func MakeGzip(input []byte) bytes.Buffer {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(input); err != nil {
		logging.Fatal("%v", err)
	}
	if err := gz.Close(); err != nil {
		logging.Fatal("%v", err)
	}

	return b
}

// UnGzip UnGzip
func UnGzip(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)
	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	resData = resB.Bytes()
	return resData, nil
}
