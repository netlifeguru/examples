# 08 Many To Many

This example shows how to scan rows from a many-to-many relationship.

The query joins `posts`, `post_tags`, and `tags`, then uses SQL aliases to map the result into a flat Go struct.