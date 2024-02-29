FROM golang:alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash make

COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/school-project-server ./cmd/school-project-server/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/school-project-server ./
COPY ./config/local.yaml ./config/local.yaml

CMD ["./school-project-server"]
