# Experiment 1

Demonstrates a race condition in HTTP middleware where a shared closure variable is modified
concurrently by different requests. The bug occurs because multiple goroutines (HTTP requests)
share and modify the same variable pointer through the closure environment.
