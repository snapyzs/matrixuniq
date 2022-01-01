## Uniq Randmon Matrix
### Заполняет матрицу рандомными уникальными числами в указанном диапазоне
```
$ go run main.go -w 10 -h 10 -max 100
[2 66 43 82 72 11 59 39 56 4]
[83 20 30 33 42 81 22 50 13 86]
[40 28 9 27 16 85 44 14 37 71]
[65 73 84 77 53 67 32 58 99 18]
[87 19 7 46 78 48 24 36 94 92]
[25 76 23 74 5 68 52 29 60 88]
[57 95 98 26 47 1 91 10 41 70]
[21 63 31 45 15 62 49 61 3 8]
[35 90 93 64 6 79 69 97 51 38]
[17 75 96 0 12 55 89 34 54 80]

```
```
$ go test -v
=== RUN   TestRandomMatrix
=== RUN   TestRandomMatrix/Case(2,2,100)
=== RUN   TestRandomMatrix/Case(4,4,150)
=== RUN   TestRandomMatrix/Case(5,10,65)
--- PASS: TestRandomMatrix (0.00s)
    --- PASS: TestRandomMatrix/Case(2,2,100) (0.00s)
    --- PASS: TestRandomMatrix/Case(4,4,150) (0.00s)
    --- PASS: TestRandomMatrix/Case(5,10,65) (0.00s)
=== RUN   TestUniqNumber
=== RUN   TestUniqNumber/Case(2,2,100)
=== RUN   TestUniqNumber/Case(4,4,150)
=== RUN   TestUniqNumber/Case(10,5,65)
--- PASS: TestUniqNumber (0.00s)
    --- PASS: TestUniqNumber/Case(2,2,100) (0.00s)
    --- PASS: TestUniqNumber/Case(4,4,150) (0.00s)
    --- PASS: TestUniqNumber/Case(10,5,65) (0.00s)
PASS
ok      github.com/snapyzs/matrixrandom 0.006s

```