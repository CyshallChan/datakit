.PHONY: default test

default: release

# 正式环境
RELEASE_DOWNLOAD_ADDR = zhuyun-static-files-production.oss-cn-hangzhou.aliyuncs.com/datakit

# 测试环境
TEST_DOWNLOAD_ADDR = zhuyun-static-files-testing.oss-cn-hangzhou.aliyuncs.com/datakit

PUB_DIR = pub
BIN = datakit
NAME = datakit
ENTRY = main.go

VERSION := $(shell git describe --always --tags)

all: test release preprod local


local:
	$(call build,local,$(LOCAL_KODO_HOST),$(LOCAL_DOWNLOAD_ADDR),$(LOCAL_SSL),$(LOCAL_PORT))


define build
	@echo "===== $(BIN) $(1) ===="
	@rm -rf $(PUB_DIR)/$(1)/*
	@mkdir -p build $(PUB_DIR)/$(1)
	@mkdir -p git
	@echo 'package git; const (Sha1 string=""; BuildAt string=""; Version string=""; Golang string="")' > git/git.go
	@go run make.go -main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir build  \
		 -download-addr $(2) -release $(1) -pub-dir $(PUB_DIR) -archs $(3) -os $(4)
	#@strip build/$(NAME)-linux-amd64/$(BIN)
	#@tar czf $(PUB_DIR)/release/$(NAME)-$(VERSION).tar.gz autostart agent -C build .
	tree -Csh $(PUB_DIR)
endef

test:
	$(call build,test,$(TEST_DOWNLOAD_ADDR),"linux/amd64","linux")

release:
	$(call build,release,$(RELEASE_DOWNLOAD_ADDR),"linux/amd64","linux")

	# @echo "===== $(BIN) release ===="
	# @rm -rf $(PUB_DIR)/release
	# @mkdir -p build $(PUB_DIR)/release
	# @mkdir -p git
	# @echo 'package git; const (Sha1 string=""; BuildAt string=""; Version string=""; Golang string="")' > git/git.go
	# @go run make.go -main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir build -archs "linux/amd64" \
	# 	 -download-addr $(RELEASE_DOWNLOAD_ADDR) -release release -pub-dir $(PUB_DIR)
	# #@strip build/$(NAME)-linux-amd64/$(BIN)
	# #@tar czf $(PUB_DIR)/release/$(NAME)-$(VERSION).tar.gz autostart agent -C build .
	# tree -Csh $(PUB_DIR)


define pub
	echo "publish $(1) $(NAME) ..."
	go run make.go -pub -release $(1) -pub-dir $(PUB_DIR) -name $(NAME) -download-addr $(2) -archs $(3) -os $(4)
endef


pub_test:
	$(call pub,test,$(TEST_DOWNLOAD_ADDR),"linux/amd64","linux")

pub_release:
	$(call pub,release,$(RELEASE_DOWNLOAD_ADDR),"linux/amd64","linux")

test_win:
	@echo "===== $(BIN) test_win ===="
	@rm -rf $(PUB_DIR)/test_win
	@mkdir -p build $(PUB_DIR)/test_win
	@mkdir -p git
	@echo 'package git; const (Sha1 string=""; BuildAt string=""; Version string=""; Golang string="")' > git/git.go
	@go run make.go -main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir build -archs "windows/amd64" \
		 -download-addr $(TEST_DOWNLOAD_ADDR_WIN) -release test -pub-dir $(PUB_DIR) -windows
	#@strip build/$(NAME)-linux-amd64/$(BIN)
	#@tar czf $(PUB_DIR)/test_win/$(NAME)-$(VERSION).tar.gz -C windows agent.exe -C ../build .
	tree -Csh $(PUB_DIR)

release_win:
	@echo "===== $(BIN) release_win ===="
	@rm -rf $(PUB_DIR)/release_win
	@mkdir -p build $(PUB_DIR)/release_win
	@mkdir -p git
	@echo 'package git; const (Sha1 string=""; BuildAt string=""; Version string=""; Golang string="")' > git/git.go
	@go run make.go -main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir build -archs "windows/amd64" \
		 -download-addr $(RELEASE_DOWNLOAD_ADDR) -release release -pub-dir $(PUB_DIR) -windows
	#@strip build/$(NAME)-linux-amd64/$(BIN)
	#@tar czf $(PUB_DIR)/test_win/$(NAME)-$(VERSION).tar.gz -C windows agent.exe -C ../build .
	tree -Csh $(PUB_DIR)


test_mac:
	@echo "===== $(BIN) test_mac ===="
	@rm -rf $(PUB_DIR)/test_mac
	@mkdir -p build $(PUB_DIR)/test_mac
	@go run make.go -main $(ENTRY) -binary $(BIN) -name $(NAME) -build-dir build -archs "darwin/amd64" \
		 -download-addr $(TEST_DOWNLOAD_ADDR_WIN) -release test -pub-dir $(PUB_DIR) -mac
	#@strip build/$(NAME)-linux-amd64/$(BIN)
	@tar czf $(PUB_DIR)/test_mac/$(NAME)-$(VERSION).tar.gz -C mac agent -C ../build .
	tree -Csh $(PUB_DIR)

pub_local:
	$(call pub,local)

# pub_test:
# 	@echo "publish test ${BIN} ..."
# 	@go run make.go -pub -release test -pub-dir $(PUB_DIR) -name $(NAME)

pub_test_win:
	@echo "publish test windows ${BIN} ..."
	@go run make.go -pub -release test -pub-dir $(PUB_DIR) -archs "windows/amd64" -name $(NAME) -windows

pub_release_win:
	@echo "publish release windows ${BIN} ..."
	@go run make.go -pub -release release -pub-dir $(PUB_DIR) -archs "windows/amd64" -name $(NAME) -windows

pub_preprod:
	@echo "publish preprod ${BIN} ..."
	@go run make.go -pub -release preprod -pub-dir $(PUB_DIR) -name $(NAME)

# pub_release:
# 	@echo "publish release ${BIN} ..."
# 	@go run make.go -pub -release release -pub-dir $(PUB_DIR) -name $(NAME)

clean:
	rm -rf build/*
	rm -rf $(PUB_DIR)/*
