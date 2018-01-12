#!/bin/bash
set -eo pipefail

cd $(git rev-parse --show-toplevel)

finalize() {
    sed -i '1s;^;/* tslint:disable */\n;' ./src/pb/${1}.d.ts
}

export PATH=$PATH:$(pwd)/node_modules/.bin
pbjs -t static-module -w commonjs -o ./timestamp.js timestamp.proto
pbts -o ./timestamp.d.ts ./timestamp.js
