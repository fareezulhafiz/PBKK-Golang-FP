FROM docker.io/golang:1.23 AS build
WORKDIR /build/
COPY go.mod go.sum .
RUN go mod download
COPY src .
RUN CGO_ENABLED=0 go build -o /app

FROM scratch
COPY --from=build /app /app
COPY html/ /html/
EXPOSE 3000
CMD ["/app"]
