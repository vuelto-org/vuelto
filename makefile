VERSION = none
MESSAGE = Release version $(VERSION)

TARGET_TESTS = \
	./tests/rundown \
	./tests/backend \
	./tests/backend2

TARGET_EXAMPLES = \
	./examples/basic-window \
	./examples/shapes \
	./examples/images \
	./examples/two-windows

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
		echo "building $$dir..."; \
		go build -o bin/$$dir $$dir || exit 1; \
	done

examples:
	@for dir in $(TARGET_EXAMPLES); do \
		echo "Building $$dir..."; \
		go build -o bin/$$dir $$dir || exit 1; \
	done
