#!/usr/bin/env bash

set -ex

# Clear old logs
rm -rf test/integration/log/*

# RM and Generate Go Code
rm -rf test/integration/go/gen/*
frugal --gen go:package_prefix=github.com/Workiva/frugal/ -r --out='test/integration/go/gen' test/integration/frugalTest.frugal

# Create Go binaries
rm -rf test/integration/go/bin/*
go build -o test/integration/go/bin/testclient test/integration/go/src/bin/testclient/main.go
go build -o test/integration/go/bin/testserver test/integration/go/src/bin/testserver/main.go

# RM and Generate Dart Code
rm -rf test/integration/dart/gen-dart/
frugal --gen dart -r --out='test/integration/dart/gen-dart' test/integration/frugalTest.frugal

# Pub get hackery.  This can be fixed when skynet has pub credentials.
cd test/integration/dart/test_client
pub get
cd ../gen-dart/frugal_test
pub get
cd ../../../../..
echo $PWD

# Run tests
# -v flag for verbose output
# --server for specific server languages (only go supported currently)
# --client for specific client languages (go and dart supported currently)
# Example: python test/integration/test.py --server go --client go
python test/integration/test.py --retry-count=0
