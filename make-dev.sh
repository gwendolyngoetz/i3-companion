#!/bin/bash

VERSION="-$(date '+%Y%m%d%H%M%S')-dev" make
./build/i3-companion --version
