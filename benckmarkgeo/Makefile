build:
	go build -o benchmark main.go

gen:
	make build
	./benchmark gen -f geo.txt -n 10000000
	./benchmark gen -f geo_search.txt -n 10000

benchmark_redis:
	make build
	./benchmark redis init -f geo.txt >> benchmark.log
	./benchmark redis -f geo_search.txt >> benchmark.log
