##	Duplicate seacher v0.1.4

Search duplicates in all files in current directory.

![Alt text](/screenshot.jpg?raw=true "Usage")

Sequence of the option keys is not critical.

12.09.2015 - Firts commit.

13.09.2015 - Algorithm for finding duplicates has been optimized.

13.10.2015 - Changed algorithm and view for progress bar. Fixed calculation for strings in files.

### Install:
```
go get github.com/gkiryaziev/go-duplicate-seacher
```

### Build and Run
```
go build && go-duplicate-seacher
```

### Usage:
```
go-duplicate-seacher -ext dic -new Dict_New.dic
```