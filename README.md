skpr-crond
==========

A lightweight crond for local development environments. Also leveraged the Skpr Preview environments. 

## What problem are we solving?

The standard Linux cron daemon is had to run in containers as it was built to run as root and then
handle subprocesses as a specific user.

Ultimately this is alot to maintain when all developers want is a local developer approach for routinely
executing a command to simulate cron tasks running in production.

## Usage

**Command**

```bash
$ skpr-crond --config=path/to/config/file.yaml
```

**Configuration**

```yaml
example:
  schedule: "* * * * *"
  command: echo 1
```
