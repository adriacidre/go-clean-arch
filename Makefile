BINARY=api_test
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

help: ##@other Show this help
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

build: ##@dev Builds current package
	${TESTS}
	go build -o ${BINARY}

run: ##@dev Run current package
	${TESTS}
	go run main.go

install: ##@dev Installs current package
	${TESTS}
	go build -o ${BINARY}

test: ##@test Runs associated tests
	go test -cover -short $$(go list ./... | grep -v /vendor/)

dep-ensure: ##@dev Installs current package dependencies
	dep ensure

lint-install:: ##@lint Installs necessary packages to run linting tools
	@# The following installs a specific version of golangci-lint, which is appropriate for a CI server to avoid different results from build to build
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $(GOPATH)/bin v1.9.1

lint: ##@lint runs the linter
	@echo "lint"
	@golangci-lint run ./...

rest-fetch: ##@rest Fetch a resource by id
	curl http://localhost:9090/payment/1

rest-add: ##@rest Create a resource
	curl -d '{"payment_id":"supu","organisation_id":"tupu"}' -H "Content-Type: application/json" -X POST http://localhost:9090/payment

rest-update: ##@rest Update a resource
	curl -d '{"payment_id":"supu","organisation_id":"modified"}' -H "Content-Type: application/json" -X PATCH http://localhost:9090/payment/1

rest-list: ##@rest List a collection of payment resources
	curl http://localhost:9090/payment

rest-delete: ##@rest Delete a resource
	curl -X "DELETE" http://localhost:9090/payment/1

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest


# This is a code for automatic help generator.
# It supports ANSI colors and categories.
# To add new item into help output, simply add comments
# starting with '##'. To add category, use @category.
GREEN  := $(shell echo "\e[32m")
WHITE  := $(shell echo "\e[37m")
YELLOW := $(shell echo "\e[33m")
RESET  := $(shell echo "\e[0m")

HELP_FUN = \
		   %help; \
		   while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z0-9\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
		   print "Usage: make [target]\n\n"; \
		   for (sort keys %help) { \
			   print "${WHITE}$$_:${RESET}\n"; \
			   for (@{$$help{$$_}}) { \
				   $$sep = " " x (32 - length $$_->[0]); \
				   print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
			   }; \
			   print "\n"; \
		   }
