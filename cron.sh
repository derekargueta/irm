#!/bin/bash

if [ ! -f "/code/irm/tmp/irm-data/results.csv" ]; then
  mkdir -p /code/irm/tmp
  git clone git@github.com:derekargueta/irm-data.git > /code/irm/tmp/irm-data
fi

cd /code/irm/tmp/irm-data
git pull origin master
cd ../..
make build-docker
make cron-docker
cd tmp/irm-data
git push origin master

exit 0
