N = 200000000

CFLAGS   = -std=c99   -O3 -Wall -Werror
CPPFLAGS = -std=c++1y -O3 -Wall -Werror 
################################################
EXECsGo = hsg hsgf hsgs
EXECsCC = hscc hsccf hsc
EXECs = $(EXECsGo) $(EXECsCC)

top: $(EXECs)
	for f in $(EXECs); do \
		time ./$$f $(N); \
		echo; echo; \
	done;

.PHONY: $(EXECs)

execs: $(EXECs)

################################################
clean:
	rm -f $(EXECs) *.o

################################################
%: %.go
	@echo '====================================='
	time go build $?
	echo; echo

################################################
% : %.cc
	@echo '====================================='
	time g++ $(CPPFLAGS) -o $@ $? pl.c
	echo; echo

################################################
% : %.c
	@echo '====================================='
	time gcc $(CFLAGS) -o $@ $? pl.c
	echo; echo

