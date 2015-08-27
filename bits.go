package bitidx

type Bits []byte

const NilBit = -1

// Get bit at position i, zero based. When the position i is out of range, it returns NilBit.
func (b Bits) GetBit(i int) int {
    j, k := i / 8, i % 8
    if j >= len(b) {
        return NilBit
    }
    if (b[j] & byte(1 << uint(7-k))) == 0 {
        return 0
    }
    return 1
}
