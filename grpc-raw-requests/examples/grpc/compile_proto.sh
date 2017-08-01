#!/bin/bash

DIR="$( pwd )"
PROTOFILE=$(ls $DIR | egrep '.proto$')

protoc --proto_path=$DIR/ --go_out=plugins=grpc:$DIR "$DIR/$PROTOFILE"