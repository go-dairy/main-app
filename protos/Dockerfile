FROM bufbuild/buf:latest AS buf


FROM golang:alpine
ENV GO111MODULE=on
RUN go mod init github.com/go-dairy/main-app && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
  && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
  && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf
ENV PATH="/go/bin:${PATH}"
ENTRYPOINT ["/usr/local/bin/buf"]