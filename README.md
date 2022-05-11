# zsearch
Search for file names or matching file contents according to the specified directory

# download
After downloading, configure the system environment variables and you can use them directly  
The download address is: https://github.com/zztroot/zsearch/releases

# example
### Search file name and file content (all)
```cmd
$ zsearch -d error
[f] zsearch/error.go
[c] .idea/workspace.xml -> 2
[c] zsearch/search.go -> 3

StartTime:2022-05-11 10:30:13 | EndTime:2022-05-11 10:30:13 | FileNum:13 | DirNum:4 | Second:0ms
```
-d：specify search content  
[f] search by file name [type]  
[c] search by file content [type]  
*more usage -h
### Search by file name
```cmd
$ zsearch -f -d error
[f] zsearch/error.go

StartTime:2022-05-11 10:35:24 | EndTime:2022-05-11 10:35:24 | FileNum:13 | DirNum:4 | Second:0ms
```
-f：search by file name  
*more usage -h
### Search by file content 
```cmd
$ zsearch -c -d error
[c] zsearch/search.go -> 3
[c] .idea/workspace.xml -> 2

StartTime:2022-05-11 10:38:41 | EndTime:2022-05-11 10:38:41 | FileNum:13 | DirNum:4 | Second:0ms
```
-c：search by file content  
*more usage -h
##### -d parameter is a required item, and other parameters are optional. The default is to search the file name and file content of the current directory
##### -p you can specify a search directory