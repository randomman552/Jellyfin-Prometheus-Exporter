FROM golang:1.21.5-alpine3.19 AS build
WORKDIR /build
COPY . .
RUN go build -o entrypoint .

FROM alpine:3.19 AS final

COPY --from=build /build/entrypoint /

EXPOSE 2112
ENTRYPOINT [ "/entrypoint" ]
