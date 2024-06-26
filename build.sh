#!/usr/bin/env bash

GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=jsoniter,poll_opt -gcflags "all=-N -l" -ldflags="-X 'auxiliary/cmd/auxiliary/internal.DefaultConfigPath=/etc/auxiliary/config.yaml'" -o build/release/auxiliary cmd/auxiliary/main.go
scp build/release/auxiliary root@10.10.10.221:/usr/local/bin/auxiliary
