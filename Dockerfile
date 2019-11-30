FROM golang:1.13.4-alpine3.10 as build
WORKDIR /go/src/app
RUN apk --no-cache add git
COPY . .
RUN go mod download
RUN sh bin/build

FROM alpine:3.7
COPY --from=build /go/src/app/app /usr/local/bin/app
COPY --from=build /go/src/app/templates /usr/local/bin/templates
WORKDIR /usr/local/bin
CMD ["app"]