# Bit Index for golang

This is a general structue to build index for bit array. It was initially created for query of IP address range.

It is a recursive pair structure. This structure has a simple serialization to json. Here goes an example.

```
map:
  0b110 -> 1
  0b101 -> 2

serialized to json:
  [NilId, [[NilId, 2], [1, NilId]]]
```
