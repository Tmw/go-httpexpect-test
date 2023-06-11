gotestsumpath := $(shell which gotestsum)
testargs = -v -race -count=1 ./...
.PHONY: test
test:
ifeq ($(gotestsumpath),)
	go test $(testargs)
else
	gotestsum -- $(testargs)
endif