package ipidx

type Node []interface{}

const NilId = 0

const (
    AssignAborted = -1
    AssignNone = 0
    AssignUpdated = 1
    AssignOverwritten = 2
)

func NewNode() Node {
    return Node{NilId, NilId}
}

// AssignNone: no update is performed
// AssignUpdated: updated
// AssignOverwritten: overwritten
// AssignAborted: the node tree is not as expected, error or panic
func (n Node) Assign(bits Bits, length int, id int, overwrite bool) int {
    p, x := n, bits.GetBit(0)
    for i := 1; i < length; i++ {
        switch v := p[x].(type) {
        case Node:
            p = v
        case int:
            q := Node{v, v}
            p[x] = q
            p = q
        default:
            return AssignAborted
        }
        x = bits.GetBit(i)
    }

    if _, ok := p[x].(int); ok {
        p[x] = id
        return AssignUpdated
    } else if overwrite {
        p[x] = id
        return AssignOverwritten
    }

    return AssignNone
}

func (n Node) Find(bits Bits, length int) int {
    for p, i := n, 0; i < length; i++ {
        x := bits.GetBit(i)
        if x == NilBit {
            break
        }
        switch v := p[x].(type) {
        case Node:
            p = v
        case int:
            return v
        default:
            return NilId
        }
    }
    return NilId
}

func (n Node) FindNode(bits Bits, length int) (node Node) {
    for p, i := n, 0; i < length; i++ {
        x := bits.GetBit(i)
        if x == NilBit {
            break
        }
        switch v := p[x].(type) {
        case Node:
            p = v
            node = v
        default:
            return nil
        }
    }
    return
}

// When the structure is import json, array should be translated to Node and float to int.
func (n Node) Consolidate() {
    n.consolidate(0)
    n.consolidate(1)
}

func (n Node) consolidate(i int) {
    switch v := n[i].(type) {
    case Node:
        v.Consolidate()
    case []interface{}:
        vn := Node(v)
        vn.Consolidate()
        n[i] = vn
    case float64:
        n[i] = int(v)
    }
}
