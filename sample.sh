#!/bin/bash

bazel run //go:main -- \
	--taste_file_path=$(pwd)/sample.taste \
	--recipe_file_path=$(pwd)/sample.recipe \
	--output_file_path=$(pwd)/out.tmp \
	--override_output
