# This file is based on code originally licensed unter the Apache License, version 2.0.
# It has been modified by Cloudogu GmbH and is distributed under the AGPL-3.0-only as part of the velero-plugin-for-restore-exclude Project.
# Original code Copyright 2017, 2019, 2020 the Velero contributors.
# Modification Copyright 2025 - present, Cloudogu GmbH

FROM golang:1.24-bookworm AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /go/src/github.com/cloudogu/velero-plugin-for-restore-exclude
COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor/ vendor/

COPY main.go main.go
COPY internal/ internal/

RUN CGO_ENABLED=0 go build -mod=vendor -o /go/bin/velero-plugin-for-restore-exclude .

FROM busybox:1.33.1 AS busybox

FROM scratch
COPY --from=build /go/bin/velero-plugin-for-restore-exclude /plugins/
COPY --from=busybox /bin/cp /bin/cp
USER 65532:65532
ENTRYPOINT ["cp", "/plugins/velero-plugin-for-restore-exclude", "/target/."]
