# CLI-Agenda

[![Build Status](https://travis-ci.org/smallGum/CLI-agenda.svg?branch=master)](https://travis-ci.org/smallGum/CLI-agenda)



## Our Team

15331050 Jiezhu Cheng

15331052 Rong Cheng

## Usage

```shell
$ ./CLI-agenda --help
CLI-agenda is cooperative program for meeting management using cobra package.
	It supports commands such as register, login, creatingMeeting, clearMeetings and so on.

Usage:
  CLI-agenda [command]

Available Commands:
  cancelMeeting cancel meetings you sponsored with specified title
  cancelUser    remove an account from users
  clearMeetings clear all meetings with you as sponsor
  createMeeting Create a new meeting
  help          Help about any command
  login         for guest to login
  logout        logout
  queryMeetings Query meetings of current login user between specific time interval
  quit          quit from all meetings with you as participator
  register      to register a new user
  users         list all users

Flags:
      --config string   config file (default is $HOME/.CLI-agenda.yaml)
  -h, --help            help for CLI-agenda
  -t, --toggle          Help message for toggle

Use "CLI-agenda [command] --help" for more information about a command.
```

## Commands Design

Please refer to 

[cmd-design]: ./cmd/cmd-design.md

