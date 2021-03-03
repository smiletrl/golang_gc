## Example description

This is a simple http server `http:/127.0.0.1:8080`, with only one registered endpoint `/foo`.

According to [escape analysis](https://github.com/smiletrl/golang_escape/blob/master/pkg/escape/README.md), variable `var f []employee` from `line 40` at `main.go` will escape to heap. After `line 34` execution, this variable will no longer be used, i.e, it will become garbage in heap memory. GC will find it, and free its memory.

A request to endpoint `http:/127.0.0.1:8080/foo` will triger this `expensiveCall()`, and then we can observe the gc behaviour with some commands.

## Commands to run

Use `GODEBUG="gctrace=1"` to enable garbage collection trace

```
go build .
GODEBUG="gctrace=1" ./example1
// open a new terminal, and run this command
curl http:/127.0.0.1:8080/foo
```

Result looks like

```
smiletrl@Rulins-MacBook-Pro example1 % GODEBUG="gctrace=1" GOGC=100  ./example1
gc 1 @3.842s 0%: 0.029+0.31+0.005 ms clock, 0.46+0/0.42/0.25+0.084 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @3.845s 0%: 0.026+0.26+0.004 ms clock, 0.42+0/0.32/0.24+0.079 ms cpu, 10->10->10 MB, 11 MB goal, 16 P
gc 3 @3.851s 0%: 0.061+0.29+0.006 ms clock, 0.98+0/0.35/0.28+0.098 ms cpu, 20->20->20 MB, 21 MB goal, 16 P
gc 4 @3.863s 0%: 0.032+0.30+0.004 ms clock, 0.51+0/0.34/0.22+0.071 ms cpu, 40->40->40 MB, 41 MB goal, 16 P
gc 5 @3.885s 0%: 0.059+0.26+0.004 ms clock, 0.95+0/0.21/0.17+0.068 ms cpu, 80->80->80 MB, 81 MB goal, 16 P
gc 6 @7.660s 0%: 0.12+0.19+0.004 ms clock, 2.0+0/0.23/0.21+0.079 ms cpu, 157->157->157 MB, 161 MB goal, 16 P
gc 7 @7.732s 0%: 0.033+0.29+0.004 ms clock, 0.53+0/0.24/0.29+0.072 ms cpu, 308->308->308 MB, 315 MB goal, 16 P
gc 8 @7.848s 0%: 0.058+0.24+0.005 ms clock, 0.93+0/0.27/0.21+0.085 ms cpu, 603->603->268 MB, 617 MB goal, 16 P
gc 9 @7.896s 0%: 0.039+0.27+0.005 ms clock, 0.63+0/0.26/0.46+0.087 ms cpu, 526->526->191 MB, 536 MB goal, 16 P
gc 10 @7.928s 0%: 0.058+0.20+0.004 ms clock, 0.93+0/0.27/0.18+0.077 ms cpu, 375->375->40 MB, 382 MB goal, 16 P
gc 11 @7.935s 0%: 0.056+0.23+0.004 ms clock, 0.91+0/0.27/0.14+0.078 ms cpu, 80->80->80 MB, 81 MB goal, 16 P
gc 12 @7.948s 0%: 0.057+0.22+0.005 ms clock, 0.91+0/0.29/0.18+0.082 ms cpu, 157->157->157 MB, 161 MB goal, 16 P
gc 13 @7.974s 0%: 0.058+0.20+0.005 ms clock, 0.94+0/0.27/0.24+0.082 ms cpu, 308->308->308 MB, 315 MB goal, 16 P
gc 14 @8.027s 0%: 0.063+0.28+0.005 ms clock, 1.0+0/0.29/0.23+0.081 ms cpu, 603->603->268 MB, 617 MB goal, 16 P
...
gc 205 @9.612s 0%: 0.061+0.28+0.005 ms clock, 0.98+0/0.29/0.14+0.080 ms cpu, 308->308->308 MB, 315 MB goal, 16 P
druation is: 5.775688745s
GC forced
gc 206 @133.539s 0%: 0.078+0.22+0.004 ms clock, 1.2+0/0.44/0.26+0.065 ms cpu, 335->335->0 MB, 617 MB goal, 16 P
GC forced
gc 207 @253.540s 0%: 0.11+0.43+0.013 ms clock, 1.8+0/0.80/0.22+0.20 ms cpu, 0->0->0 MB, 4 MB goal, 16 P
GC forced
gc 208 @373.543s 0%: 0.074+0.25+0.006 ms clock, 1.1+0/0.50/0.18+0.11 ms cpu, 0->0->0 MB, 4 MB goal, 16 P
GC forced
gc 209 @524.928s 0%: 0.038+0.18+0.005 ms clock, 0.61+0/0.37/0.45+0.085 ms cpu, 0->0->0 MB, 4 MB goal, 16 P
...
```

Description for the meaning of each line item's sections can be found at [Go Runtime doc](https://golang.org/pkg/runtime/). Find this line `gctrace: setting gctrace=1 causes the garbage collector to emit a single line to standard...` from that doc.

We will see the initial gc started with a heap goal 5MB. In other words, when memory allocated at heap grows closer to 5MB (in this example's case, it is 4MB), gc gets started. Later on, gc will set up the heap size goal to be double of previous gc's live heap(i.e, heap memory left after previous gc clean). When heap size grows closer to the heap goal again, gc will be triggered again.

Gc cycle `2-7` don't actually free any memory, as heap size keeps growing doubly. It means the variable `var f []employee` is still growing inside code block between `line 41` to `line 56`.

From gc cycle `gc 8-10`, the live heap size after gc gets reduced a lot, from 603MB to 40MB. It means the first variable `var f []employee` (this variable has been declared 100 times at `line 32`) finally finishes its life span at `line 40` from `main.go`. Then Gc starts freeing it.

Next GC cycle will repeat above process until this endpoint `/foo` finishes processing.

Finally after gc 205, process is complete. But there're still 308MB live heap size remaining. We know GC will set up next gc heap size goal to be 617MB (double 308MB), i.e, next gc won't start until heap size grows near 617MB. But this process has already finished, will this heap size never be freed until a new request to endpoint `/foo` is made again?

Then we see `gc 206, 207 - ...`. Gc will be forced to run every 120 seconds (2 minutes). We see nothing left on heap anymore after gc 206. We don't send request to `/foo`, no heap size grows, and gc is always forced to run every 2 minutes.

## GOGC parameter

In the above example, GC starts when the live heap size grows near a goal. And the goal changes from time to time. The formula to get the goal is `goal = (1 + GOGC/100) * previous-live-heap-size`. The default value of GOGC is 100. We might be interested about the first live heap size used for this formula. The initial heap goal set up in the above example is 5MB.

We may play with this parameter like

```
go build .
GODEBUG="gctrace=1" GOGC=200 ./example1
// open a new terminal, and run this command
curl http:/127.0.0.1:8080/foo
```

Result looks like

```
smiletrl@Rulins-MacBook-Pro example1 % GODEBUG="gctrace=1" GOGC=200 ./example1
gc 1 @3.292s 0%: 0.055+0.33+0.007 ms clock, 0.89+0/0.41/0.26+0.11 ms cpu, 11->11->10 MB, 12 MB goal, 16 P
gc 2 @3.302s 0%: 0.028+0.24+0.004 ms clock, 0.45+0/0.30/0.23+0.076 ms cpu, 30->30->30 MB, 31 MB goal, 16 P
gc 3 @3.325s 0%: 0.035+0.27+0.005 ms clock, 0.56+0/0.30/0.48+0.082 ms cpu, 84->84->84 MB, 91 MB goal, 16 P
gc 4 @3.399s 0%: 0.049+0.23+0.005 ms clock, 0.78+0/0.41/0.38+0.088 ms cpu, 241->241->241 MB, 252 MB goal, 16 P
gc 5 @3.594s 0%: 0.034+0.18+0.004 ms clock, 0.55+0/0.23/0.21+0.076 ms cpu, 700->700->30 MB, 724 MB goal, 16 P
gc 6 @3.605s 0%: 0.038+0.25+0.006 ms clock, 0.61+0/0.33/0.42+0.10 ms cpu, 90->90->90 MB, 91 MB goal, 16 P
gc 7 @3.643s 0%: 0.063+0.25+0.005 ms clock, 1.0+0/0.38/0.36+0.080 ms cpu, 265->265->265 MB, 272 MB goal, 16 P
gc 8 @3.756s 0%: 0.035+0.23+0.005 ms clock, 0.56+0/0.32/0.32+0.089 ms cpu, 771->771->100 MB, 795 MB goal, 16 P
gc 9 @3.791s 0%: 0.040+0.24+0.005 ms clock, 0.65+0/0.33/0.34+0.080 ms cpu, 295->295->295 MB, 302 MB goal, 16 P
gc 10 @3.927s 0%: 0.033+0.25+0.005 ms clock, 0.53+0/0.35/0.25+0.091 ms cpu, 858->858->188 MB, 885 MB goal, 16 P
gc 11 @3.999s 0%: 0.064+0.23+0.006 ms clock, 1.0+0/0.32/0.38+0.10 ms cpu, 546->546->211 MB, 564 MB goal, 16 P
gc 12 @4.081s 0%: 0.066+0.22+0.005 ms clock, 1.0+0/0.29/0.36+0.081 ms cpu, 613->613->278 MB, 634 MB goal, 16 P
gc 13 @4.191s 0%: 0.034+0.27+0.005 ms clock, 0.55+0/0.38/0.41+0.087 ms cpu, 808->808->137 MB, 835 MB goal, 16 P
gc 14 @4.245s 0%: 0.057+0.41+0.006 ms clock, 0.91+0/0.49/0.52+0.10 ms cpu, 402->402->67 MB, 413 MB goal, 16 P
gc 15 @4.271s 0%: 0.033+0.21+0.004 ms clock, 0.53+0/0.32/0.21+0.078 ms cpu, 198->198->198 MB, 202 MB goal, 16 P
...
gc 101 @10.007s 0%: 0.034+0.23+0.005 ms clock, 0.55+0/0.34/0.37+0.087 ms cpu, 701->701->30 MB, 724 MB goal, 16 P
gc 102 @10.018s 0%: 0.033+0.21+0.005 ms clock, 0.53+0/0.23/0.23+0.086 ms cpu, 90->90->90 MB, 91 MB goal, 16 P
gc 103 @10.049s 0%: 0.067+0.24+0.005 ms clock, 1.0+0/0.35/0.29+0.080 ms cpu, 265->265->265 MB, 272 MB goal, 16 P
druation is: 6.839089277s
GC forced
gc 104 @131.184s 0%: 0.071+0.33+0.013 ms clock, 1.1+0/0.73/0.48+0.21 ms cpu, 670->670->0 MB, 795 MB goal, 16 P
GC forced
gc 105 @287.712s 0%: 0.052+0.17+0.003 ms clock, 0.83+0/0.32/0.081+0.063 ms cpu, 0->0->0 MB, 8 MB goal, 16 P
GC forced
gc 106 @423.069s 0%: 0.043+0.17+0.005 ms clock, 0.70+0/0.39/0.17+0.093 ms cpu, 0->0->0 MB, 8 MB goal, 16 P
...
```

## Turn off GC

We may turn off GC entirely with this parameter `GOGC` as

```
go build .
GODEBUG="gctrace=1" GOGC=off ./example1
// open a new terminal, and run this command
curl http:/127.0.0.1:8080/foo
```

Result looks like

```
smiletrl@Rulins-MacBook-Pro example1 % GODEBUG="gctrace=1" GOGC=off ./example1
druation is: 28.177963689s
```

No GC cycle is logged. The interesting thing is with GC disabled, the performance is actually getting better. This tells us GC is not always a bad thing for performance.

## GC pacer

GC sets up the heap goal with GOGC (ratio based on previous live heap size). But in the above examples, we can see the next GC cycle starts with a smaller heap size than the goal size. This is controlled by GC pacer. We might cover GC pacer in following examples.
