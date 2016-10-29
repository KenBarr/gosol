#!/bin/bash

. ./env.sh

./bin/direct_example   ../properties/demotr-ext.properties
./bin/persistent_example ../properties/demotr-ext.properties
./bin/persistent_ack_example ../properties/demotr-ext.properties
./bin/persistent_streaming_example ../properties/demotr-ext.properties
# ./bin/cache_example   ../properties/demotr-ext.properties

