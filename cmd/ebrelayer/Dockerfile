#
# Build
#
FROM golang:1.18 AS build

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

# Empty dir for the db data
RUN mkdir /data

WORKDIR /grid
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make install

#
# Main
#
FROM alpine

# Copy the smart contract ABIs over.
ADD cmd/ebrelayer/contract/generated  /grid/cmd/ebrelayer/contract/generated

# Copy the compiled binaires over.
COPY --from=build /go/bin/ebrelayer /usr/bin/ebrelayer
COPY --from=build /go/bin/gridnoded /usr/bin/gridnoded