CP = gcc
LIB_FILE = bin/lib/libFesti.so
SOURCES = $(wildcard src/festi/*.c)
OBJECTS = $(patsubst %.c,%.o,$(SOURCES))
FLAGS = -Wdeprecated-declarations -Wdeprecated

.PHONY: build test

build: clean $(LIB_FILE)

$(LIB_FILE): $(OBJECTS)
	$(CP) -shared -o $@ $^ -lglfw -framework OpenGL $(FLAGS)

test: build
	$(CP) -o bin/test/test test/test.c -Isrc -Lbin/lib -lFesti 
	./bin/test/test
	make clean

clean:
	rm -f $(LIB_FILE) $(OBJECTS)
