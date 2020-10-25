FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /build
COPY . .

RUN go build .

FROM golang:1.15.2-alpine3.12

ENV USER_SERVICE_BASE_URL http://localhost:8000
ENV JWT_SIGNING_KEY MockSigningKey
ENV PORT 8010
ENV APPINSIGHTS_INSTRUMENTATIONKEY=Empty

WORKDIR /app
COPY --from=builder /build/auth-service .

CMD [ "./auth-service" ]