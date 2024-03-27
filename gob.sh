#!/usr/bin/bash
#go build script a.k.a. gob.sh
# from https://belief-driven-design.com/build-time-variables-in-go-51439b26ef9/ with edits

# STEP 0: Git to have it

git add .
git commit -m "GOB.SH the IT BOT"
# no push for purpose

# STEP 1: Determinate the required values

# "/v2"
BRANCHVERSYNTAX=""

PACKAGE="github.com/$(basename $HOME)/$(basename $(pwd))$BRANCHVERSYNTAX"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S%Z')
DYNAMIC="dynamic"

# STEP 2: The dynamic ldflags 

LDFLAGS=(
  "-X '${PACKAGE}/consts.Version=${VERSION}-${COMMIT_HASH}'"
  "-X '${PACKAGE}/consts.BuildTime=${BUILD_TIMESTAMP}'"
  "-X '${PACKAGE}/consts.Dynamic=${DYNAMIC}'"
)

# STEP 3: Build shared C object

go generate
go build -ldflags="${LDFLAGS[*]}" -buildmode=c-shared
# rename .so dynamic lib
mv goali goali.so 

# STEP 4: Build the ldflags

LDFLAGS=(
  "-X '${PACKAGE}/consts.Version=${VERSION}-${COMMIT_HASH}'"
  "-X '${PACKAGE}/consts.BuildTime=${BUILD_TIMESTAMP}'"
)

# STEP 5: Actual Go build process

go generate
go build -ldflags="${LDFLAGS[*]}"
