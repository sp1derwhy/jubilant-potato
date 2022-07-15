#! /bin/bash

sourceFilePath=""
targetFilePath=""

currentDir=$(dirname "$0")

cd $currentDir && go run ../cmd -source $sourceFilePath -target $targetFilePath

