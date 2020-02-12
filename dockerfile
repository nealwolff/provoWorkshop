#build stage
FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/nealwolff/provoWorkshop

RUN go get github.com/gorilla/mux && \
    go get go.mongodb.org/mongo-driver || true

COPY . .

ENV CGO_ENABLED 0
WORKDIR /go/src/github.com/nealwolff/provoWorkshop/api
RUN go test -c -o tests

RUN CGO_ENABLED=0 GOOS=linux go install github.com/nealwolff/provoWorkshop/api

#final image stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/api .
RUN chmod 755 -R api
CMD ["./api"]
EXPOSE 8080