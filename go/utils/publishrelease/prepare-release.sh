#!/bin/bash
set -e
set -o pipefail

script_dir=$(dirname "$0")
cd $script_dir

[ ! -z "$1" ] || (echo "Please supply version string as first parameter"; exit 1)
version_str=$1

dolt_cmd_entry_point=../../cmd/dolt/dolt.go
release_branch=release

echo "Checking out release branch"
git fetch origin $release_branch
git checkout -b $release_branch --track "origin/$release_branch"
git merge -m "Merging latest master into $release_branch for release $version_str" master

echo "Updating the version to $version"
echo "script dir is $script_dir"
sed -i '' -e 's/Version = ".*"/Version = "'"$version_str"'"/' $dolt_cmd_entry_point

echo "Creating commit for $release_branch"
git add $dolt_cmd_entry_point
git commit -m "Updated version for release of version $version_str"

ehco "Creating tag for version $version_str"
git tag -a $version_str -m "Tag for release of Dolt version $version_str"

echo "Pushing $release_branch"
git push origin $release_branch

echo "Pushing tag for $version_str"
git push origin $version_str
