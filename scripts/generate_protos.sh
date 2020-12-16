#!/bin/bash

DIR=$(pwd)
rm -rf $DIR/gen
LYFT_IMAGE="lyft/protocgenerator:8167e11d3b3439373c2f033080a4b550078884a2"

docker run --rm -u $(id -u):$(id -g) -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos -l go --go_source_relative --validate_out

languages=("python" "cpp" "java")
for lang in "${languages[@]}"
do
    docker run --rm -u $(id -u):$(id -g) -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos -l $lang
done

# Docs generated
docker run --rm -u $(id -u):$(id -g) -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos -l protodoc
