# Informer

## What

Returns information at your will. Should enable you to turn off all notifications, without loosing important notices.

Format should be `yaml`. *Should* is used as there are no `tests`.

### Github

Queries for `user`, `repo` and `kind`. By default the search is specified to type `open` and `sort:updated-desc` (showing the newest first).

In case you want to `query` something completely differnt, use `query`.

Lets pretend we name the `BIN=`[`igith`](https://www.dict.cc/?s=Igitt%21):

```bash
Usage of igith:
  -q string
        query by hand, if set ignores all other flags
  -r string
        filter to given repository
  -t string
        type to search for (default "pr")
  -u string
        user name to check (default "ibihim")
  -v    prints out additional data
```

## Why

We receive a lot of notifications. Those notifications are counter on apps, notifications, sounds, alerts, emails and what not.

It is too much.

Jobs that demand a high level of focus are in ever bigger demand.

This is a huge conflict.

## Example Usage

```bash
# Check for updates on co-workers aka dream team.
for user in "s-urbaniak" "stlaz" "slaskawi" "ibihim" "emilym1"; do
    igith -u "$user"
done
```
