echo "Building with travis commit of $BUILD_NAME ..."
docker build -t gopocservice:$BUILD_NAME .
