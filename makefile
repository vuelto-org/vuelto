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

define BUILD_ALL
@for dir in $($1); do \
    t=$$(perl -MTime::HiRes -e 'print Time::HiRes::time()'); \
    printf "\033[1m\033[36mBuilding\033[0m $$dir...\n"; \
    go build -o bin/$$dir $$dir || exit 1; \
    dur=$$(perl -MTime::HiRes -e 'printf "%.2f", Time::HiRes::time() - shift' $$t); \
    printf "\033[1m\033[32mBuilt\033[0m $$dir in $${dur}s\n"; \
    \
    t=$$(perl -MTime::HiRes -e 'print Time::HiRes::time()'); \
    printf "\033[1m\033[36mBuilding w/o CGO:\033[0m $$dir...\n"; \
    CGO_ENABLED=0 go build -gcflags="github.com/ebitengine/purego/internal/fakecgo=-std" -o bin/$$dir $$dir || exit 1; \
    dur=$$(perl -MTime::HiRes -e 'printf "%.2f", Time::HiRes::time() - shift' $$t); \
    printf "\033[1m\033[32mBuilt w/o CGO\033[0m $$dir in $${dur}s\n"; \
done
endef

.PHONY: release upload tests examples

release:
	git add .
	git commit -m "$(MESSAGE)"
	git tag $(VERSION)
	git push origin $(VERSION)
	make upload VERSION=$(VERSION)

upload:
	GOPROXY=proxy.golang.org go list -m vuelto.pp.ua@$(VERSION)

tests:
	$(call BUILD_ALL,TARGET_TESTS)

examples:
	$(call BUILD_ALL,TARGET_EXAMPLES)
