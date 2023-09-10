#!/bin/bash

GO_SRC_PATH=${GO_SRC_PATH:-.}
CONFIG_FILE=${CONFIG_FILE:-}

if [[ -n $CONFIG_FILE ]]; then
  _CONFIG_FILE_OPT="--config $CONFIG_FILE"
else
  _CONFIG_FILE_OPT=""
fi

cd $GO_SRC_PATH && mockery $_CONFIG_FILE_OPT
