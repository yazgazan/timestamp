timestamp
=========

A simple cli utility to prepend timestamps to commands.

```
Usage of timestamp:
  -delta
    	log time relative to last line
  -format string
    	date format (see https://golang.org/pkg/time/#Time.Format) (default "Mon Jan 2 15:04:05")
  -relative
    	log time relative to start
```

Examples:

```
$ ./test.sh
absolute date:
$ ./timestamp
Thu Feb 16 14:59:48: 0
Thu Feb 16 14:59:49: 1
Thu Feb 16 14:59:50: 2
time elapsed (relative from the start):
$ ./timestamp --relative
55.586µs: 0
1.000598304s: 1
2.004221692s: 2
time elapsed (relative from last line):
$ ./timestamp --delta
52.206µs: 0
1.003321839s: 1
1.007405483s: 2
```

