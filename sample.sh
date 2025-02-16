#!/bin/bash

bazel run //go:main -- \
	--taste=$(pwd)/sample.taste \
	--recipe=$(pwd)/sample.recipe
