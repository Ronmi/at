AT runs some bash scripts at specific time.

This tool aims to provide shortcuts of something like `do_some_work.sh && wait_until_midnight.sh && do_some_other_work.sh`.

# Usage

## What to do

AT reads bash script commands from `stdin`, you can just redirect it or pipe some command `cat /some/bash/script.sh`.

## When to do

You can specify time ablosutely or relatively.

```sh
# absolute
echo ls | at 2006-01-02 15:04:05
echo ls | at 15:04:05

# relative, see https://golang.org/pkg/time/#ParseDuration for detailed format
echo ls | at 10s
```

# Accuracy

Since the user specific time is converted to relative format internally in AT, it will have few milliseconds delay (and even more on heavy-loading machine).

# License

GPLv3
