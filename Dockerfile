FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o octoptimist .

FROM alpine as build
WORKDIR /out
COPY --from=builder /app/octoptimist /out/octoptimist
RUN chmod +x /out/octoptimist
RUN mkdir -m777 -p /out/static
RUN mkdir -m777 -p /out/static/css
RUN mkdir -m777 -p /out/private

COPY /static/. /out/static/
COPY /static/css/. /out/static/css
COPY /private/. /out/private/

ENV CLIENT_ID=""
ENV CLIENT_SECRET=""
ENV REDIRECT_URL=""
ENV OCTOPOD_DOMAIN=""
ENV PORT="9090"

EXPOSE ${PORT}
WORKDIR /out
ENTRYPOINT ["/out/octoptimist"]