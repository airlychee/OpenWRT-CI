#!/bin/bash
# SPDX-License-Identifier: MIT

set -e

PKG_PATH="$GITHUB_WORKSPACE/wrt/package"
DDNS_GO_DIR="$PKG_PATH/luci-app-ddns-go/ddns-go"
DDNS_GO_PATCH="$GITHUB_WORKSPACE/Patches/ddns-go/950-silence-unchanged-ip.patch"

if [ ! -f "$DDNS_GO_DIR/Makefile" ] || [ ! -s "$DDNS_GO_PATCH" ]; then
	echo "ddns-go package or unchanged-IP patch missing; check Packages.sh and Patches/ddns-go!" >&2
	exit 1
fi

mkdir -p "$DDNS_GO_DIR/patches"
install -m 0644 \
	"$DDNS_GO_PATCH" \
	"$DDNS_GO_DIR/patches/950-silence-unchanged-ip.patch"

echo "ddns-go unchanged-IP patch installed; other logs remain enabled."
