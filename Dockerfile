ARG TERRAFORM_VERSION=latest
ARG OPENTOFU_VERSION=latest

FROM hashicorp/terraform:$TERRAFORM_VERSION AS terraform
FROM ghcr.io/opentofu/opentofu:$OPENTOFU_VERSION AS opentofu

FROM golang:1.21-alpine3.18
RUN apk --no-cache add make git bash

# A workaround for a permission issue of git.
# Since UIDs are different between host and container,
# the .git directory is untrusted by default.
# We need to allow it explicitly.
# https://github.com/actions/checkout/issues/760
RUN git config --global --add safe.directory /work

COPY --from=terraform /bin/terraform /usr/local/bin/
COPY --from=opentofu /usr/local/bin/tofu /usr/local/bin/
WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make install

ENTRYPOINT ["./entrypoint.sh"]
