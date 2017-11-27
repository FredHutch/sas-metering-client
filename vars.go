package main

// Version of the agent
const version = "0.1"

// Name of the Windows Service
const svcName = "sas-metering-client"

// Process to monitor
const procName = "sas.exe"

// Prometheus PushGateway Server FQDN and TCP port
const promServer = "yourserver.name.here"
const promServerPort = "some-port-number"

// Prometheus PushGateway Username and Password
const promUser = "some-username "
const promPass = "some-password"

// How often (seconds) to send metrics to the server
const reportingInterval = 60
