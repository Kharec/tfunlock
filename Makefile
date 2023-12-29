TARGET := tfunlock

.PHONY: all clean

all:
	go get .
	go build -o $(TARGET)

clean:
	rm -f $(TARGET)

