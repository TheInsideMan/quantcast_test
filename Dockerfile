FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/TheInsideMan/quantcast_test
RUN cd /build && git clone https://github.com/TheInsideMan/quantcast_test/GoTerminal.git
RUN cd /build/GoTerminal/main && go build

ENTRYPOINT [ "/build/GoTerminal/main/main" ]