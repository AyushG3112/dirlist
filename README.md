# DirList

DirList allows you to quickly spin up an HTTP server to view and browse directory listing of any of your folders, and browse them on your browser. 

## Installing

```
go install github.com/ayushg3112/dirlist/cmd/dirlist@latest
```

## Usages

### View Help

```
$ dirlist --help 
Usage: 
      --help                View help
  -p, --port int            Port on which to start the listing server (default 8000)
  -d, --root-dir string     Root directory to start directory listing
  -f, --sort-field string   Field to sort by (default "modifiedAt")
      --sort-order string   Sorting order. ASC/DESC. (default "ASC")
```

### Start a Server

```
$ dirlist --root-dir "/Users/ayushgupta/Downloads" --sort-order ASC --sort-field modifiedAt
2022/09/03 16:15:24 starting the server at port 8000
```
