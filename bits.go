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
func (bits *Bits) GetBit(i int) int {
    if (i < 0) || (i >= bits.count) {
        return NilBit
    }

    j, k := i / 8, i % 8
    if (bits.bytes[j] & byte(1 << uint(7-k))) == 0 {
        return 0
    }
    return 1
}

// Get count of bits.
func (bits *Bits) Count() int {
    return bits.count
}

// Dump string representation of bits.
func (bits *Bits) String() string {
    s := make([]byte, bits.count)
    for i := 0; i < bits.count; i++ {
        j, k := i / 8, i % 8
        if (bits.bytes[j] & byte(1 << uint(7-k))) == 0 {
            s[i] = '0'
        } else {
            s[i] = '1'
        }
    }
    return string(s)
}
