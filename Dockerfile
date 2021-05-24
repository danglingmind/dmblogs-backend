FROM golang:latest as base
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download 
RUN go mod vendor
COPY . /app 

FROM base as build 
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o app .


FROM base as unit-test
RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v .

FROM alpine:latest
COPY --from=build /app/app .
EXPOSE 8000:8000

CMD [ "./app" ]