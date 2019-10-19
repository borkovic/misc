PROG=main

#	go build types.go proc.go $(PROG).go

r:
	rm -f main
	time go run --race $(PROG).go 2>&1 | tee r.log

f:
	rm -f $(PROG)
	time go run $(PROG).go 2>&1 | tee f.log

five:
	rm -f $(PROG)
	N=5; echo; n=0; while let "$$n < $$N"; do go run $(PROG).go; echo; echo; let "n=$$n+1"; done

b:
	time go build $(PROG).go


fmt:
	for f in *.go; do go fmt $$f; done

vet:
	go vet *.go

s:
	sha1sum *.go

c:
	rm -f $(PROG)
