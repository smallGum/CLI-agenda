# Commands Design of CLI-agenda

## logout

Command `logout` allows the current user to log out and switch into default user `guest` :

```shell
$ logout
```

## users

Command `users` lists all registered users:

```shell
$ users
```

## cancelUser

Command `cancelUser` removes current user account and meetings sponsored by him. And also, removes him from meetings he has participated in:

```shell
$ cancelUser
```

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