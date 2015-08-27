package bitidx

import(
    "testing"
)

func TestNote_Put(t *testing.T) {
    n := NewNode(0)

    if a := n.Put(NewBits([]byte{1, 2}, -1), 10, false); a != PutUpdated {
        t.Error("Node->Put not expected", a)
    }

    if a := n.Put(NewBits([]byte{1, 2}, 15), 10, false); a != PutNone {
        t.Error("Node->Put not exptected", a)
    }

    if a := n.Put(NewBits([]byte{1, 2, 3}, -1), 12, false); a != PutUpdated {
        t.Error("Node->Put not exptected", a)
    }

    if n.String() != "[[[[[[[[0,[[[[[[[0,[[[[[[[[10,[10,12]],10],10],10],10],10],10],0]],0],0],0],0],0],0]],0],0],0],0],0],0],0]" {
        t.Error("structure not expected", s)
    }

    if _, f := n.Find(NewBits([]byte{1, 2, 4}, 24)); f != 10 {
        t.Error("Node->Find failed", f)
    }

    if _, f := n.Find(NewBits([]byte{1, 2, 3}, 24)); f != 12 {
        t.Error("Node->Find failed", f)
    }

    if q, _ := n.Find(NewBits([]byte{1, 2}, 16)); q != nil {
        if q.String() != "[[[[[[[10,[10,12]],10],10],10],10],10],10]" {
            t.Error("structure not expected", s)
        }
    } else {
        t.Error("not found")
    }
}
