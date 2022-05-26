#!/bin/bash
git pull
rm -rf hego
go build -o hego cmd/search/main.go
chmod 777 hego
./hego