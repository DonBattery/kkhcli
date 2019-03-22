#!/bin/bash

# You need to edit this file manually to change target architecture or OS
# List of available options
# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

Architecture="amd64"
OS="linux"


echo "Building KKHCLI..."

docker run --rm -e GOARCH=$Architecture -e GOOS=$OS -e GOPATH=/go -v `pwd`:/go/src/github.com/donbattery/kkhcli golang /bin/bash -c "cd /go/src/github.com/donbattery/kkhcli; go build ."

if [ $? -eq 0 ]; then
    echo "Built kkhcli successfully..."
    exit 0
else
    echo "Build failed..."
    exit 1
fi
