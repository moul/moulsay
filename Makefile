GOPKG ?=	moul.io/moulsay
DOCKER_IMAGE ?=	moul/moulsay
GOBINS ?=	.

all: test install

-include rules.mk
