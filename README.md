Timestamp - a timestamp conversion CLI utility
==============================================

Yes, I know I can get around with `date` but this is easier to use.


Install
-------

```shell
$ make install
```


Flags
-----

- `-t`: Show time (default: false). If doubled up, also convert to local and show TZ (default: UTC)
- `-v`: Show version
- `-h`: Show help


Usage
-----

```shell
$ timestamp 1653881338
2022-05-30

$ timestamp 1629484202017
2021-08-20

$ timestamp -t 1653881338
2022-05-30T03:28:58

$ timestamp -tt 1653881338
2022-05-30T05:28:58Z+02:00
```
