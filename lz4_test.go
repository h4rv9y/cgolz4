package cgolz4_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/h4rv9y/cgolz4"
)

var (
	testdataMTTS = mustReadFile("testdata/Mark.Twain-Tom.Sawyer.txt")
)

func TestCompressAndDecompress(t *testing.T) {
	cmpBuf := make([]byte, cgolz4.CompressBound(len(testdataMTTS)))

	n, err := cgolz4.Compress(testdataMTTS, cmpBuf)
	if err != nil {
		t.Error(err)
	}

	rawBuf := make([]byte, len(testdataMTTS))
	n, err = cgolz4.Decompress(cmpBuf[:n], rawBuf)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(testdataMTTS, rawBuf) {
		t.Errorf("decompressed data not match")
	}
}

func mustReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}
