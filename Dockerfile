FROM golang:1.9-alpine as builder
WORKDIR /go/src/git.containerum.net/ch/resource-service
COPY . .
RUN CGO_ENABLED=0 go build -v -ldflags="-w -s -extldflags '-static'" -o /bin/resource-service

FROM scratch
COPY --from=builder /bin/resource-service /
COPY --from=builder /go/src/git.containerum.net/ch/resource-service/migrations /migration
ENV MIGRATION_URL="file:///migration" \
    AUTH_ADDR="localhost:1001" \
    BILLING_ADDR="localhost:1002" \
    KUBE_ADDR="localhost:1003" \
    DB_URL="postgres://user:password@localhost:5432/resource_service?sslmode=disable"
EXPOSE 1213
ENTRYPOINT ["/resource-service"]