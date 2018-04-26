#!/bin/sh

set -e

if [ -e /firstrun ]; then
   echo "Not first run so skipping initialization"
else
    dep ensure
fi

tail -f /dev/null