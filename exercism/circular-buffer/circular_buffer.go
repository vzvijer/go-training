package circular

import "errors"

type Buffer struct {
	data  []byte
	size  int
	index int
	len   int
}

func NewBuffer(size int) *Buffer {
	buff := Buffer{}
	buff.data = make([]byte, size)
	buff.size = size
	return &buff
}

func (buff *Buffer) Reset() {
	buff.index = 0
	buff.len = 0
}

func (buff *Buffer) ReadByte() (byte, error) {
	if buff.len == 0 {
		return 0, errors.New("Error: Reading from an empty buffer!")
	}
	b := buff.data[buff.index]
	buff.index = (buff.index + 1) % buff.size
	buff.len--
	return b, nil
}

func (buff *Buffer) WriteByte(b byte) error {
	if buff.len == buff.size {
		return errors.New("Error: Writing to a full buffer!")
	}
	writeIndex := (buff.index + buff.len) % buff.size
	buff.data[writeIndex] = b
	buff.len++
	return nil
}

func (buff *Buffer) Overwrite(b byte) {
	if buff.len == buff.size {
		buff.index = (buff.index + 1) % buff.size
		buff.len--
	}
	buff.WriteByte(b)
}
