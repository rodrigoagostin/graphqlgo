FROM golang:1.16 as build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o graphgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/app
COPY --from=0 /go/src/app/graphgo .
CMD ["./graphgo"]