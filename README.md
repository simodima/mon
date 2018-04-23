# MON (the process monitor)

Process monitor is a simple command to wrap your application an monitor it. That's it!
Whether you're running a cron command, a long-running process or the entire application you can prepend the `pm` to the execution to monitor the CPU and Memory

# How to

```
    ~/ $ MON_PROCESSNAME=infinite-consumer MON_TICK=5 mon  my-queue-consumer  -m 200 -p 1
         ____________________________________________ ___  _________________  ___________ 
         |                                              |           |               |
    mon _|                                              |           |               |
    configuration (optional)                            |           |               | 
                                                        |           |               |
    mon ________________________________________________|           |               |
                                                                    |               |
                                                                    |               |
    your command  __________________________________________________|               |
                                                                                    |
    your command  __________________________________________________________________|
    options
``` 

# Configuration 

We provide sensible defaults for every configuration, but if you need to customize metrics here the list of env variables to provide:

#### StatsD host & port configuration

MON_STATSDHOST=localhost

MON_STATSDPORT=8125

#### Process name to customize metrics

MON_PROCESSNAME=command_name

(you will receive `command_name.error`, `command_name.cpu`, `command_name.memory` metrics inside statsd)


#### The interval for data sampling (secods)

MON_TICK=5

# Development

Initialize the container env: `bin/init.sh`

Install dependencies: `bin/provision.sh`

Run a one-off command to the go container: `bin/run.sh`

##### Examples
`bin/run.sh go run main.go sleep 2`

`bin/run.sh go run main.go dd if=/dev/zero of=/dev/null`

`bin/run.sh go run main.go exit 1`

##### Build the dist

`bin/build.sh`