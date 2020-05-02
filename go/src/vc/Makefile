
GOs := ext_ifc.go vc.go op.go types.go

top:
	go build ext_ifc.go vc.go types.go
	go build ext_ifc.go op.go vc.go types.go
	go build ext_ifc.go eng.go op.go vc.go types.go
	go build ext_ifc.go graph.go eng.go op.go vc.go types.go

fmt:
	go fmt $(GOs)

