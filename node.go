package bitidx

import (
    "encoding/json"
)

type Node []interface{}

const (
    PutNone = 0
    PutUpdated = 1
    PutOverwritten = 2
)

// New Node with a default value.
func NewNode(defaultValue interface{}) Node {
    return Node{defaultValue, defaultValue}
}

// PutNone: no update is performed
// PutUpdated: updated
// PutOverwritten: overwritten
func (n Node) Put(bits *Bits, content interface{}, overwrite bool) int {
    p, x := n, bits.GetBit(0)
    if x == NilBit {
        return PutNone
    }

    for i, count := 1, bits.Count(); i < count; i++ {
        switch v := p[x].(type) {
        case Node:
            p = v
        default:
            q := Node{v, v}
            p[x] = q
            p = q
        }
        x = bits.GetBit(i)
    }

    if _, ok := p[x].(Node); !ok {
        p[x] = content
        return PutUpdated
    } else if overwrite {
        p[x] = content
        return PutOverwritten
    }

    return PutNone
}

// For the returned node and id, both or one of them must be nil.
func (n Node) Find(bits *Bits) (node Node, content interface{}) {
    for p, i := n, 0; ; i++ {
        x := bits.GetBit(i)
        if x == NilBit {
            break
        }
        switch v := p[x].(type) {
        case Node:
            p = v
            node = v
        default:
            return nil, v
        }
    }
    return node, nil
}

// When the structure is imported from json. A user function f is supplied to convert none structure objects.
func (n Node) Consolidate(f func(interface{}) interface{}) {
    n.consolidate(0, f)
    n.consolidate(1, f)
}

// Consolidate float64 numbers when node is imported from json.
func (n Node) ConsolidateNum() {
    n.Consolidate(func(x interface{}) interface{} {
        if v, ok := x.(float64); ok {
            return int(v)
        }
        return x
    })
}

func (n Node) consolidate(i int, f func(interface{}) interface{}) {
    switch v := n[i].(type) {
    case Node:
        v.Consolidate(f)
    case []interface{}:
        if len(v) == 2 {
            vn := Node(v)
            vn.Consolidate(f)
            n[i] = vn
        } else {
            n[i] = f(v)
        }
    default:
        n[i] = f(v)
    }
}

// Serialize to string in JSON.
func (n Node) String() string {
    v, _ := json.Marshal(n)
    return string(v)
}
