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
gco --ticket <ticket-type> <ticket-id> <some-arbitrary-length-ticket-description>
```

## Configuration

```yaml
repositories:
  default:
    ticket_prefix: MY_PREFIX
    bug: bugfix // optional
    // add whatever additional mappings you'd like for whatever <ticket-type> value you pass in
separator: "/"
```

## Flags


| Flag | Name | Description |
|--- | --- | --- |
| -r       | remote | Include parsing through remote branches |
| --ticket | ticket   | Create a ticket branch with the mapped prefix |
---

