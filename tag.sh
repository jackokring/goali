#!/usr/bin/bash
# push git tags with version numbers
test $# == 1 && git tag "v$1" || echo "api.feature.bug" && git tag
git push --tags
