#!/bin/bash
pushd `dirname $0` > /dev/null
SCRIPTPATH=`pwd -P`
cd "$SCRIPTPATH/frontend/canjs/compiler"
gulp
gulp watch
popd > /dev/null
