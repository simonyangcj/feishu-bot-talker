TAG ?=
ifeq ($(TAG),)
	TAG = $(COMMIT_REF)
endif

.PHONY: test-all
test-all:
	ginkgo -r -v --cover --coverprofile=coverage.out

.PHONY: fmt
fmt:
	gofmt -w pkg cmd

.PHONY: go-build
go-build:
	go build -o cmd/feishu-bot-talker/feishu-bot-talker -ldflags "-X 'main.app.version.version=${TAG}'" cmd/feishu-bot-talker/main.go