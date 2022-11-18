#!/bin/bash

VERSION="$(date '+%Y%m%d%H%M%S')-dev" 
VERSION=$VERSION make
./build/i3-companion --version

if [[ "$1" == "package" ]]; then
  VERSION=$VERSION make package
fi
