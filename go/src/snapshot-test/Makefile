
#	go build types.go proc.go driver.go

r:
	rm -f driver
	time go run --race driver.go 2>&1 | tee r.log

f:
	rm -f driver
	time go run driver.go 2>&1 | tee f.log

five:
	rm -f driver
	N=5; n=0; while let "$$n < $$N"; do go run driver.go; echo; echo; let "n=$$n+1"; done

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
