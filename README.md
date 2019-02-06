Insertion time
-----
On average 5s per 10,000 rows with unique index

500,000 accounts, 2 account keys, 1 application key, no index
-----
```
BenchmarkFindAccount-8                50          24286577 ns/op
PASS
ok      github.com/xinsnake/key-rotation        1.247s
```

500,000 accounts, 2 account keys, 1 application key, with unique index
-----
```
BenchmarkFindAccount-8             20000            100215 ns/op
PASS
ok      github.com/xinsnake/key-rotation        3.016s
```


500,000 accounts, 2 account keys, 10,000 application key, with unique index
-----
```
BenchmarkFindAccount-8                 1        1006530007 ns/op
PASS
ok      github.com/xinsnake/key-rotation        1.009s
```