# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY serving-operator/go.mod serving-operator/go.mod
COPY serving-operator/go.sum serving-operator/go.sum

# Copy the go source
COPY proto/ proto/
COPY ray-operator/ ray-operator/
COPY serving-operator/ serving-operator/

WORKDIR /workspace/serving-operator

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/serving-operator/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
