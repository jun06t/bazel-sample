#!/bin/bash

if [ "${IMAGE_TAG}" = "latest" ]; then
    BUILD_VERSION_PREFIX=`echo ${BUILDABLE}`
    echo GIT_COMMIT $(git describe --tags --match ${BUILD_VERSION_PREFIX}/*)
else
    echo GIT_COMMIT $(git describe --tags)
fi
echo DATE $(date '+%Y/%m/%d %H:%M:%S %Z')
