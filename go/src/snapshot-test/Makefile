
#	go build types.go proc.go driver.go

r:
	go run --race driver.go

b:
	go build driver.go


fmt:
	for f in *.go; do go fmt $$f; done

s:
	sha1sum *.go
