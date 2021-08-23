#!/bin/bash

if [ $# -eq 2 ]
then
    version=$1
    token=$2

    export GITHUB_TOKEN=$2

    git add .
    git commit -m "$version"
    git tag $version
    git push origin main

    goreleaser release

    rm -rf dist/
else
    echo arg1 is the version tag, arg2 is the github token
fi
