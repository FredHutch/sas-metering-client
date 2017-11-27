# SAS Metering Client

Windows service to report the utilization of SAS on the desktop for license metering.

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
sc.exe description "sas-metering-client" "Sends SAS utilization metrics to the FredHutch metrics server"
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

promtheus server: yourserver.yourdomain.org 
prometheus port: someport-number 
username: someuser 
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


### Build Notes

**Dependencies**

Dependencies are vendored in the repo using the "godep" tool. If you add any additional dependencies, vendor them by running "godep save" and committing the dependencies.

**Binary information**

To include information and an icon in the compiled binary we are using "goversioninfo" (https://github.com/josephspurrier/goversioninfo). To use it, download in and istall it:

```
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
```

Edit the "versioninfo.json" file in this repo to include the required changes. Make sure that "//go:generate goversioninfo -icon=icon.ico" build instructions are at the top of the "main.go" source file. Run the following command in the root of the repo to generate/update the "resource.syso" file:

```
go generate
```

The code is now read to compile.

**Reducing binary size**

To reduce the binary size, build with the following flags:

```
go build -ldflags="-s -w"
```

**Hacks**

I had to hack the wmi.go source file vendored in this repo at "vendor\github.com\StackExchange\wmi\wmi.go" because it would file with a type error (can't assigned nil to string) is any of the processes properties were empty. The following hack fixes this issue (need to do an upstream PR):


```
// DefaultClient is the default Client and is used by Query, QueryNamespace
//var DefaultClient = &Client{}   <--- this was the default
//rmcdermo: the ^above^ line is the default, the line below replaces it, changes required
//so processes with nil valumes don't raise errors when loading into Win32_Process struct
var DefaultClient = &Client{NonePtrZero: true, PtrNil: false}  // <----- this is my hack
````

If you need to update the wmi packgage for any reason, you'll need to re-implement this hack if it's not fixed upstream yet.
