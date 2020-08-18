package binchunk

import "encoding/binary"
import "math"

type reader struct {
	data []byte
}