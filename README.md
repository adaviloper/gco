# gco

## Installation

```sh
go get github.com/adaviloper/gco@latest
```

## Description

Git utility to simplify creating and switching branches based on a provided pattern matching.

### Branch switching

```sh
gco some-full-branch-name
gco <ticket-number>
gco <partial-string-match>
```

### Branch Creation
```sh
gco -b <ticket-id> some arbitrary length ticket description
gco -e <ticket-id> some arbitrary length ticket description
gco -s <ticket-id> some arbitrary length ticket description
gco -p <ticket-id> some arbitrary length ticket description
gco -t <ticket-id> some arbitrary length ticket description
```

## Configuration

```yaml
repositories:
  my_repo:
    ticket_prefix: MY_PREFIX
    bug: bug_type // optional
    epic: epic_type // optional
    story: story_type // optional
    spike: spike_type // optional
    task: task_type // optional
```

## Flags


| Flag | Name | Description |
|--- | --- | --- |
| -r | remote | Include parsing through remote branches |
| -b | bug    | Create a bug ticket branch |
| -e | epic   | Create a epic ticket branch |
| -s | story  | Create a story ticket branch |
| -p | spike  | Create a spike ticket branch |
| -t | task   | Create a task ticket branch |
---

