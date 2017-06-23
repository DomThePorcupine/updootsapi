#! /bin/bash
# Set the gopath to our current working directory
export GOPATH=$(pwd)
# For convenience, add the workspace's bin subdirectory to your PATH
export PATH=$PATH:$(go env GOPATH)/bin