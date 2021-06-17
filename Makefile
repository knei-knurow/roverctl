GC := go
EXE := 

all: main.go movement.go
	$(GC) build -o roverctl main.go movement.go

clean:
	rm ./roverctl
