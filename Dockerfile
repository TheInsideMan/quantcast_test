FROM --platform=${BUILDPLATFORM} golang:1.13.3-alpine AS build
WORKDIR /go/src/QuantCast
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/quantcast .

FROM scratch AS bin-unix
COPY --from=build /out/quantcast /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/quantcast /example.exe

FROM bin-${TARGETOS} AS bin