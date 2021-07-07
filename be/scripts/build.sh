SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPTDIR}" )" &> /dev/null && pwd )"
CMD_DIR=${ROOT_DIR}/cmd
BIN_DIR=${ROOT_DIR}/bin

mkdir -p ${BIN_DIR}
cd ${BIN_DIR}
go build ${CMD_DIR}/gateserver/gate_server.go
go build ${CMD_DIR}/userserver/user_server.go

