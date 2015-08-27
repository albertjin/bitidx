package bitidx

type Bits struct {
    bytes []byte
    count int
}

const NilBit = -1

// New *Bits with a byte array. When count is -1, the whole byte array is used.
func NewBits(bytes []byte, count int) *Bits {
    if count == -1 {
        count = len(bytes) * 8
    } else if count > (len(bytes)*8) {
        return nil
    }

    return &Bits{bytes, count}
}

// Get bit at position i, zero based. When the position i is out of range, it returns NilBit.
func (b *Bits) GetBit(i int) int {
    if (i < 0) || (i >= b.count) {
        return NilBit
    }

    j, k := i / 8, i % 8
    if (b.bytes[j] & byte(1 << uint(7-k))) == 0 {
        return 0
    }
    return 1
}

// Bit count
func (b *Bits) Count() int {
    return b.count
}
