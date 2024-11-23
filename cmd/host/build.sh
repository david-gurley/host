#!/usr/bin/env bash

BINARY=/mnt/share/software/bin/slickhost
go build -o ${BINARY} 
chmod 755 ${BINARY}
${BINARY} --help
