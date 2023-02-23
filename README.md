# network

### host_ip.go

It gets an IP address of a network adaptor attached to the local host. 
There will be either multiple IP addresses for a network adaptor or multiple network adpators to select the IP address from. 

In order to get an correct IP address among multiple addresses and multiple network adaptors, you will need to provide an IP address range that the IP address you want will be in.

The format of the IP address range will be for example "192.168.56.1/24" for IPv4, and ::ffff:192.2.16.10/96 for IPv6.
  
