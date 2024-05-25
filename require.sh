#!/usr/bin/bash
#sudo apt install git python-is-python3 golang python3-pip python3-dev postgressql
# possibly use godeb for more updated go install
#go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
#and set up a new database
git pull
git submodule update --init --recursive
python -m venv .
pip install -r requirements.txt
source gob.sh
