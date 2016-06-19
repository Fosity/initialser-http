#!/usr/bin/env bash
gox -ldflags "-s -w" -osarch="windows/386" -output dist/initialser-win32
gox -ldflags "-s -w" -osarch="windows/amd64" -output dist/initialser
gox -ldflags "-s -w" -os="linux"   -output dist/initialser-linux
gox -ldflags "-s -w" -os="darwin"  -output dist/initialser-mac
