#!/bin/bash

docker run -e GOOS=darwin -e GOPATH=/go/gege -v `pwd`:/go/src/github.com/donbattery/kkhcli golang go build go/src/github.com/donbattery/kkhcli/.