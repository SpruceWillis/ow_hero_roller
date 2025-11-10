FROM golang AS builder

ARG OS=linux
ARG ARCH=amd64

WORKDIR /usr/src/app
COPY go.* main.go /usr/src/app/
RUN  GOOS=${OS} GOARCH=${ARCH} CGO_ENABLED=0 go build -o /usr/local/bin/ow_hero_roller .

FROM scratch
COPY --from=builder /usr/local/bin/ow_hero_roller /usr/local/bin/ow_hero_roller

ENTRYPOINT [ "/usr/local/bin/ow_hero_roller" ]
