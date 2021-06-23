SERVICE_NAME = tiny-url

.PHONY: run
run:
	go run $(PWD)/cmd/$(SERVICE_NAME)

