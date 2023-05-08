FROM golang:1.19 as base

ENV GO111MODULE=on

WORKDIR /cosmo

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

FROM base as builder

# Copy the go source
COPY cmd/ cmd/
COPY api/ api/
COPY pkg/ pkg/
COPY proto/ proto/
COPY internal/ internal/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-server ./cmd/auth-server/main.go

FROM node:18-alpine as ui-builder

# Create build environment
ENV PATH web/auth-server-ui/node_modules/.bin:$PATH
RUN mkdir -p web/auth-server-ui
WORKDIR /cosmo/web/auth-server-ui
# Copy the package.json and install
COPY web/auth-server-ui/package.json package.json
COPY web/auth-server-ui/tsconfig.json tsconfig.json
COPY web/auth-server-ui/yarn.lock yarn.lock
RUN yarn install

# Copy the source and build
COPY ./web/auth-server-ui .
RUN yarn build --base=/cosmo-auth-server

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /cosmo/auth-server .
COPY --from=ui-builder /cosmo/web/auth-server-ui/dist ./public

USER 65532:65532

ENTRYPOINT ["/app/auth-server"]
