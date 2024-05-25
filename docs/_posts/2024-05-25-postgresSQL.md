---
layout: post
title:  PostgreSQL
date:   2024-05-25 20:31:00 +0000
categories: gaming database glm
---
## PostgreSQL
So I'm on a Chromebook, therefore the localhost in the Debian container is not exposed to the network. Setting a `\password` though was easy enough, and allows the use of `pgadmin4` to construct the schema. From there the schema can be dumped with `pg_dump --schema-only > schema.sql` and `sqlc generate` can generate the wrapper code as part of `gob.sh` (the build script).

I left the example code's SSL connection method as it's bad to forget it later if the database server migrates off the localhost into a sharded cluster. Umm, prepositions seem to be adverbs of a kind. I also decided that the `OpenConnection()` function should not pass on errors but internally manage them or `Fatal()`, and also return a function closure of the connection `close()` function for an easy `defer ... ()` application in an outer function. Nice.

I'll have to look at the docs to see if foreign key indexes are in a sense merged for efficient columnar data storage. The old "file" per column cache locality mechanisms.

## Grown Language Models
The preposition that all current popular LLMs have an intrinsic ontological mismatch between the complexity of the training data and the connectedness of the simplicity of the evolved instinctual "central" thalamic cortex controller. If you take the coordinative cerebellum, as a highly compact 70% of all brain cells as the residual intent corrector to simplify the core textual "grunt" constraints. Umm, in deeds.