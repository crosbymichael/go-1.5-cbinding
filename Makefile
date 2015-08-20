all:
	go build -buildmode c-archive stuff.go
	gcc main.c stuff.a -lpthread

clean:
	rm a.out
	rm stuff.h
	rm stuff.a
