#!/bin/sh

nohup ./cow -debug &
export GOPACKAGESDEBUG=1 GO111MODULE=on GOPACKAGESPRINTGOLISTERRORS=1
nohup /home/yaoguai/workspace/go/bin/gocode-gomod -s -sock unix -addr 127.0.0.1:37373 -debug &
