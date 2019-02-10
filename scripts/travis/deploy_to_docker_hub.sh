echo "Pushing service docker images to docker hub ...."
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push anthonydenecheau/gopocservice:$BUILD_NAME
