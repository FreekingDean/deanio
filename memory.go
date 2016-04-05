package deanio

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

var counter int

func ReadMem(RAM []byte) uint64 {
	var sum uint64
	sum = 0
	for i := 0; i < len(RAM)/4; i++ {
		var buf uint32
		binary.Read(bytes.NewReader(RAM[i*4:(i+1)*4]), binary.LittleEndian, &buf)
		sum += uint64(buf)
	}

	updateMemory(strconv.FormatUint(sum, 10))
	return sum
}
