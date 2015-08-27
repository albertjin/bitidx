package bitidx

import(
    "testing"
)

func TestBits_GetBit(t *testing.T) {
    b := Bits([]byte{1, 0x80})
    f := []int{
        0, 0, 0, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 0, 0, 0,
        -1, -1,
    }
    for i, end := 0, len(f); i < end; i++ {
        x := b.GetBit(i)
        if x != f[i] {
            t.Error("bit", i, "got", x, ", expected", f[i])
        }
    }
}
