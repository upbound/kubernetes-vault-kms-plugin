# Project Setup
PROJECT_NAME := kubernetes-vault-kms-plugin
PROJECT_REPO := github.com/turkenh/$(PROJECT_NAME)

PLATFORMS ?= linux_amd64 linux_arm64
include build/makelib/common.mk

# Setup Go
GO_LDFLAGS += -X $(GO_PROJECT)/pkg/version.Version=$(VERSION)
GO_STATIC_PACKAGES = $(GO_PROJECT)/server
GO111MODULE = on
include build/makelib/golang.mk

# Docker images
DOCKER_REGISTRY = turkenh
IMAGES = kubernetes-vault-kms-plugin
include build/makelib/image.mk