#########################################################################
# File Name: build-linux.sh
# Description: 編譯 linux 用的執行檔
# Author: holmes.lin
# mail: holmes.lin@gmail.com
# Created Time: 2020-11-19 13:33:39
#########################################################################
#!/bin/bash

rm -f ./iotAPI &&  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iotAPI -ldflags "-w -s" .
