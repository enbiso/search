FROM golang AS builder
WORKDIR /go/src/search
COPY . .
RUN go get -d .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/search/search .
ENTRYPOINT [ "/app/search" ]