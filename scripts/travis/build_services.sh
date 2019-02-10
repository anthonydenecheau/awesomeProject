echo "Building with travis commit of $BUILD_NAME ..."
docker build -t awesomeproject:$BUILD_NAME .
