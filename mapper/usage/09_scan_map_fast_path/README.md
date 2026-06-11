# 09 ScanMap Fast Path

This example shows how to scan rows into `map[string]any` and manually map values using a custom `ScanMap` method.

This can be useful when you want full control over conversions and validation without relying on reflection-based struct assignment.