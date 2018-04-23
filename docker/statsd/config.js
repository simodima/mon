(function() {
  return {
    port: parseInt(process.env.STATSD_PORT) || 8125,
    backends: ["./backends/console"],
    debug: true,
    flushInterval: 20000
  };
})()