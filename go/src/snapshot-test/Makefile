
#	go build types.go proc.go driver.go

r:
	rm -f driver
	time go run --race driver.go 2>&1 | tee r.log

f:
	rm -f driver
	time go run driver.go 2>&1 | tee f.log

five:
	rm -f driver
	for f in 0 1 2 3 4; do go run driver.go; echo; echo; done

b:
	time go build driver.go


fmt:
	for f in *.go; do go fmt $$f; done

vet:
	go vet *.go

s:
	sha1sum *.go

c:
	rm -f driver
