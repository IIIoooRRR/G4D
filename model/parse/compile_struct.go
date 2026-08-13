package parse

import "unsafe"

type eface struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}
