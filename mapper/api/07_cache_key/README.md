# 07 Cache Key

This example shows how to use the optional named scan-plan cache.

```go
mapper.ScanStructRowsWithCacheKey[T](rows, cacheKey, callback)
```

The default ScanStructRows function already uses a safe cache based on the destination type and returned columns.

ScanStructRowsWithCacheKey adds an explicit named cache key. Use it only for stable hot-path queries where the result
shape is expected to stay the same.

The mapper validates the returned columns signature before reusing a named plan.

If the query shape changes, update the cache key version:

```go
const usersListCacheKey = "users:list:created_at_desc:v2"
```