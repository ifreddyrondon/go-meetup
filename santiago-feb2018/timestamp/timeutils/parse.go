package timeutils

// #include "approxidate.h" // HL
// #include <stdlib.h>
// #cgo LDFLAGS: -lm
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func ParseDateString(dt string) (time.Time, error) {
	date := C.struct_timeval{}
	cStr := C.CString(dt)
	ok := C.approxidate(cStr, &date) // HL
	C.free(unsafe.Pointer(cStr))
	if int(ok) != 0 {
		return time.Time{}, fmt.Errorf("Invlid Date Format %s", dt)
	}

	return time.Unix(int64(date.tv_sec), int64(date.tv_usec)*1000), nil
}
