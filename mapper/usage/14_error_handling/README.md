# 14 Error Handling

This example shows how the mapper reports conversion errors.

The `created_at` column contains a datetime value, but the destination struct uses an `int` field. The mapper returns an error instead of silently assigning an invalid value.

Mapper returned error:
sql: Scan error on column index 2, name "created_at": converting driver.Value type time.Time ("2026-05-20 16:03:35 +0000 UTC") to a int: invalid syntax