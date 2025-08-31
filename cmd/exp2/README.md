# Experiment 2

Demonstrates a race condition when storing `http.ResponseWriter` directly in the request context.
Multiple goroutines accessing the same response writer through context can cause concurrent
read/write operations on the underlying connection, leading to corrupted responses or panics.
