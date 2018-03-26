#! /bin/sh

GOOS=darwin go build -ldflags="-s -w" -o $2.temp "$1"
upx -f --brute -o $2 $2.temp
rm -rf $2.temp

ls -l $2*