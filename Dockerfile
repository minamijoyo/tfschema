ARG TERRAFORM_VERSION=latest
FROM hashicorp/terraform:$TERRAFORM_VERSION AS terraform

FROM golang:1.17.11-alpine3.16
RUN apk --no-cache add make git bash

# A workaround for a permission issue of git.
# Since UIDs are different between host and container,
# the .git directory is untrusted by default.
# We need to allow it explicitly.
# https://github.com/actions/checkout/issues/760
RUN git config --global --add safe.directory /work

COPY --from=terraform /bin/terraform /usr/local/bin/
WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make install

ENTRYPOINT ["./entrypoint.sh"]
