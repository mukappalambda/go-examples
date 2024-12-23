# Demo Usage of Delve

Delve github: https://github.com/go-delve/delve

## Install Delve

Ref: https://github.com/go-delve/delve/tree/master/Documentation/installation#installation

```bash
$ go version
go version go1.20.5 linux/amd64
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

After having delve installed, one can check the delve version:

```bash
$ dlv version
Delve Debugger
Version: 1.21.0
Build: $Id: fec0d226b2c2cce1567d5f59169660cf61dc1efe $
```

Change to the current directory, there is a `main.go` file. Executing this file will see:

```bash
$ go run main.go
{alex 20 [1.1 2.2 3.3]}
{alex 20 [1.1 2.2 3.3 4.4 5.5]}
```

Now we use delve to interact with `main.go`.

```bash
$ dlv debug main.go
Type 'help' for list of commands.
(dlv)
```

Show source code of `main.main`:

```bash
(dlv) list main.main
Showing /home/mklan/workspace/go-examples/debugging/delve-example/main.go:11 (PC: 0x49d152)
     6:         Name string
     7:         Age  int
     8:         Data []float64
     9: }
    10:
    11: func main() {
    12:         var p Person
    13:         p.Name = "alex"
    14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
```

Set a breakpoint in line 13:

```bash
(dlv) b main.go:13
Breakpoint 1 set at 0x49d184 for main.main() ./main.go:13
```

Again, this time we set another breakpoint in line 17:

```bash
(dlv) b main.go:17
Breakpoint 2 set at 0x49d2ca for main.main() ./main.go:17
```

To print out all the active breakpoints, type `bp`:

```bash
(dlv) bp
Breakpoint runtime-fatal-throw (enabled) at 0x437a60,0x437b60 for (multiple functions)() <multiple locations>:0 (0)
Breakpoint unrecovered-panic (enabled) at 0x437f00 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1145 (0)
        print runtime.curg._panic.arg
Breakpoint 1 (enabled) at 0x49d184 for main.main() ./main.go:13 (0)
Breakpoint 2 (enabled) at 0x49d2ca for main.main() ./main.go:17 (0)
```

Step over to next source line:

```bash
(dlv) n
> main.main() ./main.go:13 (hits goroutine(1):1 total:1) (PC: 0x49d184)
     8:         Data []float64
     9: }
    10:
    11: func main() {
    12:         var p Person
=>  13:         p.Name = "alex"
    14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
```

Print the variable `p`:

```bash
(dlv) p p
main.Person {
        Name: "",
        Age: 0,
        Data: []float64 len: 0, cap: 0, nil,}
```

Step over to the next source line:

```bash
(dlv) n
> main.main() ./main.go:14 (PC: 0x49d19f)
     9: }
    10:
    11: func main() {
    12:         var p Person
    13:         p.Name = "alex"
=>  14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
    19: }
```

Print the variable `p` again, you will see the value of `p.Name` becomes `"alex"`:

```bash
(dlv) p p
main.Person {
        Name: "alex",
        Age: 0,
        Data: []float64 len: 0, cap: 0, nil,}
```

Step over to the next line and then print the variable `p`. Now `p.Age` becomes `20`:

```bash
(dlv) n
> main.main() ./main.go:15 (PC: 0x49d1ab)
    10:
    11: func main() {
    12:         var p Person
    13:         p.Name = "alex"
    14:         p.Age = 20
=>  15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
    19: }
    20:
(dlv) p p
main.Person {
        Name: "alex",
        Age: 20,
        Data: []float64 len: 0, cap: 0, nil,}
```

Run to the next breakpoint:

```bash
(dlv) c
{alex 20 [1.1 2.2 3.3]}
> main.main() ./main.go:17 (hits goroutine(1):1 total:1) (PC: 0x49d2ca)
    12:         var p Person
    13:         p.Name = "alex"
    14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
=>  17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
    19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
```

Single step through the function:

```bash
(dlv) s
> main.(*Person).appendData() ./main.go:21 (PC: 0x49d3ef)
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
    19: }
    20:
=>  21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
    23:         p.Data = append(p.Data, s...)
    24:         *pp = p
    25: }
(dlv) n
> main.(*Person).appendData() ./main.go:22 (PC: 0x49d426)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
    19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
=>  22:         p := *pp
    23:         p.Data = append(p.Data, s...)
    24:         *pp = p
    25: }
```

Print the variables `pp` and `s`:

```bash
(dlv) p pp
*main.Person {
        Name: "alex",
        Age: 20,
        Data: []float64 len: 3, cap: 3, [1.1,2.2,3.3],}
(dlv) p s
[]float64 len: 2, cap: 2, [4.4,5.5]
```

Step over to the next two line, and print the variable `p`:

```bash
(dlv) n
> main.(*Person).appendData() ./main.go:23 (PC: 0x49d448)
    18:         fmt.Println(p)
    19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
=>  23:         p.Data = append(p.Data, s...)
    24:         *pp = p
    25: }
(dlv) n
> main.(*Person).appendData() ./main.go:24 (PC: 0x49d52c)
    19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
    23:         p.Data = append(p.Data, s...)
=>  24:         *pp = p
    25: }
(dlv) p p
main.Person {
        Name: "alex",
        Age: 20,
        Data: []float64 len: 5, cap: 6, [1.1,2.2,3.3,4.4,5.5],}
```

Step out of the function `appendData`:

```bash
(dlv) so
> main.main() ./main.go:18 (PC: 0x49d32d)
Values returned:

    13:         p.Name = "alex"
    14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
=>  18:         fmt.Println(p)
    19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
    23:         p.Data = append(p.Data, s...)
```

```bash
(dlv) p p
main.Person {
        Name: "alex",
        Age: 20,
        Data: []float64 len: 5, cap: 6, [1.1,2.2,3.3,4.4,5.5],}
(dlv) n
{alex 20 [1.1 2.2 3.3 4.4 5.5]}
> main.main() ./main.go:19 (PC: 0x49d3bd)
    14:         p.Age = 20
    15:         p.Data = []float64{1.1, 2.2, 3.3}
    16:         fmt.Println(p)
    17:         p.appendData([]float64{4.4, 5.5})
    18:         fmt.Println(p)
=>  19: }
    20:
    21: func (pp *Person) appendData(s []float64) {
    22:         p := *pp
    23:         p.Data = append(p.Data, s...)
    24:         *pp = p
```

```bash
(dlv) bp
Breakpoint runtime-fatal-throw (enabled) at 0x437b60,0x437a60 for (multiple functions)() <multiple locations>:0 (0)
Breakpoint unrecovered-panic (enabled) at 0x437f00 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1145 (0)
        print runtime.curg._panic.arg
Breakpoint 1 (enabled) at 0x49d184 for main.main() ./main.go:13 (1)
Breakpoint 2 (enabled) at 0x49d2ca for main.main() ./main.go:17 (1)
```

Delete the breakpoint 1:

```bash
(dlv) clear 1
Breakpoint 1 cleared at 0x49d184 for main.main() ./main.go:13
(dlv) bp
Breakpoint runtime-fatal-throw (enabled) at 0x437b60,0x437a60 for (multiple functions)() <multiple locations>:0 (0)
Breakpoint unrecovered-panic (enabled) at 0x437f00 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1145 (0)
        print runtime.curg._panic.arg
Breakpoint 2 (enabled) at 0x49d2ca for main.main() ./main.go:17 (1)
```

Delete all the breakpoints:

```bash
(dlv) clearall
Breakpoint 2 cleared at 0x49d2ca for main.main() ./main.go:17
(dlv) bp
Breakpoint runtime-fatal-throw (enabled) at 0x437b60,0x437a60 for (multiple functions)() <multiple locations>:0 (0)
Breakpoint unrecovered-panic (enabled) at 0x437f00 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1145 (0)
        print runtime.curg._panic.arg
```

Exit the debugger:

```bash
(dlv) q
$
```

That's the basic introduction to delve.
