all: bin/quantcast
test: unit-test

PLATFORM=local

.PHONY: bin/quantcast
bin/quantcast:
    @docker build . --target bin \
    --output bin/ \
    --platform ${PLATFORM}

.PHONY: unit-test
unit-test:
    @docker build . --target unit-test