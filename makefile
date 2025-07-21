VERSION = none
MESSAGE = Release version $(VERSION)

TARGET_TESTS = \
	./tests/rundown \
	./tests/backend \
	./tests/backend2 \
	./tests/fps

TARGET_EXAMPLES = \
	./examples/basic-window \
	./examples/fonts \
	./examples/images \
	./examples/shapes \
	./examples/transparency \
	./examples/two-windows \
	./examples/web

.PHONY: release upload tests examples

release:
	git add .
	git commit -m "$(MESSAGE)"
	git tag $(VERSION)
	git push origin $(VERSION)
	make proxy-release

upload:
	GOPROXY=proxy.golang.org go list -m vuelto.pp.ua@$(VERSION)

tests:
	@for dir in $(TARGET_TESTS); do \
		echo "Building $$dir..."; \
		go build -o bin/$$dir $$dir || exit 1; \
	done

examples:
	@for dir in $(TARGET_EXAMPLES); do \
		echo "Building $$dir..."; \
		go build -o bin/$$dir $$dir || exit 1; \
	done
