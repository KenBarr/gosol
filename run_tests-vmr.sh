#!/bin/bash

. ./env.sh

./bin/direct_example   ../properties/vmr.properties
./bin/persistent_example ../properties/vmr.properties
./bin/persistent_ack_example ../properties/vmr.properties
./bin/persistent_streaming_example ../properties/vmr.properties
# ./bin/cache_example   ../properties/vmr.properties

