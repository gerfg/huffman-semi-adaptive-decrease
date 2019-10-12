# huffman-semi-adaptive-decrease

1. Build the execution file
```shell
$ go build -o huffman
```

2. Execute with the instance following the flags to compress or uncompress

```shell
$ ./huffman [targeting] [file]
```
#### file

* **file**: Path to the file

##### targeting

* **-c**: compress the file
* **-u**: uncompress the file

---

##### example:

```shell
$ ./huffman -c instances/t1.bin

--	  Encoding started

| File Compressed, File Location: instances/t1.bin.compr |

$ ./huffman -u instances/t1.bin.compr

 --	  Decoding started.

| File Uncompressed, File Location: instances/t1.bin.uncmp |
```
