#!/bin/bash

echo GIT_COMMIT $(git rev-parse --verify HEAD)
echo DATE $(date '+%Y/%m/%d %H:%M:%S %Z')
