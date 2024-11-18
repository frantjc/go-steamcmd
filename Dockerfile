FROM golang:1.23 AS build0
WORKDIR $GOPATH/github.com/frantjc/go-steamcmd
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM build0 AS build1
RUN go build -o /steamcmdw ./cmd/steamcmdw

FROM build0 AS build2
RUN go build -o /app-info-print ./examples/app_info_print

FROM build0 AS build3
RUN go build -o /app-update ./examples/app_update

FROM debian AS base
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        lib32gcc-s1 \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build1 /steamcmdw /usr/local/bin/
RUN steamcmdw +quit
ENTRYPOINT ["steamcmdw"]

FROM base
COPY --from=build2 /app-info-print /usr/local/bin/
COPY --from=build3 /app-update /usr/local/bin/
