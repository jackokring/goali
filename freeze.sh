#!/usr/bin/bash
. xdg.sh
pip freeze | tee requirements.txt
#rm -rf "$(find extras-backup -type d -name .git)"
# add commit push
gacp() {
    date=$(date +"%A %Y-%m-%d %H:%M:%S")
    message="${1:-$date}"
    git add .
    git commit -m "$message"
    git push
}
gacp "$1"
