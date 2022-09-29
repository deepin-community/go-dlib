GOPATH_DIR = gopath
GOPKG_PREFIX = pkg.deepin.io/lib

prepare:
	@if [ ! -d ${GOPATH_DIR}/src/${GOPKG_PREFIX} ]; then \
		mkdir -p ${GOPATH_DIR}/src/$(dir ${GOPKG_PREFIX}); \
		ln -sf ../../.. ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
		fi

print_gopath: prepare
	GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}"

clean:
	rm -rf ${GOPATH_DIR}

check_code_quality: prepare
	env GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}" go vet ./...

test: prepare
	env GOPATH="${CURDIR}/${GOBUILD_DIR}:${GOPATH}" go test -v ./...

test-coverage: prepare
	env GOPATH="${CURDIR}/${GOBUILD_DIR}:${GOPATH}" go test -cover -v ./... | awk '$$1 ~ "^(ok|\\?)" {print $$2","$$5}' | sed "s:${CURDIR}::g" | sed 's/files\]/0\.0%/g' > coverage.csv
