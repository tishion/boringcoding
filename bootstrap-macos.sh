#!/bin/zsh

# restart mysql service
# make sure mysql service is running on local host
# connetion string "test:test@tcp(localhost:3306)/test"
brew services restart mysql

# restart redis service
# make sure redis service is running on local host with default port
brew services restart redis

# get the absolute path of current script folder
ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# build server applications 
mkdir -p ${ROOT_DIR}/be/bin
cd ${ROOT_DIR}/be/bin
go build ${ROOT_DIR}/be/cmd/gateserver/gate_server.go
go build ${ROOT_DIR}/be/cmd/userserver/user_server.go

# run business server
cd ${ROOT_DIR}
nohup ${ROOT_DIR}/be/bin/user_server > user_server.log 2>&1 &

# run gate server
cd ${ROOT_DIR}
nohup ${ROOT_DIR}/be/bin/gate_server > gate_server.log 2>&1 &

# run the fe project
cd ${ROOT_DIR}/fe/assbook
npm i
npm run serve