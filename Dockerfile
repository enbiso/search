FROM golang AS builder
WORKDIR /go/src/search
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/search/search .
EXPOSE 8080
ENTRYPOINT [ "/app/search" ]