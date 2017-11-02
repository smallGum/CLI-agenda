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

Please refer to [cmd-design](./cmd/cmd-design.md)

## Test

At first, there is only one public user `guest`, so let's create some users first

create user1

```shell
$ ./CLI-agenda register -u user1 -p user1
2017/11/02 08:30:21  created successfully
user:user1 created successfully!
```

create user2


```shell
$ ./CLI-agenda register --username=user2 --password=user2
2017/11/02 08:30:42  created successfully
user:user2 created successfully!

```

now list all users, guest user has no access to this.

```
$ ./CLI-agenda users
there are 3  users:
--------------------------
2017/11/02 08:30:49 only users loged in have access to this
```

so we have to login first

```
$ ./CLI-agenda login -u user1 -p user1
2017/11/02 08:31:08  log in successfully
user:user1 log in successfully
```

then

```
aaron@chengr25:~/work/src/github.com/jack-cheng/CLI-agenda$ ./CLI-agenda usersthere are 3  users:
--------------------------
user:user2
email:
tel:
--------------------------
user:guest
email:
tel:
--------------------------
user:user1
email:
tel:
--------------------------
```

user1 logout

```
$ ./CLI-agenda logout
2017/11/02 08:31:43  log out
now you are a guest
```
user2 login

```
$ ./CLI-agenda login --username=user2 --password=user2
2017/11/02 08:32:47  log in successfully
user:user2 log in successfully
```

set user2's telephone numbe and email
```
$ ./CLI-agenda setTel -t 15521165845
2017/11/02 08:33:33  set telephone to be 15521165845
```
```
$ ./CLI-agenda setEmail -e 123456@qq.com
2017/11/02 08:33:49  set email to be 123456@qq.com
```
and list users again to see result

```
$ ./CLI-agenda users
there are 3  users:
--------------------------
user:guest
email:
tel:
--------------------------
user:user1
email:
tel:
--------------------------
user:user2
email:123456@qq.com
tel:15521165845
--------------------------
```
cancel user2's account

```
$ ./CLI-agenda cancelUser
2017/11/02 08:34:17 cancel its account
```

and list users again to ensure user2 has been canceled
```
$ ./CLI-agenda users
there are 2  users:
--------------------------
2017/11/02 08:34:31 only users loged in have access to this
```
```
$ ./CLI-agenda login -u user1 -p user1
2017/11/02 08:34:45  log in successfully
user:user1 log in successfully
```
```
$ ./CLI-agenda users
there are 2  users:
--------------------------
user:guest
email:
tel:
--------------------------
user:user1
email:
tel:
--------------------------
```
