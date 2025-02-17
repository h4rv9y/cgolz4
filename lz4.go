package cgolz4

// #cgo CFLAGS: -O3 -Wno-deprecated-declarations
// #include "src/lz4.h"
// #include "src/lz4.c"
import "C"
import (
	"fmt"
	"unsafe"
)

func p(in []byte) *C.char {
	if len(in) == 0 {
		return (*C.char)(unsafe.Pointer(nil))
	}
	return (*C.char)(unsafe.Pointer(&in[0]))
}

func CompressBound(srcLen int) int {
	return srcLen + ((srcLen / 255) + 16)
}

func Compress(src []byte, dst []byte) (n int, err error) {
	outSize := int(C.LZ4_compress_default(p(src), p(dst), C.int(len(src)), C.int(len(dst))))
	if outSize == 0 {
		return 0, fmt.Errorf("insufficient space for compression")
	}
	return outSize, nil
}

func Decompress(src []byte, dst []byte) (n int, err error) {
	outSize := int(C.LZ4_decompress_safe(p(src), p(dst), C.int(len(src)), C.int(len(dst))))
	if outSize == 0 {
		return 0, fmt.Errorf("insufficient space for decompression")
	}
	return outSize, nil
}
