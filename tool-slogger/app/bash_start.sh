#!/bin/sh

SV="0.2"

echo "###############################################################"
echo "##                   SLOGGER VERSION: $SV                    ##"
echo "###############################################################"
echo "##                   START SYSLOG SERVICE                    ##"
echo "###############################################################"
sudo /app/init/syslog-ng start
/app/syslog-ng/sbin/syslog-ng -V
echo "###############################################################"
echo "##     RUN NETSTAT TO ENSURE SYSLOG IS LISTENING ON 1999     ##"
echo "###############################################################"
netstat -l

tail -f /app/logs/all-remote-logs
