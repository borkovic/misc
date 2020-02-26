
GOs := vc.go op.go types.go

top:
	go build vc.go types.go
	go build op.go vc.go types.go
	go build eng.go op.go vc.go types.go
	go build graph.go eng.go op.go vc.go types.go

fmt:
	go fmt $(GOs)

