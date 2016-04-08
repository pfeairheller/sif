#!/bin/sh

PWD=`pwd`
export GOPATH=$PWD:$PWD/vendor
export PATH=$PWD/bin:$PWD/vendor/bin:$PATH
