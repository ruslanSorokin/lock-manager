#!/bin/bash

GO_SRC_PATH=${GO_SRC_PATH:-.}
COMPANY_PREFIXES=${COMPANY_PREFIXES:-}

if [[ -n $COMPANY_PREFIXES ]]; then
  _COMPANY_PREFIXES_OPT="-company-prefixes $COMPANY_PREFIXES"
else
  _COMPANY_PREFIXES_OPT=""
fi

goimports-reviser "$GO_SRC_PATH/..." $_COMPANY_PREFIXES_OPT -format -set-alias -rm-unused -use-cache
