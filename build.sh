#!/bin/bash

printf "** Building linux/amd64\n"
go build -a github.com/sedarasecurity/ossecsvc
go install github.com/sedarasecurity/ossecsvc

printf "** Building linux/386\n"
go-linux-386 build -a github.com/sedarasecurity/ossecsvc
go-linux-386 install github.com/sedarasecurity/ossecsvc

printf "** Building windows/386\n"
go-windows-386 build -a github.com/sedarasecurity/ossecsvc
go-windows-386 install github.com/sedarasecurity/ossecsvc

printf "** Building windows/amd64\n"
go-windows-amd64 build -a github.com/sedarasecurity/ossecsvc
go-windows-amd64 install github.com/sedarasecurity/ossecsvc
