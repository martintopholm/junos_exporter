# JUNOS_EXPORTER 8 "NOVEMBER 2020" "Prometheus helper" "System Manager's Manual"

## NAME

junos_exporter - prometheus junos exporter via ssh

## SYNOPSIS

`junos_exporter` [`-web.listen-address` *address:port*] [`-config.file` *file*] ...

## DESCRIPTION

`junos_exporter` collects metrics from Juniper devices via ssh. The set of
metrics is minimal to increase performance. This is an alternative to
snmp_exporter that is very flexible but also very complex and may put a higher
load on the device.


## OPTIONS

`-accounting.enabled`
  Scrape accounting flow metrics

`-alarms.filter string`
  Regex to filter for alerts to ignore

`-bgp.enabled`
  Scrape BGP metrics (default true)

`-config.file string`
  Path to config file

`-config.ignore-targets`
  Ignore check if target is specified in config

`-debug`
  Show verbose debug output in log

`-dynamic-interface-labels`
  Parse interface descriptions to get labels dynamicly (default true)

`-environment.enabled`
  Scrape environment metrics (default true)

`-firewall.enabled`
  Scrape Firewall count metrics (default true)

`-fpc.enabled`
  Scrape line card metrics (default true)

`-ifdiag.enabled`
  Scrape optical interface diagnostic metrics (default true)

`-interfaces.enabled`
  Scrape interface metrics (default true)

`-ipsec.enabled`
  Scrape IPSec metrics

`-isis.enabled`
  Scrape ISIS metrics

`-l2circuit.enabled`
  Scrape l2circuit metrics

`-ldp.enabled`
  Scrape ldp metrics (default true)

`-nat.enabled`
  Scrape NAT metrics

`-ospf.enabled`
  Scrape OSPFv3 metrics (default true)

`-queues.enabled`
  Scrape interface queue metrics

`-routes.enabled`
  Scrape routing table metrics (default true)

`-routingengine.enabled`
  Scrape Routing Engine metrics (default true)

`-ssh.keep-alive-interval duration`
  Duration to wait between keep alive messages (default 10s)

`-ssh.keep-alive-timeout duration`
  Duration to wait for keep alive message response (default 15s)

`-ssh.keyfile string`
  Public key file to use when connecting to junos devices using ssh

`-ssh.password string`
  Password to use when connecting to junos devices using ssh

`-ssh.reconnect-interval duration`
  Duration to wait before reconnecting to a device after connection got lost (default 30s)

`-ssh.targets string`
  Hosts to scrape

`-ssh.user string`
  Username to use when connecting to junos devices using ssh (default "junos_exporter")

`-storage.enabled`
  Scrape system storage metrics (default true)

`-version`
  Print version information.

`-web.listen-address string`
  Address on which to expose metrics and web interface. (default ":9326")

`-web.telemetry-path string`
  Path under which to expose metrics. (default "/metrics")

## AUTHOR

Daniel Czerwonk

## SEE ALSO

prometheus(1), [Prometheus](http://prometheus.io/)
