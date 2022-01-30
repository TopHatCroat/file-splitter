# File splitter

Small tool for splitting files found in a path into multiple groups. Usefull for parallelisation of whatever can be 
paralleled with multiple files.

### Options

```text
  -glob string
        Glob pattern for files to split (default "*")
  -index int
        Index of current split
  -total int
        Total number of splits (default 2)
```

### Examples
When running in this repository it gives the following output 
```bash
file-splitter -glob "*" -index=0 -total=2
# .gitignore README.md file-splitter.iml go.sum main_test.go

file-splitter -glob "*" -index=1 -total=2
# LICENSE file-splitter go.mod main.go
```