# cronenberg #

[![Build Status](https://travis-ci.org/ess/cronenberg.svg?branch=master)](https://travis-ci.org/ess/cronenberg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ess/cronenberg)](https://goreportcard.com/report/github.com/ess/cronenberg)
[![Documentation](https://godoc.org/github.com/ess/cronenberg?status.svg)](http://godoc.org/github.com/ess/cronenberg)

It's like cron, but kinda twisted.

## Installation ##

You can either download the appropriate release from the [releases page](https://github.com/ess/cronenberg/releases) or install via `go get`:

```
go get -u github.com/ess/cronenberg/cmd/cronenberg
```

## Basics ##

Cronenberg works similarly to the standard `cron` that you already know and love, but with a few differences:

* It is not meant to be run as a system service
* It does not use a traditional crontab
* Jobs for a given `cronenberg` instance are loaded from a YAML file
* It supports locking jobs (jobs that can only ever have one running instance)
* It logs to STDOUT
* It supports both incoming (from the system) and job-scoped environment variables

In short, this is the `cron` implementation that you run along side your 12-factor application.

### Jobs File ###

There is not a hard-coded location for your jobs file. Instead, you pass the jobs file as the first argument to the `cronenberg` command like so:

```
cronenberg /path/to/my/jobs/file.yml
```

This is just a YAML file containing an array of Job objects. Here's an example jobs file:

```yaml
# Job objects require a name, a command, and a schedule (when). You can also
# specify a description, whether or not the job locks, and a hash of environment
# variables for the job.

# This is just a normal job that runs every minute
- name: what-am-i
  command: echo "I am a command"
  when: "* * * * *"

# This is a locking job that is scheduled to run every minute so long as there
# is not already an instance of the job in progress.
- name: picky-picky
  command: 'echo "I am a picky command" ; sleep 70'
  when: "* * * * *"
  lock: true

# This is a job that runs every five minutes and specifies an environment
# variable. If this variable is also set in cronenberg's executing shell, the
# value configured here takes precedence.
- name: know-your-environment
  command: echo $flibberty
  when: "*/5 * * * *"
  env:
    flibberty: gibbets
```
