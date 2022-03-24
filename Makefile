.PHONY: default testing local deps prepare man gofmt lfparser_disable_line plparser_disable_line

default: local

# 正式环境
PRODUCTION_DOWNLOAD_ADDR = zhuyun-static-files-production.oss-cn-hangzhou.aliyuncs.com/datakit

# 测试环境
TESTING_DOWNLOAD_ADDR = zhuyun-static-files-testing.oss-cn-hangzhou.aliyuncs.com/datakit

# 本地环境: 需配置环境变量，便于完整测试采集器的发布、更新等流程
# export LOCAL_OSS_ACCESS_KEY='<your-oss-AK>'
# export LOCAL_OSS_SECRET_KEY='<your-oss-SK>'
# export LOCAL_OSS_BUCKET='<your-oss-bucket>'
# export LOCAL_OSS_HOST='oss-cn-hangzhou.aliyuncs.com' # 一般都是这个地址
# export LOCAL_OSS_ADDR='<your-oss-bucket>.oss-cn-hangzhou.aliyuncs.com/datakit'
# 如果只是编译，LOCAL_OSS_ADDR 这个环境变量可以随便给个值
LOCAL_DOWNLOAD_ADDR=${LOCAL_OSS_ADDR}


PUB_DIR = dist
BUILD_DIR = dist

BIN = datakit
NAME = datakit
ENTRY = cmd/datakit/main.go

LOCAL_ARCHS:="local"
DEFAULT_ARCHS:="all"
MAC_ARCHS:="darwin/amd64"
NOT_SET="not-set"
VERSION?=$(shell git describe --always --tags)
DATE:=$(shell date -u +'%Y-%m-%d %H:%M:%S')
GOVERSION:=$(shell go version)
COMMIT:=$(shell git rev-parse --short HEAD)
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
COMMITER:=$(shell git log -1 --pretty=format:'%an')
UPLOADER:=$(shell hostname)/${USER}/${COMMITER}
DOCKER_IMAGE_ARCHS:="linux/arm64,linux/amd64"

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
GO_PATCH_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f3)
MINIMUM_SUPPORTED_GO_MAJOR_VERSION = 1
MINIMUM_SUPPORTED_GO_MINOR_VERSION = 16
GO_VERSION_VALIDATION_ERR_MSG = Golang version is not supported, please update to at least $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION).$(MINIMUM_SUPPORTED_GO_MINOR_VERSION)
BUILDER_GOOS_GOARCH=$(shell go env GOOS)-$(shell go env GOARCH)

GOLINT_VERSION = "$(shell golangci-lint --version | cut -c 27- | cut -d' ' -f1)"
SUPPORTED_GOLINT_VERSION = "1.42.1"
SUPPORTED_GOLINT_VERSION_ANOTHER = "v1.42.1"
GOLINT_VERSION_VALIDATION_ERR_MSG = golangci-lint version($(GOLINT_VERSION)) is not supported, please use version $(SUPPORTED_GOLINT_VERSION)

#####################
# Large strings
#####################

define GIT_INFO
//nolint
package git

const (
	BuildAt  string = "$(DATE)"
	Version  string = "$(VERSION)"
	Golang   string = "$(GOVERSION)"
	Commit   string = "$(COMMIT)"
	Branch   string = "$(GIT_BRANCH)"
	Uploader string = "$(UPLOADER)"
);
endef
export GIT_INFO

define build
	@if [ $(GO_MAJOR_VERSION) -gt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		exit 0 ; \
	elif [ $(GO_MAJOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MINOR_VERSION) ] ; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	fi

	@rm -rf $(PUB_DIR)/$(1)/*
	@mkdir -p $(BUILD_DIR) $(PUB_DIR)/$(1)
	@echo "===== $(BIN) $(1) ===="
	@GO111MODULE=off CGO_ENABLED=0 go run cmd/make/make.go \
		-main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir $(BUILD_DIR) \
		-release $(1) -pub-dir $(PUB_DIR) -archs $(2) -download-addr $(3)
	@tree -Csh -L 3 $(BUILD_DIR)
endef

define pub
	@echo "publish $(1) $(NAME) ..."
	@GO111MODULE=off go run cmd/make/make.go \
		-pub -release $(1) -pub-dir $(PUB_DIR) \
		-name $(NAME) -download-addr $(2) \
		-build-dir $(BUILD_DIR) -archs $(3)
endef

define build_docker_image
	@if [ $(2) = "registry.jiagouyun.com" ]; then \
		echo 'publish to $(2)...'; \
		sudo docker buildx build --platform $(1) \
			-t $(2)/datakit/datakit:$(VERSION) . --push ; \
		sudo docker buildx build --platform $(1) \
			-t $(2)/datakit/logfwd:$(VERSION) -f Dockerfile_logfwd . --push ; \
	else \
		echo 'publish to $(2)...'; \
		sudo docker buildx build --platform $(1) \
			-t $(2)/datakit/datakit:$(VERSION) \
			-t $(2)/dataflux/datakit:$(VERSION) \
			-t $(2)/dataflux-prev/datakit:$(VERSION) . --push; \
		sudo docker buildx build --platform $(1) \
			-t $(2)/datakit/logfwd:$(VERSION) \
			-t $(2)/dataflux/logfwd:$(VERSION) \
			-t $(2)/dataflux-prev/logfwd:$(VERSION) -f Dockerfile_logfwd . --push; \
	fi
endef

define check_golint_version
	@case $(GOLINT_VERSION) in \
	$(SUPPORTED_GOLINT_VERSION)) \
	;; \
	$(SUPPORTED_GOLINT_VERSION_ANOTHER)) \
	;; \
	*) \
		echo '$(GOLINT_VERSION_VALIDATION_ERR_MSG)'; \
		exit 1; \
	esac;
endef

local: deps
	$(call build,local, $(LOCAL_ARCHS), $(LOCAL_DOWNLOAD_ADDR))

pub_local:
	$(call pub, local,$(LOCAL_DOWNLOAD_ADDR),$(LOCAL_ARCHS))

testing: deps
	$(call build, testing, $(DEFAULT_ARCHS), $(TESTING_DOWNLOAD_ADDR))
	$(call pub, testing,$(TESTING_DOWNLOAD_ADDR),$(DEFAULT_ARCHS))

testing_image:
	$(call build_docker_image, $(DOCKER_IMAGE_ARCHS), 'registry.jiagouyun.com')
	# we also publish testing image to public image repo
	$(call build_docker_image, $(DOCKER_IMAGE_ARCHS), 'pubrepo.jiagouyun.com')

production: deps # stable release
	$(call build, production, $(DEFAULT_ARCHS), $(PRODUCTION_DOWNLOAD_ADDR))
	$(call pub, production, $(PRODUCTION_DOWNLOAD_ADDR),$(DEFAULT_ARCHS))

production_image:
	$(call build_docker_image, $(DOCKER_IMAGE_ARCHS), 'pubrepo.jiagouyun.com')

production_mac: deps
	$(call build, production, $(MAC_ARCHS), $(PRODUCTION_DOWNLOAD_ADDR))
	$(call pub,production,$(PRODUCTION_DOWNLOAD_ADDR),$(MAC_ARCHS))

testing_mac: deps
	$(call build, testing, $(MAC_ARCHS), $(TESTING_DOWNLOAD_ADDR))
	$(call pub, testing,$(TESTING_DOWNLOAD_ADDR),$(MAC_ARCHS))

# not used
pub_testing_win_img:
	@mkdir -p embed/windows-amd64
	@wget --quiet -O - "https://$(TESTING_DOWNLOAD_ADDR)/iploc/iploc.tar.gz" | tar -xz -C .
	@sudo docker build -t registry.jiagouyun.com/datakit/datakit-win:$(VERSION) -f ./Dockerfile_win .
	@sudo docker push registry.jiagouyun.com/datakit/datakit-win:$(VERSION)

# not used
pub_release_win_img:
	# release to pub hub
	@mkdir -p embed/windows-amd64
	@wget --quiet -O - "https://$(PRODUCTION_DOWNLOAD_ADDR)/iploc/iploc.tar.gz" | tar -xz -C .
	@sudo docker build -t pubrepo.jiagouyun.com/datakit/datakit-win:$(VERSION) -f ./Dockerfile_win .
	@sudo docker push pubrepo.jiagouyun.com/datakit/datakit-win:$(VERSION)

# Config samples should only be published by production release,
# because config samples in multiple testing releases may not be compatible to each other.
pub_conf_samples:
	@echo "upload config samples to oss..."
	@go run cmd/make/make.go -dump-samples -release production

# testing/production downloads config samples from different oss bucket.
check_testing_conf_compatible:
	@go run cmd/make/make.go -download-samples -release testing
	@LOGGER_PATH=nul ./dist/datakit-$(BUILDER_GOOS_GOARCH)/datakit --check-config --config-dir samples
	@LOGGER_PATH=nul ./dist/datakit-$(BUILDER_GOOS_GOARCH)/datakit --check-sample

check_production_conf_compatible:
	@go run cmd/make/make.go -download-samples -release production
	@LOGGER_PATH=nul ./dist/datakit-$(BUILDER_GOOS_GOARCH)/datakit --check-config --config-dir samples
	@LOGGER_PATH=nul ./dist/datakit-$(BUILDER_GOOS_GOARCH)/datakit --check-sample

define build_ip2isp
	rm -rf china-operator-ip
	git clone -b ip-lists https://github.com/gaoyifan/china-operator-ip.git
	@GO111MODULE=off CGO_ENABLED=0 go run cmd/make/make.go -build-isp
endef

define do_lint
	truncate -s 0 lint.err
	golangci-lint --version 
	GOARCH=$(1) GOOS=$(2) golangci-lint run --fix --allow-parallel-runners
endef

ip2isp:
	$(call build_ip2isp)

deps: prepare man gofmt lfparser_disable_line plparser_disable_line

man:
	@packr2 clean
	@packr2

# ignore files under vendor/.git/git
# install gofumpt: go install mvdan.cc/gofumpt@latest
gofmt:
	@GO111MODULE=off gofumpt -w -l $(shell find . -type f -name '*.go'| grep -v "/vendor/\|/.git/\|/git/\|.*_y.go")

vet:
	@go vet ./...

ut: deps
	@GO111MODULE=off CGO_ENABLED=1 go run cmd/make/make.go -ut

# all testing

all_test: deps
	@truncate -s 0 test.output
	@echo "#####################" | tee -a test.output
	@echo "#" $(DATE) | tee -a test.output
	@echo "#" $(VERSION) | tee -a test.output
	@echo "#####################" | tee -a test.output
	i=0; \
	for pkg in `go list ./... | grep -vE 'datakit/git'`; do \
		echo "# testing $$pkg..." | tee -a test.output; \
		GO111MODULE=off CGO_ENABLED=1 LOGGER_PATH=nul go test -timeout 1m -cover $$pkg; \
		if [ $$? != 0 ]; then \
			printf "\033[31m [FAIL] %s\n\033[0m" $$pkg; \
			i=`expr $$i + 1`; \
		else \
			echo "######################"; \
			fi \
	done; \
	if [ $$i -gt 0 ]; then \
		printf "\033[31m %d case failed.\n\033[0m" $$i; \
		exit 1; \
	else \
		printf "\033[32m all testinig passed.\n\033[0m"; \
	fi

test_deps: prepare man gofmt lfparser_disable_line plparser_disable_line vet

lint: deps
	$(call do_lint,386,windows)
	$(call do_lint,amd64,windows)
	$(call do_lint,amd64,linux)
	$(call do_lint,386,linux)
	$(call do_lint,arm,linux)
	$(call do_lint,arm64,linux)
	$(call do_lint,amd64,darwin)

lfparser_disable_line:
	@rm -rf io/parser/gram_y.go
	@rm -rf io/parser/gram.y.go
	@rm -rf io/parser/parser.y.go
	@rm -rf io/parser/parser_y.go
	@goyacc -l -o io/parser/gram_y.go io/parser/gram.y # use -l to disable `//line`

plparser_disable_line:
	@rm -rf pipeline/parser/gram_y.go
	@rm -rf pipeline/parser/gram.y.go
	@rm -rf pipeline/parser/parser.y.go
	@rm -rf pipeline/parser/parser_y.go
	@goyacc -l -o pipeline/parser/gram_y.go pipeline/parser/gram.y # use -l to disable `//line`

prepare:
	@mkdir -p git
	@echo "$$GIT_INFO" > git/git.go

check_man:
	grep --color=always -nrP "[a-zA-Z0-9][\p{Han}]|[\p{Han}][a-zA-Z0-9]" man > bad-doc.log
	if [ $$? != 0 ]; then \
		echo "check manuals ok"; \
	else \
		cat bad-doc.log; \
		rm -rf bad-doc.log; \
	fi

clean:
	@rm -rf build/*
	@rm -rf io/parser/gram_y.go
	@rm -rf io/parser/gram.y.go
	@rm -rf pipeline/parser/parser.y.go
	@rm -rf pipeline/parser/parser_y.go
	@rm -rf pipeline/parser/gram.y.go
	@rm -rf pipeline/parser/gram_y.go
	@rm -rf check.err
	@rm -rf $(PUB_DIR)/*
