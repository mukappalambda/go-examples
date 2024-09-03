# Task Aborted on Timeout

```bash
$ ./task-aborted-on-timeout -help
Usage of ./task-aborted-on-timeout:
  -dt duration
        duration (default 10ms)
  -timeout duration
        timeout (default 5ms)
$ ./task-aborted-on-timeout -dt 19ms -timeout 20ms
Task execution time: 19ms; timeout: 20ms
Task completed.
$ ./task-aborted-on-timeout -dt 20ms -timeout 20ms
Task execution time: 20ms; timeout: 20ms
Task completed.
$ ./task-aborted-on-timeout -dt 21ms -timeout 20ms
Task execution time: 21ms; timeout: 20ms
Timeout is reached. Task aborted.
```
