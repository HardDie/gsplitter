# gsplitter
Allows you to split files in the current folder into many small directories.

## Example
```bash
$ ls -l
-rw-r--r-- 1 user users     0 Feb 14 18:21 0001.txt
-rw-r--r-- 1 user users     0 Feb 14 18:21 0002.txt
-rw-r--r-- 1 user users     0 Feb 14 18:21 0003.txt
-rw-r--r-- 1 user users     0 Feb 14 18:21 0004.txt
...
-rw-r--r-- 1 user users     0 Feb 14 18:21 1000.txt
```
```bash
$ gsplitter -count 200
2023/02/14 18:24:57 Done!
```
```bash
$ ls -l
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 0/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 1/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 2/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 3/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 4/
```

## How to install
```
go install github.com/HardDie/gsplitter@latest
```
