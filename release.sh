#!/bin/bash

version=$1

git add .
git commit -m "$version"
git tag $version
git push origin main

goreleaser release

rm -rf dist/