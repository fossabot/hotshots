# Hotshots 🔥
[![Build Status](https://travis-ci.org/kochman/hotshots.svg?branch=master)](https://travis-ci.org/kochman/hotshots)&nbsp;[![codecov](https://codecov.io/gh/kochman/hotshots/branch/master/graph/badge.svg)](https://codecov.io/gh/kochman/hotshots)

Hotshots automatically uploads photos from remote cameras to a server, allowing news organization editors or event photographers to quickly curate and share content in real-time through social media.

## Development

```
go get github.com/kochman/hotshots
go get -u github.com/kardianos/govendor
cd $GOPATH/src/github.com/kochman/hotshots
govendor sync
go build -i
./hotshots -h
```
