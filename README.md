# SAS Metering Client

Windows service to report the utilization of SAS to a Prometheus push gateway to license metering.

### Requirements

- Microsoft Windows 7+
- Administrative privileges to install it as a service

### Installation

- Download "binary/sas-metering-client.exe" from this repo to your Windows system.
- Open a command-line shell with administrative permissions

Run the follow commands:

```
sas-metering-client.exe install
sas-metering-client.exe start
```

### SAS Simulator

SAS is not free and obtaining a trial version for testing is not trivial, so to test this monitoring agent, I’ve written a SAS simulator. 

To run it just download "sas-simulator/binary/sas.exe" from this repo to your windows computer and run it via the command-line of double-click in Windows explorer.
