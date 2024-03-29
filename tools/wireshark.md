From https://www.wireshark.org/docs/wsug_html_chunked/index.html 

Currently up to 5.0 (not yet read) (https://www.wireshark.org/docs/wsug_html_chunked/ChapterIO.html)

# Overview

Wireshark is a network packet analyzer – think of it as a measuring device for examining what's happening inside a network cable (like an electrician uses a voltmeter for an electric cable). It has many practical applications, of course, but it can even be used to learn the internals of various network protocols. Wireshark doesn't send packets or do anything active, it is purely a measurement tool. If run from a useful point in the network, it can even capture network traffic not intended for your local machine.

Wireshark can be augmented to work with new unsupported (or, theoretically, custom) protocols via plugins, or even by directly modifying the source. Wireshark isn’t limited to just network interfaces—on most systems you can also capture USB, Bluetooth, and other types of packets.

# GUI

Packet list and detail navigation can be done entirely from the keyboard (see https://www.wireshark.org/docs/wsug_html_chunked/ChUseMainWindowSection.html#ChUseTabNav).

### The Packet List Pane

Each line in the packet list corresponds to *one packet*. The details of the selected packet are displayed in the Packet Details and Packet Bytes panes below.

Symbols in the first column describe how packets in the list are related to each other:

![[Wireshark Related Packet Indicators.png]]

### The Packet Details Pane

Information enclosed in square brackets is not actually present in the captured data--this data is generated by Wireshark itself.

### The Packet Bytes Pane
Each line contains sixteen hexadecimal bytes and then sixteen ASCII bytes. Non-printable bytes are replaced with a period (".")

# Capturing Live Network Data

### Setup
Appropriate priveleges are necessary to capture network data (see https://gitlab.com/wireshark/wireshark/-/wikis/CaptureSetup/CapturePrivileges).

### Capturing Your Own Local Traffic
Simply double click the network interface with which you're concerned and traffic should be displayed (if there is any).

##### Loopback Capture Setup
When you are trying to capture traffic from your machine to itself, that traffic will not be sent over a real network interface, *even if it's being sent to an address on one of the machine's real network interfaces*. That means that even if that traffic you are interested in is destined for one of your machine's network adapters, you won't see it by selecting the destination interface as the interface on which to capture, but only on the "Adapter for loopback traffic capture" interface. See https://gitlab.com/wireshark/wireshark/-/wikis/CaptureSetup/Loopback.

### Filtering
Filters are written in libpcap filter language (see https://www.tcpdump.org/manpages/pcap-filter.7.html for a full reference). A capture filter takes the form of a series of primitive expressions connected by conjunctions *(and/or)* and optionally preceded by *not*:

`[not] primitive [and|or [not] primitive ...]`

i.e.

`tcp port 23 and host 10.0.0.5`

or

`tcp port 23 and not src host 10.0.0.5`

##### Primitives

```
[src|dst] host <host>
```
Allows you to filter on a host IP address or name. Preceding with `src` or `dst` is optional, and allows you to specify that you're only interested traffic where the specified `host` is the source or the destination.

---

```
ether [src|dst] host <ehost>
```
Like the above primitive, but allows you to filter on Ethernet host addresses only.

---

```
gateway host <host>
```
Allows you to filter on packets that used *host* as a gateway (that is, where the Ethernet source or destination was `host`, but neither the source nor destination IP address was `host`.

---

```
[src|dst] net <net> [{mask <mask>}|{len <len>}]
```

Allows you to filter on network numbers. In addition, you can specify either the netmask or the CIDR prefix for the network if they are different from your own.

---

```
[tcp|udp] [src|dst] port <port>
```

Allows you to filter on TCP and UDP port numbers.

---

```
less|greater <length>
```

Allows you to filter on packets whose length was less-than-equal or greater-than-equal to the specified `length`.

---

```
ip|ether proto <protocol>
```

Allows you to filter on the specified `protocol` at either the Ethernet layer or the IP layer.

---

```
ether|ip broadcast|multicast
```

Allows you to filter on either Ethernet or IP `broadcast`s or `multicast`s

---

```
<expr> relop <expr>
```

Allows you to create complex filter expressions that select specific bytes or ranges of bytes within packets. See https://www.tcpdump.org/manpages/pcap-filter.7.html.

### Automatic Remote-Traffic Filtering
If Wireshark is running remotely (i.e. using SSH or RDP), it attempts to automatically filter out the packets related to the remote connection.