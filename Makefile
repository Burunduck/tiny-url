SERVICE_NAME = user-crud-service

.PHONY: run
run:
	go run $(PWD)/cmd/$(SERVICE_NAME)

