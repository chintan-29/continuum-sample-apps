# SLOGGER

#### Version 0.2

Slogger is a simple syslog server that is built via the bash staging process.

Source code for `eventlog` and `syslog-ng` is downloaded and compiled into the
`/app` directory during staging. Afterwards, the completed slogger package is
uploaded.

### The Slogger startup process:
1. Start `syslog-ng` service
2. Run `netstat -l` to ensure the service is listening on port `1999`
3. `tail -f` is run against `/app/logs/all-remote-logs`

Slogger listens on port `1999` by default. All log messages that hit this port
will be appended to the `/app/logs/all-remote-logs` file. Due to the startup
process ending with tailing this file, all log messages can be viewed in
realtime by running `apc job logs slogger`.

By adding a TCP route to the Slogger job and then adding an app drain to other
apps, this allows the collection of job logs via Slogger.

### Instructions for building Slogger and setting it as an app drain target:

1. Change working directory to Slogger root (this directory)
2. `apc app create`
3. Wait for staging process to complete and the app to start
4. Add a TCP route with `apc route add auto --tcp --app slogger -p 1999`
5. Note the TCP route value once created
6. Add an app drain to a second app with
   `apc drain add syslog://<HOST-IP>:<HOST-PORT> --app <OTHER-APP-NAME>`

For example, if my TCP route is `tcp://54.213.204.135:50154` and my other app
is called `go-app`, I would issue the command
`apc drain add syslog://54.213.204.135:50154 --app go-app`

At this point, logs that are generated by `go-app` should be viewable by
running `apc job logs slogger`.
