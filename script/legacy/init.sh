#!/bin/bash
pushd `dirname $0` > /dev/null
SCRIPTPATH=`pwd -P`
goop install
cd "$SCRIPTPATH/migrate"
goop install
goop exec go run migrate.go
cd "$SCRIPTPATH/frontend/compiler"
npm install
gulp
popd > /dev/null
