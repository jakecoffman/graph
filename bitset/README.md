This package implements some features similar to C++'s bitset.

## Why not make bitset a type?

Often you might need to do:

```
result := (bitset1 | bitset2) & bitset3
```

If bitset was a type, you'd have to cast quite a bit to achieve this:

```
result := Bitset(uint64(bitset1) | uint64(bitset2)) & uint64(bitset3))
```

It would be tempting then to make methods `And` and `Or`, etc but this also hurts readability.

```
result := (bitset1.Or(bitset2)).And(bitset3)
```
