FROM golang:1.19-alpine AS builder
WORKDIR /billing-api
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ./ ./
WORKDIR /billing-api/app
RUN CGO_ENABLED=0 go build -o /bin/app


FROM alpine:latest
RUN apk --update add ca-certificates

COPY --from=builder /bin/app /bin/app

ENV PATH=/bin
EXPOSE 9096
CMD [ "/bin/app" ]