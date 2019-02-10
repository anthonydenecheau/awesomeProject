echo "Building with travis commit of $BUILD_NAME ..."
docker build -t anthonydenecheau/gopocservice:$BUILD_NAME .
