STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/steampipe@latest/steampipe-plugin-steampipe.plugin -tags "${BUILD_TAGS}" *.go
