# 08 Schema Version

This example shows how to set a schema version for the optional named scan-plan cache.

```go
mapper.SetSchemaVersion("2026_05_15_001")
```

The schema version becomes part of the named cache key internally.

This is useful when your application uses migrations. After a schema change, update the schema version so the mapper
builds new scan plans for cached query shapes.

You can also read the current value:

```go
mapper.CurrentSchemaVersion()
```

The default ScanStructRows function does not require schema versioning. This is only relevant when using
ScanStructRowsWithCacheKey.