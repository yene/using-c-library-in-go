# How to use a C lib from Go

## Build
```bash
# Compile hiya into a static library (all code in one file).
gcc -Wall -g -c hiya.c -o hiya.o
ar ruv libhiya.a hiya.o
ranlib libhiya.a

# Compile hiya into a shared library (has external dependencies like standard libs).
gcc -c -fPIC hiya.c -o hiya.o
gcc -shared -W1,-soname,libhiya.so.1 -o libhiya.so.1.0.0 hiya.o

go run main.go
```

## Test
```
gcc -o test -I. -L. -lhiya -Wall -g test.c
./test
```

## Links
* [Example cgo (Golang) app that calls a native library with a C structure](http://www.mischiefblog.com/2014/06/26/example-cgo-golang-app-that-calls-a-native-library-with-a-c-structure/)
