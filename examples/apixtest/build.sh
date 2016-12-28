#!/bin/bash

ROOT=$GOPATH/src/github.com/dshills/apix/examples/apixtest
APPNAME=apixtest
ENV=$ROOT/local.env
FPATH=$(dirname "$1")
FNAME=$(basename "$1")
EXT=${FNAME##*.}

function stop {
	SAPP=$1
	numproc=`ps aux | grep ${SAPP} | grep -v grep | wc -l`
	if [[ numproc > 0 ]]
	then
		killall ${SAPP} 2> /dev/null
	fi
}

function start {
	SENV=$1;SAPP=$2
	set -a;. ${SENV};set +a;${SAPP} &
}

function build {
	BASEDIR=$1
	cd ${BASEDIR}
	go test && go vet && go install
}

#----------------

# Take action depending on the file type
# Doing seperate builds if not in the main folder
# allows code being worked on to be compiled even if
# it is not used by the main application and as such would not be
# compiled by simply doing go install in the root directory.
# go install ./... works but test and vet get run on the vendor directory
if [ "${EXT}" == "go" ]; then
	stop ${APPNAME}
	if [ "${FPATH}" == "${ROOT}" ]; then
		build ${FPATH} && start ${ENV} ${APPNAME}
	else
		build ${FPATH} && build ${ROOT} && start ${ENV} ${APPNAME}
	fi
fi

# watch for file changes recursivly quit after first one is found and rerun the script
fswatch -1 -r ${ROOT} | xargs -n1 -I{} ${ROOT}/build.sh {}
# fswatch
# -0 output nul seperator
# -1 run just once
# -r run recursive
# xargs
# -0 expect null sperator
# -n 1 grab only the first arg
# -I execute for each input line
# {} replaced with args from fswatch
