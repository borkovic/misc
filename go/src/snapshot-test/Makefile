
#	go build types.go proc.go driver.go

r:
	rm -f driver
	go run --race driver.go 2>&1 | tee r.log

b:
	go build driver.go


fmt:
	for f in *.go; do go fmt $$f; done

s:
	sha1sum *.go

c:
	rm -f driver
