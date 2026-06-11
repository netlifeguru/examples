# 05 Fill From Map

This example shows how to fill a Go struct from `map[string]any`.

```go
mapper.FillFromMap(&user, row)
```

`FillFromMap` uses the same struct mapping rules as the SQL row scanner, including db tags and type conversion.

This is useful when your data already exists as a map, for example from dynamic query results, decoded payloads, cached
data, or generic event data.