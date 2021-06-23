all: bin/quantcast_test
test: lint unit-test

PLATFORM=local

.PHONY: bin/quantcast_test
bin/quantcast_test:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out