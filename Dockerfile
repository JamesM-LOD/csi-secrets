FROM golang:1.13-alpine
RUN apk update

RUN mkdir /csi-secrets
WORKDIR /csi-secrets
COPY . .
RUN go build -o /go/bin/csi-secrets

FROM alpine:latest
COPY --from=0 /go/bin/csi-secrets .
EXPOSE 80
CMD ["./csi-secrets"]
