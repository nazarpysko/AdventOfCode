## Basic 
Make sure to **export sessionAoC environment variable** with your session cookie of [adventofcode.com](adventofcode.com).

To run a specific *XX day* and *Y part*, make sure to be in 2022 folder and execute the following:
```
go dayXX/common.go run dayXX/partY.go
```

## Makefile recipes

To create new day:
``` make
make new day=XX
```

To copy part1 to part2.go file
``` make
make copyPart1 day=XX
```
