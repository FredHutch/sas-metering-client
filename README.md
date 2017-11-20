# SAS Metering Client

Windows service to report the utilization of SAS to a Prometheus push gateway to license metering.

### Requirements

- Microsoft Windows 7+
- Administrative privileges to install it as a service

### Installation

- Download "binary/sas-metering-client.exe" from this repo to your Windows system ("C:\Program Files\sas-metering-client\").
- Open a command-line shell with administrative permissions and run the following commands:


```
sas-metering-client.exe install
sas-metering-client.exe start
sc.exe config "sas-metering-client" start=auto
sc.exe failure "sas-metering-client"  actions=restart/600000/restart/600000/restart/41400000 reset=86400
```

### Checking the verion and configuration

To see what version of the agent you are running run the following command:

```
sas-metering-client.exe version

version: 0.1
```

To see the configuration that the agent is using run the following command:

```
sas-metering-client.exe print-config

promtheus server: prometheus.fredhutch.org
prometheus port: 9991
username: promuser
password: '*******fi$'
TLS: true
```

### Uninstalling

To remove the SAS metering client service from your system run the following commands (assuming you installed it to "C\Program Files\sas-metering-client"):

```
C:\Program Files\sas-metering-client\sas-metering-client.exe stop
C:\Program Files\sas-metering-client\sas-metering-client.exe remove
```

After the service is stopped and removed, simply delete the binary from your system.


### SAS Simulator

SAS is not free and obtaining a trial version for testing is not trivial, so to test this monitoring agent, I’ve written a SAS simulator. 

To run it just download "sas-simulator/binary/sas.exe" from this repo to your windows computer and run it via the command-line of double-click in Windows explorer.
