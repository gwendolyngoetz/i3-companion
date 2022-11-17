build: clean
	@go build -ldflags="-X 'main.Version=v$(VERSION)'" -o build/i3-companion cmd/i3-companion/main.go
	@echo Built i3-companion

clean:
	@rm -rf build
	@echo Cleaned

install: uninstall
	@[ -d $${HOME}/.local/bin ] || @mkdir -p $${HOME}/.local/bin
	@cp build/i3-companion $${HOME}/.local/bin
	@echo Installed i3-companion

uninstall:
	@rm $${HOME}/.local/bin/i3-companion
	@echo Uninstalled i3-companion

define newline


endef

define CONTROL_FILE_BODY
Package: i3-companion
Version: $(VERSION)
Section: base
Priority: optional
Architecture: x64
Maintainer: Gwendolyn Goetz
Description: Companion helpers for i3wm
endef

package-deb:
	@mkdir -p ./package/i3-companion_$(VERSION)/tmp/usr/local/bin
	@mkdir -p ./package/i3-companion_$(VERSION)/DEBIAN

	@cp ./build/i3-companion ./package/i3-companion_$(VERSION)/tmp/usr/local/bin
	@touch ./package/i3-companion_$(VERSION)/DEBIAN/control
	
	@echo '$(subst $(newline),\n,${CONTROL_FILE_BODY})' > ./package/i3-companion_$(VERSION)/DEBIAN/control

	@dpkg-deb --build ./package/i3-companion_$(VERSION)
	@echo Package deb

package: package-deb
