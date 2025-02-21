# Archiver
<h3 id>Archiving Things</h3>

The last tool we'll implement today is the Log Rotation tool. "Log rotation" is a process by which the old log file is archived and stored so that logs don't pile up in a single file indefinitely. It should work like this:

```
# Will create file /path/to/logs/some_application_1600785299.tag.gz
# where 1600785299 is a UNIX timestamp from `some_application.log`'s [MTIME](https://linuxize.com/post/linux-touch-command/)
~$ ./myRotate /path/to/logs/some_application.log
```

```
# Will create two tar.gz files with timestamps (one for each log) 
# and put them in /data/archive directory
~$ ./myRotate -a /data/archive /path/to/logs/some_application.log /path/to/logs/other_application.log
```
