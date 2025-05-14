VERSION = none
MESSAGE = Release version $(VERSION)

TARGETS = \
	./tests/rundown \
	./tests/backend \
	./tests/backend2 \
	./examples/basic-window \
	./examples/shapes \
	./examples/images \
	./examples/two-windows

.PHONY: release upload test

release:
	git add .
	git commit -m "$(MESSAGE)"
	git tag $(VERSION)
	git push origin $(VERSION)
	make proxy-release

upload:
	GOPROXY=proxy.golang.org go list -m vuelto.pp.ua@$(VERSION)

test:
	@for dir in $(TARGETS); do \
		echo "Building $$dir..."; \
		go build -o bin/$$dir $$dir || exit 1; \
	done
