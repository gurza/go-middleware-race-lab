.PHONY: test

# Usage:
#   make test exp=1 t=TestMiddlewareRace_Racy
#   make test exp=2 t=TestWithResponseWriter
#
# 'exp' selects ./cmd/expN
# 't'   selects the test name regex

test:
	go test -race -count=1 ./cmd/exp$(exp) -run $(t)
