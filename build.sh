#!/bin/bash
set -e
cd webui-vue
yarn run build
cd -
docker build . -t gotplonline