# Start by building the application.
FROM golang:1.12 as build
LABEL maintainer="vikram.jakhr@gmail.com"
LABEL go_version="1.12"
LABEL description="Hiking simulator application"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
COPY input.yaml /
CMD ["/app"]