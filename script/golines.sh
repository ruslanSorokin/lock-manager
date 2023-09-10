#!/bin/bash

SRC_PATH=${SRC_PATH:-.}

golines -w --reformat-tags --tab-len=2 --base-formatter=gofumpt --max-len=80 $SRC_PATH
