# file to run
TARGET=main

# src file
SRCS=main.cpp

# file to object
# OBJS=test.o
# replace filename *.c to *.o
OBJS=$(SRCS:.cpp=.o)   #same to OBJS=$(patsubst%.c,%.o,$(SRCS))   ---> test.0

# compliler option
CC=g++
#CFLAGS=-Wall -std=gnu99




$(TARGET):$(OBJS)
# TARGET $@
# OBJS   $^
	@echo TARGET:$@
	@echo OBJS:$^
	$(CC) -std=c++11 -o  $@ $^  # stp2:equal to gcc -o test test.o

clean:
	rm -rf  $(TARGET) $(OBJS)

%.o:%.cpp
	@echo TARGET:$@
	@echo OBJS:$<
	$(CC) -std=c++11 -g -o  $@ -c $<   # stp1: equal to gcc -o test.o -c test.c
