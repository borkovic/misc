PROG=main

#	go build types.go proc.go $(PROG).go

r:
	@rm -f main
	@time go run --race $(PROG).go 2>&1 | tee r.log

f:
	@rm -f $(PROG)
	@time go run $(PROG).go 2>&1 | tee f.log

f2:
	@rm -f $(PROG)
	time go run $(PROG).go --chan-prob 2 100 2>&1 | tee f.log

5: five

five:
	@rm -f $(PROG)
	@go build --race $(PROG).go
	@N=5; echo; echo; n=0; while let "$$n < $$N"; do echo $$n': '; ./$(PROG) --nproc 1000; echo; let "n=$$n+1"; done


10: ten

ten:
	@rm -f $(PROG)
	@go build --race $(PROG).go
	@N=10; echo; echo; n=0; while let "$$n < $$N"; do echo $$n': '; ./$(PROG); echo; let "n=$$n+1"; done

100:
	@rm -f $(PROG)
	@go build $(PROG).go
	@N=100; echo; echo; n=0; while let "$$n < $$N"; do \
		echo $$n': '; \
		./$(PROG) --nproc 3000 --chan-prob 3 1000 ; \
		echo; let "n=$$n+1"; \
		done | tee log100

100r:
	@rm -f $(PROG)
	@go build --race $(PROG).go
	@N=100; echo; echo; n=0; while let "$$n < $$N"; do \
		echo $$n': '; \
		./$(PROG) --nproc 3000 --chan-prob 2 1000 ; \
		echo; let "n=$$n+1"; \
		done | tee log100r

20k:
	@rm -f $(PROG)
	@go build  $(PROG).go
	@N=100; echo; echo; n=0; while let "$$n < $$N"; do \
		echo $$n': '; \
		./$(PROG) --nproc 20000 --chan-prob 1 1000 ; \
		echo; let "n=$$n+1"; \
		done | tee log100r


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
