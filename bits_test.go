package bitidx

import(
    "testing"
)

func TestBits_GetBit_Full(t *testing.T) {
    b := NewBits([]byte{1, 0x80}, -1)
    f := []int{
        0, 0, 0, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 0, 0, 0,
        NilBit, NilBit,
    }

    for i, end := 0, len(f); i < end; i++ {
        x := b.GetBit(i)
        if x != f[i] {
            t.Error("bit", i, "got", x, ", expected", f[i])
        }
    }
}

func TestBits_GetBit_ZeroBased(t *testing.T) {
    if NewBits([]byte{0x80}, -1).GetBit(0) != 1 {
        t.Error("unexpected")
    }
}

func TestBits_GetBit_OutOfRange(t *testing.T) {
    if NewBits([]byte{0x80}, -1).GetBit(8) != NilBit {
        t.Error("unexpected")
    }
}
