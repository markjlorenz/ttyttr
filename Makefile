build:
	go build -o bin/ttyttr ttyttr.go

test:
	go run ttyttr.go 434135040256008192
	go run ttyttr.go 444239051965882368
	go run ttyttr.go 442781697596092418
	go run ttyttr.go 415804923591147520

.PHONEY: test
