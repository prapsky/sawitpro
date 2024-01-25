

.PHONY: clean all init generate generate_mocks generate_user_service_mocks generate_auth_service_mocks

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf generated

init: generate
	go mod tidy
	go mod vendor

test:
	go test -short -coverprofile coverage.out -v ./...

generate: generated generate_mocks generate_user_service_mocks generate_auth_service_mocks

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package generated -generate types,server,spec $< > generated/api.gen.go

INTERFACES_GO_FILES := $(shell find repository -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

USER_SERVICE_GO_FILES := $(shell find service -name "user_service.go")
USER_SERVICE_GEN_GO_FILES := $(USER_SERVICE_GO_FILES:%.go=%.mock.gen.go)

generate_user_service_mocks: $(USER_SERVICE_GEN_GO_FILES)
$(USER_SERVICE_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

AUTH_SERVICE_GO_FILES := $(shell find service -name "auth_service.go")
AUTH_SERVICE_GEN_GO_FILES := $(AUTH_SERVICE_GO_FILES:%.go=%.mock.gen.go)

generate_auth_service_mocks: $(AUTH_SERVICE_GEN_GO_FILES)
$(AUTH_SERVICE_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))
