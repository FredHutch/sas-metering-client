package main

// Name of the Windows Service
const svcName = "sas-metering-client"

// Process to monitor
const procName = "sas.exe"

// Prometheus PushGateway Server FQDN and TCP port
const promServer = "prometheus.fhcrc.org"
const promServerPort = "9991"

// Prometheus PushGateway Username and Password
const promUser = "user"
const promPass = "pass"

// How often (seconds) to send metrics to the server
const reportingInterval = 60
