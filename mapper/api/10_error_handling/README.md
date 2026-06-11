# 11 Error Handling

This example shows how to handle mapper errors.

## No rows

`ScanStructOne` returns `mapper.ErrNoRows` when the result set is empty.

```go
if errors.Is(err, mapper.ErrNoRows) {
// handle not found
}
```

## Too many rows

`ScanStructOne` returns `mapper.ErrTooManyRows` when the result set contains more than one row.

```go
if errors.Is(err, mapper.ErrTooManyRows) {
// handle unexpected multiple rows
}
```

## Conversion errors

When a database value cannot be assigned to the destination field, the mapper returns a conversion error.

For example, scanning a datetime column into an int field returns an error instead of silently assigning an invalid
value.

Output:

```text
go run .
Example 1: no rows
Handled: mapper: no rows

Example 2: too many rows
Handled: mapper: too many rows

Example 3: conversion error
Handled conversion error:
sql: Scan error on column index 2, name "created_at": converting driver.Value type time.Time ("2026-05-20 16:03:35 +0000 UTC") to a int: invalid syntax
```