FROM golang:1.20.3-alpine

ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV PACKAGES="ca-certificates git curl bash zsh make"
ENV ROOT /app

RUN apk update \
  && apk add --no-cache ${PACKAGES} \
  && update-ca-certificates

WORKDIR ${ROOT}

RUN go install github.com/go-delve/delve/cmd/dlv@v1.7.3 && \
  go install golang.org/x/tools/gopls@latest && \
  go install github.com/cweill/gotests/gotests@latest && \
  go install github.com/fatih/gomodifytags@latest && \
  go install github.com/josharian/impl@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
  go install golang.org/x/tools/cmd/goimports@v0.1.10 && \
  go install github.com/alvaroloes/enumer@latest && \
  go install github.com/haya14busa/goplay/cmd/goplay@latest && \
  go install github.com/pressly/goose/v3/cmd/goose@latest && \
  go install github.com/swaggo/swag/cmd/swag@latest

COPY ./ ./

WORKDIR ${ROOT}

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
