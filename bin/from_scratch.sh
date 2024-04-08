#!/bin/bash

# shellcheck disable=SC2034
SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

echo "SERVICE: $1";
echo "DIRECTORY: $SCRIPT_DIR/$1"

if [ -z "$1" ]
then
    echo "name of service is not specified, use ./bin/from_scratch example"
    exit 1;
fi


mkdir "$SCRIPT_DIR/$1"
cp -r "$SCRIPT_DIR/some/." "$SCRIPT_DIR/$1"

sed -i -e "s/{{SERVICE}}/$1/g" "$SCRIPT_DIR/$1/build.sh"
sed -i -e "s/{{SERVICE}}/$1/g" "$SCRIPT_DIR/$1/deploy.sh"
sed -i -e "s/{{SERVICE}}/$1/g" "$SCRIPT_DIR/$1/run.sh"
sed -i -e "s/{{SERVICE}}/$1/g" "$SCRIPT_DIR/$1/some.service"

mv "$SCRIPT_DIR/$1/some.service" "$SCRIPT_DIR/$1/$1.service"

echo "generated $1 service"

# TODO: add generate golang ./cmd/some/main.go and ./cmd/some/Dockerfile