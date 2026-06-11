# 04 Type Conversions

This example shows common value conversions handled by the mapper.

The database may return values in driver-specific forms, for example:

```txt
[]byte -> string
TINYINT(1) / 0 / 1 -> bool
time.Time -> string
```

The mapper tries to assign compatible values to the destination struct fields.

In this example, `created_at` is scanned into a `string` field and `active` is scanned into a `bool` field.