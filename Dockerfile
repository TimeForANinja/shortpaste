FROM golang:alpine AS backend-build

WORKDIR /usr/local/go/src/git.adyanth.site/adyanth/shortpaste/

RUN apk add --no-cache make build-base
COPY go.* ./
RUN go mod download

COPY *.go ./
COPY cmd cmd
COPY internal/api/resolve_shorts/templates ./templates
RUN CGO_ENABLED=1 go build -o /out/ ./...

FROM node:16.19.1-alpine3.17 as frontend-build

WORKDIR /ui/

COPY ui/package*.json ./
RUN npm install -g @vue/cli && npm install
COPY ui ./
RUN npm run build

FROM alpine

WORKDIR /usr/local/bin/shortpaste/
COPY --from=backend-build /out/shortpaste .
COPY --from=frontend-build /ui/dist/ static/

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/shortpaste/shortpaste" ]
