
# Node Express

Fun server that let's you spawn commands on the host VM.

## Setup

```
$ yarn
```

## Deploy

```
$ up
```

## Using

To list the root directory

```
$ curl `up url`/dir/var/log
total 76K
-rw------- 1 root root   0 Aug 18 20:09 btmp
-rw-r--r-- 1 root root 57K Aug 18 20:09 dracut.log
-rw------- 1 root root   0 Aug 18 20:09 tallylog
-rw-rw-r-- 1 root root   0 Aug 18 20:09 wtmp
-rw-rw-r-- 1 root root 16K Aug 18 20:10 yum.log
```

To execute the command `whomai`

```
$ curl `up url`/cmd/whoami
sbx_user1059

``` 

To execute the command with arguments

````
$ curl `up url`/cmd/ls?args=-al

total 11044
drwxr-xr-x  3 slicer  497     4096 Sep  1 05:26 .
drwxr-xr-x 22 root   root     4096 Sep  1 05:01 ..
-rw-r--r--  1 slicer  497     1200 Sep  1 05:22 app.js
-rwxr-xr-x  1 slicer  497     4710 Sep  1 05:25 byline.js
-rwxr-xr-x  1 slicer  497 11269647 Sep  1 05:25 main
drwxrwxr-x 45 slicer  497     4096 Sep  1 05:26 node_modules
-rw-r--r--  1 slicer  497       91 Sep  1 04:32 package.json
-rwxr-xr-x  1 slicer  497      931 Sep  1 05:25 _proxy.js
-rw-r--r--  1 slicer  497      145 Sep  1 04:36 up.json
```


