# src file
SRCS=dp.cpp

# file to run
TARGET=$(SRCS:.cpp=)

# file to object
# OBJS=test.o
# replace filename *.c to *.o
OBJS=$(SRCS:.cpp=.o)   #same to OBJS=$(patsubst%.c,%.o,$(SRCS))   ---> test.0

# compliler option
CC=g++
#CFLAGS=-Wall -std=gnu99

$(TARGET):$(OBJS)
	@echo TARGET:$@
	@echo OBJS:$^
	$(CC) -std=c++11 -o  $@ $^  # stp2:equal to gcc -o test test.o



# % is wildcard charactre(通配符)
$(OBJS):$(SRCS)
	@echo TARGET:$@
	@echo OBJS:$^
	$(CC) -std=c++11 -g -o  $@ -c $^   # stp1: equal to gcc -o test.o -c test.c

# dp.o:dp.cpp
# 	@echo TARGET:$@
# 	@echo OBJS:$^
# 	g++ -g -o dp.o -c dp.cpp

# main.o:main.cpp
# 	@echo TARGET:$@
# 	@echo OBJS:$^
# 	$(CC) -std=c++11 -g -o  main.o -c main.cpp


clean:
	rm -rf  $(TARGET) $(OBJS)