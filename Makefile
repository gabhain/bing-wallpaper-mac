APP_NAME=BingWallpaper
APP_BUNDLE=$(APP_NAME).app
PKG_NAME=$(APP_NAME)-macOS.pkg
INSTALL_ROOT=install_root

build:
	@echo "Building Go Binary..."
	go build -ldflags "-s -w" -o bing-wallpaper main.go

app: build
	@echo "Structuring the .app bundle..."
	mkdir -p $(APP_BUNDLE)/Contents/MacOS
	mkdir -p $(APP_BUNDLE)/Contents/Resources
	cp bing-wallpaper $(APP_BUNDLE)/Contents/MacOS/$(APP_NAME)
	cp Info.plist $(APP_BUNDLE)/Contents/Info.plist
	cp AppIcon.icns $(APP_BUNDLE)/Contents/Resources/AppIcon.icns
	chmod +x $(APP_BUNDLE)/Contents/MacOS/$(APP_NAME)
	@echo "App bundle $(APP_BUNDLE) created successfully."

pkg: app
	@echo "Packaging into a macOS .pkg installer..."
	mkdir -p $(INSTALL_ROOT)/Applications
	cp -r $(APP_BUNDLE) $(INSTALL_ROOT)/Applications/
	pkgbuild --root $(INSTALL_ROOT) \
			 --identifier com.gabhain.bingwallpaper \
			 --version 1.0.3 \
			 --install-location / \
			 --scripts scripts \
			 $(PKG_NAME)
	@echo "$(PKG_NAME) created successfully."

clean:
	rm -rf $(APP_BUNDLE)
	rm -rf $(INSTALL_ROOT)
	rm -f bing-wallpaper
	rm -f $(PKG_NAME)
