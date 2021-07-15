#!/bin/sh

cd cmd/analyze/results
cp results.csv /root/irm-data/
cd ~/irm-data/
git add --all
git commit -m "timely commit"
git push





# set -euxo pipefail

# data_directory="$HOME/code/irm/tmp/irm-data"

# if [ ! -d "$data_directory" ]; then
#   mkdir -p "$HOME/code/irm/tmp"
#   cd "$HOME/code/irm/tmp"
#   git clone git@github.com:derekargueta/irm-data.git
# fi

# cd "$data_directory"
# git pull origin master
# cd ../..
# make build-docker
# make cron-docker
# cd tmp/irm-data
# git push origin master

# exit 0
