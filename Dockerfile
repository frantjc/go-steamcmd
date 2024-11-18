FROM golang:1.23 AS build
WORKDIR $GOPATH/github.com/frantjc/go-steamcmd
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN go build -o /app-info-print ./examples/app_info_print
RUN go build -o /app-update ./examples/app_update
RUN go build -o /steamcmdw ./cmd/steamcmdw

FROM debian
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        lib32gcc-s1 \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build \
    # /app-info-print \
    /app-update \
    /steamcmdw \
    /usr/local/bin/
RUN steamcmdw +quit
ENTRYPOINT ["steamcmdw"]
