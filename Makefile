GOPATH = $(shell pwd)

SOL_WRAP = ../sol_wrap/osx/Release/libsol_wrap.dylib

all: direct_example persistent_example persistent_ack_example  persistent_streaming_example cache_example

lib: $(SOL_WRAP)
	( cd ../sol_wrap; make -f Makefile.osx lib )

gosol: lib
	( source ./env.sh; go install gosol )

direct_example: gosol
	( source ./env.sh; go install direct_example )

persistent_example: gosol
	( source ./env.sh; go install persistent_example )

persistent_example: gosol
	( source ./env.sh; go install persistent_example )

persistent_ack_example: gosol
	( source ./env.sh; go install persistent_ack_example )

persistent_streaming_example: gosol
	( source ./env.sh; go install persistent_streaming_example )

cache_example: gosol
	( source ./env.sh; go install cache_example )

test: direct_example persistent_example persistent_ack_example persistent_streaming_example cache_example
	./run_tests.sh
	

clean:
	$(RM) -rf bin pkg
