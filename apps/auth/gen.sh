#!/bin/sh

#
# Examples:
# webrpc-gen -schema=./rpc/product.ridl -target=go -pkg=product -server -out=./services/product/webrpc.gen.go
# webrpc-gen -schema=./rpc/product.ridl -target=ts -client -out=./sdk/product.rpc.gen.ts

command -v webrpc-gen >/dev/null 2>&1 || {
	echo >&2 "webrpc-gen is not installed";
	echo >&2 "go get -u github.com/webrpc/webrpc/cmd/webrpc-gen";
	exit 1;
}

mkdir -p ./sdk/

for FILE in ./rpc/*.ridl; do
	NAME=$(basename $FILE .ridl)
	DIR=./services/${NAME}

	if [ ! -d $DIR ]
	then
		echo "$DIR not present, make sure it exists first."
	fi

	echo "webrpc-gen -schema=${FILE} -target=js -client -out=./sdk/${NAME}.rpc.gen.js"
	webrpc-gen -schema=${FILE} -target=js -client -out=./sdk/${NAME}.rpc.gen.js

	echo "webrpc-gen -schema=${FILE} -target=go -pkg=${NAME} -server -out=./services/${NAME}/webrpc.gen.go"
	webrpc-gen -schema=${FILE} -target=go -pkg=${NAME} -server -out=./services/${NAME}/webrpc.gen.go
done
