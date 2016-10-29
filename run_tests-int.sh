#!/bin/bash

. ./env.sh

./bin/direct_example   ../properties/demotr-int.properties
./bin/persistent_example ../properties/demotr-int.properties
./bin/persistent_ack_example ../properties/demotr-int.properties
./bin/persistent_streaming_example ../properties/demotr-int.properties
# ./bin/cache_example   ../properties/demotr-int.properties

