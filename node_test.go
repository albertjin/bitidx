package ipidx

import(
    "testing"
    "encoding/json"
)

func TestNote_Assign(t *testing.T) {
    n := NewNode()

    if a := n.Assign(Bits{1, 2}, 16, 10, false); a != AssignUpdated {
        t.Error("Node->Assign not expected", a)
    }

    if a := n.Assign(Bits{1, 2}, 15, 10, false); a != AssignNone {
        t.Error("Node->Assign not exptected", a)
    }

    if a := n.Assign(Bits{1, 2, 3}, 24, 12, false); a != AssignUpdated {
        t.Error("Node->Assign not exptected", a)
    }

    {
        v, _ := json.Marshal(n)
        s := string(v)
        if s != "[[[[[[[[0,[[[[[[[0,[[[[[[[[10,[10,12]],10],10],10],10],10],10],0]],0],0],0],0],0],0]],0],0],0],0],0],0],0]" {
            t.Error("structure not expected", s)
        }
    }

    if _, f := n.Find(Bits{1, 2, 4}, 24); f != 10 {
        t.Error("Node->Find failed", f)
    }

    if _, f := n.Find(Bits{1, 2, 3}, 24); f != 12 {
        t.Error("Node->Find failed", f)
    }

    {
        q, _ := n.Find(Bits{1, 2}, 16)
        v, _ := json.Marshal(q)
        s := string(v)
        if s != "[[[[[[[10,[10,12]],10],10],10],10],10],10]" {
            t.Error("structure not expected", s)
        }
    }
}
