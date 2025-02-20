
# Variables
CC = gcc
CFLAGS = -Wall -g

# Targets
all: hello

hello: hello.o
	$(CC) $(CFLAGS) -o hello hello.o

hello.o: hello.c
	$(CC) $(CFLAGS) -c hello.c

clean:
	rm -f hello.o hello Login to the Bito extension to get free AI code completions…
it could write 50% of your code…

build : rm day4-app1-inventory-app-for-grocery-store go build -race   ./day4-app1-inventory-app-for-grocery-store
newapp : mkdir $1 && touch main.go && go mod init $1