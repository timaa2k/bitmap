package bitmap

import "errors"

type Bitmap struct {
	size uint64
	data []byte
}

const (
	EXCEEDS = "Requested bit exceeds bitmap"
)

func NewBitmap(size uint64) *Bitmap {
	length := uint64(size/8) + 1
	if size%8 == 0 {
		length = uint64(size / 8)
	}
	return &Bitmap{size, make([]byte, length)}
}

func (b *Bitmap) Set(bit uint64) error {
	if bit > b.size {
		return errors.New(EXCEEDS)
	}
	b.data[uint64(bit/8)] |= 1 << (bit % 8)
	return nil
}

func (b *Bitmap) Unset(bit uint64) error {
	if bit > b.size {
		return errors.New(EXCEEDS)
	}
	b.data[uint64(bit/8)] &^= 1 << (bit % 8)
	return nil
}

func (b *Bitmap) IsSet(bit uint64) (bool, error) {
	if bit > b.size {
		return false, errors.New(EXCEEDS)
	}
	return (b.data[uint64(bit/8)]>>(bit%8))&1 == 1, nil
}

func (b *Bitmap) Size() uint64 {
	return b.size
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, len(b.data))
}
