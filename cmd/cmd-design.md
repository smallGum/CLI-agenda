# Commands Design of CLI-agenda

## register
Command `register` is used to create a new user, unique username and free password is required
```
$ register -u chengr25 -p 123456
```

where

+ `-u` should be followed by **username** of user to be created, it also works by entering
`--username=[username]`
+ `-p` should be followed by **password** of user to be created, it also works by entering
`--password=[password]`

## login
Command `login` is used to login as a registered user

```
$ login -u chengr25 -p 123456
```
where

+ `-u` should be followed by **username** of user to login, it also works by entering
`--username=[username]`
+ `-p` should be followed by **password** of user to login, it also works by entering
`--password=[password]`

## logout

Command `logout` allows the current user to log out and switch into default user `guest` :

```shell
$ logout
```

## users

Command `users` is used to list all registered users:

```shell
$ users
```

## cancelUser

Command `cancelUser` removes current user account and meetings sponsored by him. And also, removes him from meetings he has participated in, then the current user of agenda will be switched to `guest`:

```shell
$ cancelUser
```

## setTel
Command `setTel` is used to set registered user's telephone number, guest users have no access to this.
```
$ setTel -t 15521122735
```
where

+ `-t` means **telephone number**, it also works by entering `--telephone=12166136151`

## setEmail

Command `setEmail` is used to set registered user's email address, guest has no access to this.
```
$ setEmail -e 38437262@qq.com
```

where

+ `-e` means **email**, it also works by entering `--email=38437262@qq.com`

## createMeeting

Command `createMeeting` creates a new meeting:

```shell
$ createMeeting -t "MeetingTitle" -p User1+User2+...+Usern -s 2017-11-12 -e 2017-11-15
```

where

+ `-t` means **title** of meeting, and it also works by entering

   `--title="MeetingTitle"` or `--title "MeetingTitle"` .

+ `-p` means **participators** of meeting, and it also works by entering

  `--participators=User1+...+Usern` or `--participators User1+...+Usern` .

+ `-s` means **start time** of meeting, and it also works by entering

  `--starttime=2017-11-12` or `--starttime 2017-11-12` .

+ `-e` means **end time** of meeting, and it also works by entering

  `--endtime=2017-11-15` or `--endtime 2017-11-15` .

## queryMeetings

Command `queryMeetings` shows details of meetings sponsored or participated by current login user within specific time interval:

```shell
$ queryMeetings -s 2017-11-12 -e 2017-11-15
```

where

+ `-s` means **start time** of the interval, and it also works by entering

  `--starttime=2017-11-12` or `--starttime 2017-11-12` .

+ `-e` means **end time** of the interval, and it also works by entering

  `--endtime=2017-11-15` or `--endtime 2017-11-15` .

## quitMeeting

Command `quitMeeting` allows the current user to quit the meeting he has participated in:

```shell
quitMeeting -t "MeetingTitle"
```

where

- `-t` means **title** of meeting, and it also works by entering

  `--title="MeetingTitle"` or `--title "MeetingTitle"` .

## clearMeetings

Command `clearMeetings` allows the current user to clear all the meetings he sponsors:

```shell
$ clearMeetings
```
