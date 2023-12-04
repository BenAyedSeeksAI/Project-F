.PHONY: run seed

run:
	go run main.go

seed:
	go run main.go DbSeed
