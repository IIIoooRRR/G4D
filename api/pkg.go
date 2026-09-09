package api

import (
	"unsafe"
)

/*
To reduce the number of allocations,
we’ve made it so that our array, which is stored on the stack, is treated as a string.
This is done to save memory and reduce the load on the garbage collector.
*/

//nolint:gosec
func GetURI(strings ...string) string {
	bytes := make([]byte, 0, 64)
	for i := 0; i < len(strings); i++ {
		bytes = append(bytes, strings[i]...)
	}

	// #nosec G103
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

//nolint:gosec
func GetBigURI(strings ...string) string {
	var bytes = make([]byte, 0, 128)
	for i := 0; i < len(strings); i++ {
		bytes = append(bytes, strings[i]...)
	}
	// #nosec G103
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}
func GetURL(path, uri string) string {
	return path + uri
}
