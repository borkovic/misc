
b:
	go build types.go proc.go driver.go

fmt:
	for f in *.go; do go fmt $$f; done

vet:
	go vet *.go

s:
	sha1sum *.go
