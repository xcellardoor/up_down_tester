[![Stories in Ready](https://badge.waffle.io/xcellardoor/up_down_tester.png?label=ready&title=Ready)](https://waffle.io/xcellardoor/up_down_tester?utm_source=badge)
# Up-Down Tester

Build Status Be Like - [![Build Status](https://travis-ci.org/xcellardoor/up_down_tester.svg?branch=master)](https://travis-ci.org/xcellardoor/up_down_tester)

This program is meant to be run by systems that lie within a person VPN or Intranet and to verify whether systems that should be available or operational are indeed so.

Since I run a lot of VPN-only services myself and a few small websites for friends, I wanted a one-click-style solution I could use to validate whether things were working or if I needed to start running off to fix things.

*WIll insert usage instructions here*


This code is part of my effort to better myself in Golang and coding in general, as well as Git management.

##Permissions
Some Linux distros are pre-configured to treat ICMP transmissions as a privileged function. If you wish to use PING as a testing mechanism you can either:
* Run this program as root (do you trust me?!)
* Reconfigure your machine to allow regular users, or members of a group, to ping (GIYF)
