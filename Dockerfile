FROM golang:1.19 as build

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 make build

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/src/app/stealcamoor /stealcamoor
CMD ["/stealcamoor"]
