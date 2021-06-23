FROM --platform=${BUILDPLATFORM} golang:1.15.2-alpine AS base
WORKDIR ./src/QuantCast
ENV CGO_ENABLED=0
COPY . ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/cookie_parser .

FROM base AS unit-test
RUN mkdir -p /out && go test -v -coverprofile=/out/cover.out ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM scratch AS bin-unix
COPY --from=build /out/cookie_parser /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/cookie_parser /cookie_parser.exe

FROM bin-${TARGETOS} as bin