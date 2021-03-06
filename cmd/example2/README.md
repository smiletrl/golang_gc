## Example description

An example of [memory leak](https://en.wikipedia.org/wiki/Memory_leak).

This is similar as example1, but with a gloabl variable to store the produced `employee`. It means this global variable `var emps []employee` will always be in use, and never be freed in GC.

## Commands to run

```
cd cmd/example2
go build .
GODEBUG="gctrace=1" ./example2
// open a new terminal, and run this command
curl http:/127.0.0.1:8080/foo
```

Result looks like

```
smiletrl@Rulins-MacBook-Pro example2 % GODEBUG="gctrace=1" ./example2
gc 1 @1.318s 0%: 0.025+0.27+0.004 ms clock, 0.41+0/0.31/0.27+0.072 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @1.321s 0%: 0.048+0.23+0.004 ms clock, 0.78+0/0.23/0.20+0.076 ms cpu, 10->10->10 MB, 11 MB goal, 16 P
gc 3 @1.325s 0%: 0.055+0.23+0.003 ms clock, 0.88+0/0.27/0.26+0.063 ms cpu, 20->20->20 MB, 21 MB goal, 16 P
gc 4 @1.333s 0%: 0.057+0.28+0.003 ms clock, 0.91+0/0.40/0.28+0.062 ms cpu, 40->40->40 MB, 41 MB goal, 16 P
gc 5 @1.348s 0%: 0.057+0.20+0.004 ms clock, 0.91+0/0.22/0.23+0.077 ms cpu, 80->80->80 MB, 81 MB goal, 16 P
gc 6 @1.384s 0%: 0.064+0.28+0.005 ms clock, 1.0+0/0.31/0.16+0.082 ms cpu, 157->157->157 MB, 161 MB goal, 16 P
gc 7 @1.458s 0%: 0.032+0.19+0.005 ms clock, 0.52+0/0.25/0.23+0.086 ms cpu, 308->308->308 MB, 315 MB goal, 16 P
gc 8 @1.598s 0%: 0.053+0.19+0.004 ms clock, 0.85+0/0.28/0.34+0.070 ms cpu, 603->603->603 MB, 617 MB goal, 16 P
gc 9 @1.827s 0%: 0.060+0.27+0.004 ms clock, 0.96+0/0.32/0.29+0.078 ms cpu, 1180->1180->1180 MB, 1207 MB goal, 16 P
gc 10 @2.273s 0%: 0.033+0.25+0.004 ms clock, 0.53+0/0.33/0.33+0.068 ms cpu, 2303->2303->2302 MB, 2360 MB goal, 16 P
gc 11 @3.164s 0%: 0.082+0.51+0.004 ms clock, 1.3+0/0.35/1.2+0.076 ms cpu, 4492->4492->4491 MB, 4605 MB goal, 16 P
gc 12 @5.012s 0%: 0.061+0.82+0.004 ms clock, 0.97+0/0.38/1.9+0.075 ms cpu, 8759->8759->8758 MB, 8983 MB goal, 16 P
gc 13 @14.595s 0%: 0.93+21+0.012 ms clock, 14+0.76/60/180+0.19 ms cpu, 17079->17089->17087 MB, 17516 MB goal, 16 P
gc 14 @27.192s 0%: 1.8+7.4+0.006 ms clock, 30+1.1/8.4/56+0.10 ms cpu, 33320->33326->33322 MB, 34174 MB goal, 16 P
druation is: 26.086605042s
GC forced
gc 15 @147.416s 0%: 0.098+3.7+0.006 ms clock, 1.5+0/2.9/19+0.10 ms cpu, 33517->33517->33516 MB, 66645 MB goal, 16 P
GC forced
gc 16 @267.431s 0%: 0.088+3.1+0.007 ms clock, 1.4+0/4.9/15+0.12 ms cpu, 33516->33516->33516 MB, 67033 MB goal, 16 P
GC forced
gc 17 @387.445s 0%: 0.054+3.0+0.005 ms clock, 0.87+0/4.1/9.7+0.088 ms cpu, 33516->33516->33516 MB, 67033 MB goal, 16 P
...
```

We will find out GC will not reclaim the memory used by global variable. Because this varialble `var emps []employee` keeps growing with new requests, and never be released, this turns out to be a memory leak case.
