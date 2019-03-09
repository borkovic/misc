################################################
top: hs1 hs2 hs3

################################################
hs1: hs1.go 
	go build hs1.go 
	time ./hs1

################################################
#CPPFLAGS = -O3 -Wall -Werror -std=c++11
#CPPFLAGS = -O3 -Wall -Werror -std=c++14
CPPFLAGS = -O3 -Wall -Werror -std=c++1y

hs2 : hs2.o
	g++ $(CPPFLAGS) -o hs2 hs2.o
	time ./hs2

hs2.o: hs2.cc
	g++ $(CPPFLAGS) -c hs2.cc

################################################
CFLAGS = -std=c99 -O3 -Wall -Werror

hs3 : hs3.o
	gcc $(CFLAGS) -o hs3 hs3.o
	time ./hs3

hs3.o: hs3.c
	gcc $(CFLAGS) -c hs3.c

