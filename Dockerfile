FROM golang AS builder

WORKDIR /usr/src/app
COPY go.* main.go /usr/src/app/
RUN  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /usr/local/bin/ow_hero_roller .

FROM scratch
COPY --from=builder /usr/local/bin/ow_hero_roller /usr/local/bin/ow_hero_roller

ENTRYPOINT [ "/usr/local/bin/ow_hero_roller" ]
