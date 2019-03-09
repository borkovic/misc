################################################
top: hsg hscc hsc

.PHONY: hsg hscc hsc

################################################
clean:
	rm -f hsg hscc hsc
################################################
hsg: hsg.go 
	@echo '====================================='
	time go build hsg.go 
	echo
	time ./hsg
	echo; echo

################################################
#CPPFLAGS = -O3 -Wall -Werror -std=c++11
#CPPFLAGS = -O3 -Wall -Werror -std=c++14
CPPFLAGS = -O3 -Wall -Werror -std=c++1y

hscc : hscc.cc
	@echo '====================================='
	time g++ $(CPPFLAGS) -o hscc hscc.cc
	echo
	time ./hscc
	echo; echo

################################################
CFLAGS = -std=c99 -O3 -Wall -Werror

hsc : hsc.c
	@echo '====================================='
	time gcc $(CFLAGS) -o hsc hsc.c
	echo
	time ./hsc
	echo; echo

