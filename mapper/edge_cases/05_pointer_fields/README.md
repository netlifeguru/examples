# 05 Pointer Fields

This example shows how the mapper handles pointer fields.

When a database column contains a value, the mapper allocates a pointer and assigns the value:

```go
Bio *string `db:"bio"`
```

When a database column is `NULL`, the field stays `nil`.

This is useful when you want to distinguish between an empty value and a missing value.