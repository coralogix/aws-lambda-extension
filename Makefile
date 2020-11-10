.PHONY: vendor

fmt:
	@go fmt ./{cmd,pkg}/*

vet:
	@go vet ./{cmd,pkg}/*

clean:
	@rm -rf extensions layer.zip

build: clean fmt vet
	@CGO_ENABLED=0 go build -ldflags "-s -w" -o extensions/coralogix-extension cmd/coralogix-extension/main.go

package: build
	@zip -9 -q -r layer.zip extensions
	@rm -rf extensions

publish: package
	@aws lambda publish-layer-version \
		--layer-name coralogix-extension \
		--compatible-runtimes python3.6 python3.7 python3.8 \
		--description "Lambda function extension for logging to Coralogix" \
		--license-info "Apache-2.0" \
		--zip-file fileb://layer.zip \
		--output json | tee metadata.json