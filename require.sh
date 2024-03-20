#!/usr/bin/bash
#sudo apt install git python-is-python3 golang python3-pip python3-dev
git pull
python -m venv .
pip install -r requirements.txt
source gob.sh
