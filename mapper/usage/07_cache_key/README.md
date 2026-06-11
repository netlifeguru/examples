# 07 Cache Key

This example shows how to use the optional named scan-plan cache.

`ScanStructRowsWithCacheKey` can be useful for stable hot-path queries that are executed repeatedly.

The default `ScanStructRows` function already uses a safe cache based on the destination type and returned columns. Use a named cache key only when the query result shape is stable.

If your database schema changes, update the cache key version or call the mapper cache invalidation function.

const usersListCacheKey = "users:list:created_at_desc:v1"

const usersListCacheKey = "users:list:created_at_desc:v2"