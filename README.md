# CLI-Agenda

[![Build Status](https://travis-ci.org/smallGum/CLI-agenda.svg?branch=master)](https://travis-ci.org/smallGum/CLI-agenda)



## Our Team

15331050 Jiezhu Cheng

15331052 Rong Cheng

## Usage

```shell
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
  setEmail      set registered user's email
  setTel        set registered user's telephone number
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

when you want create a user with existed user name, you will get

```
./CLI-agenda register -u user1 -p 123456
2017/11/01 18:00:15 this username has been occupied
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

Now we test about meeting operations. First we create three users:

```
./CLI-agenda users
there are 4  users:
--------------------------
user:guest
email:
tel:
--------------------------
user:user1
email:
tel:
--------------------------
user:Alice
email:
tel:
--------------------------
user:Bob
email:
tel:
--------------------------
```

create two new meetings:

```
./CLI-agenda createMeeting -t "Season You Like" -p Alice+Bob -s 2017-07-12 -e 2017-07-15
2017/11/01 18:08:46 successfully create a new meeting "Season You Like".

./CLI-agenda logout
2017/11/01 18:18:41  log out
now you are a guest

./CLI-agenda login -u Alice -p 123456
2017/11/01 18:19:00  log in successfully
user:Alice log in successfully

./CLI-agenda createMeeting -t "How to Cook" -p user1 -s 2017-07-08 -e 2017-07-12
2017/11/01 18:19:50 successfully create a new meeting "How to Cook".
```

query one's meetings:

```
./CLI-agenda queryMeetings -s 2017-07-08 -e 2017-07-13
Alice's meetings between 2017-07-08 and 2017-07-13:

-------------------------------
title: How to Cook
participators: [user1]
start time: 2017-07-07
end time: 2017-07-12
sponsor: Alice
-------------------------------


-------------------------------
title: Season You Like
participators: [Alice Bob]
start time: 2017-07-12
end time: 2017-07-15
sponsor: user1
-------------------------------

2017/11/01 22:38:12 query the meetings between 2017-07-08 and 2017-07-13.
```

cancel one's meetings:

```
./CLI-agenda cancelMeeting -t "How to Cook"
cancel meeting How to Cook called!

./CLI-agenda queryMeetings -s 2017-07-08 -e 2017-07-13
Alice's meetings between 2017-07-08 and 2017-07-13:

-------------------------------
title: Season You Like
participators: [Alice Bob]
start time: 2017-07-12
end time: 2017-07-15
sponsor: user1
-------------------------------

2017/11/01 22:44:57 query the meetings between 2017-07-08 and 2017-07-13.
```

quit meeting:

```
./CLI-agenda quitMeeting -t "Season You Like"
quitMeeting Called

./CLI-agenda queryMeetings -s 2017-07-08 -e 2017-07-13
Alice's meetings between 2017-07-08 and 2017-07-13:
none.
2017/11/01 22:47:53 query the meetings between 2017-07-08 and 2017-07-13.
```
