## Getting Started

[odin installation](https://odin-lang.org/docs/install/) wanted me to downloaded a nightly build and install that.  MacOS wanted none of that, since the executables and shared libraries were not signed.

Instead, I winged it:
```
$ brew install odin
$ which odin
/usr/local/bin/odin
$ odin version 
odin version dev-2025-04:d9f990d42
```

## Running the tests
```
$ odin test tests
```
or
```
$ odin test tests -define:ODIN_TEST_THREADS=1
```

## Running the program
```
$ odin run .
```
