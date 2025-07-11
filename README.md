# gsplitter
Allows you to divide the files in the current folder into many smaller directories. Division options: by number of files, by extension, by creation date, or by first letter.

## Installation
```
go install github.com/HardDie/gsplitter@latest
```

## Usage

<details>
	<summary>Split files by count of files</summary>

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
$ gsplitter count -i 200
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
</details>

<details>
	<summary>Split files by extension</summary>

```bash
$ ls -l
-rw-r--r-- 1 user users     0 Feb 14 18:21 0001.txt
-rw-r--r-- 1 user users     0 Feb 14 18:21 0002.TXT
-rw-r--r-- 1 user users     0 Feb 14 18:21 0003.jpg
-rw-r--r-- 1 user users     0 Feb 14 18:21 0004.JPG
...
-rw-r--r-- 1 user users     0 Feb 14 18:21 1000
```
```bash
$ gsplitter ext
2023/02/14 18:24:57 Done!
```
```bash
$ ls -l
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 txt/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 jpg/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 unknown/
```
</details>

<details>
	<summary>Split files by first letter</summary>

```bash
$ ls -l
-rw-r--r-- 1 user users     0 Feb 14 18:21 aa
-rw-r--r-- 1 user users     0 Feb 14 18:21 ab
-rw-r--r-- 1 user users     0 Feb 14 18:21 ba
-rw-r--r-- 1 user users     0 Feb 14 18:21 Bb
```
```bash
$ gsplitter letter
2023/02/14 18:24:57 Done!
```
```bash
$ ls -l
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 A/
drwxr-xr-x 2 user users 4.0K Feb 14 18:24 B/
```
</details>

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
