# Bit Index for golang

This is a general structue to build index for bit array. It was initially created for query of IP address range.

It is a recursive binary structure, which can be serialized to json. Here goes an example.

```
map:
  0b110 -> 10
  0b101 -> 20

serialized to json:
  [0, [[0, 20], [10, 0]]]
```
