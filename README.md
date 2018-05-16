<p align="center"><img src="doc/images/mon.png" width="360"></p>

# MON (the process monitor)

Process monitor is a simple command to wrap your application and monitor it. That's it!
Whether you're running a cron command, a long-running process or the entire application you can prepend `mon` to the execution command to monitor the CPU, memory and network usage.

# Supported Backends

The only supported backend is [statsd][statsd]

#### StatsD

To send metrics to a statsd backend you can setup the following env variables

```bash
export MON_STATSD=true
export MON_STATSDHOST=localhost
export MON_STATSDPORT=8125
export MON_PROCESSNAME=command_name

mon dd if=/dev/zero of=/dev/null
```

Your statsd server will receive the following metrics 
- `command_name.net.byte_sent`
- `command_name.net.byte_recv`
- `command_name.cpu`
- `command_name.memory`
- `command_name.error`

# How to use it 

```
~/ $ MON_PROCESSNAME=infinite-consumer MON_TICK=5 mon my-queue-consumer -m 200 -p 1   
     ────────────────────────────────────┬─────── ─┬─ ────────┬──────── ─────┬─────   
                                         │         │          │              │        
mon configuration (optional) ────────────┘         │          │              │        
mon ───────────────────────────────────────────────┘          │              │        
your command ─────────────────────────────────────────────────┘              │        
your command options ────────────────────────────────────────────────────────┘        
``` 

# Configuration 

We provide defaults for every configuration, but if you need to customize metrics here the list of env variables to provide:

- `MON_STATSD`: enable\disable statsd metrics, default disabled (example: `MON_STATSD=true`)
- `MON_STATSDHOST`: the hostname of statsd daemon (example: `MON_STATSDHOST=localhost`)
- `MON_STATSDPORT`: the post on which statsd is listening (example: `MON_STATSDPORT=8125`)
- `MON_PROCESSNAME`: the name of the process which will identify the monitored command (example: `MON_PROCESSNAME=command_name`)
- `MON_TICK`: the sampling interval in seconds (example: `MON_TICK=5`)

# Development

Initialize the container env

`bin/init.sh`

Install dependencies

`bin/provision.sh`

Run a one-off command to the go container

`bin/run.sh`

##### Examples
`bin/run.sh go run main.go sleep 2`

`bin/run.sh go run main.go dd if=/dev/zero of=/dev/null`

`bin/run.sh go run main.go exit 1`

##### Build the binary

`bin/build.sh`


[statsd]: https://github.com/etsy/statsd
