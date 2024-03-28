#!/usr/bin/bash
#sudo apt install git python-is-python3 golang python3-pip python3-dev postgressql
git pull
git submodule updsate --init --recursive
python -m venv .
pip install -r requirements.txt
source gob.sh
