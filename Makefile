build:
	go build main.go

test: build
	./main test.wav 

clean:
	- rm -f main test.mp3 test.aac test.flac test.ogg
