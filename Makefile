BUILD_TAG ?= suapapa/khaiii-api:latest
# BUILD_ARCHS ?= linux/amd64,linux/arm64
BUILD_ARCHS ?= linux/arm64
DOCKERFILE ?= Dockerfile
CONTEXT ?= .

build_image:
	docker buildx build --platform $(BUILD_ARCHS) -t $(BUILD_TAG) -f $(DOCKERFILE) $(CONTEXT) 

push_image:
	docker push $(BUILD_TAG)

.PHONY: build_image push_image