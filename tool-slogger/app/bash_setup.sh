#!/bin/sh

cd /app
echo "DOWNLOADING SYSLOG SOURCE!"
wget https://my.balabit.com/downloads/syslog-ng/open-source-edition/3.4.8/source/eventlog_0.2.13.tar.gz
wget https://my.balabit.com/downloads/syslog-ng/open-source-edition/3.6.4/source/syslog-ng_3.6.4.tar.gz
echo "UNPACKING SOURCE!"
tar -xvzf eventlog_0.2.13.tar.gz
tar -xvzf syslog-ng_3.6.4.tar.gz
cd /app/eventlog-0.2.13
echo "BUILDING EVENTLOG!"
./configure --prefix=/app/eventlog/
sudo make
sudo make install
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/app/eventlog/lib/pkgconfig
echo "BUILDING SYSLOG-NG!"
cd /app/syslog-ng-3.6.4
./configure --prefix=/app/syslog-ng/
sudo make
sudo make install
echo "INSTALLING SYSLOG-NG CONFIG!"
sudo cp /app/syslog-ng.conf.template /app/syslog-ng/etc/syslog-ng.conf
echo "CLEANUP!"
sudo rm -rf /app/eventlog_0.2.13.tar.gz
sudo rm -rf /app/syslog-ng_3.6.4.tar.gz
sudo rm -rf /app/eventlog-0.2.13/
sudo rm -rf /app/syslog-ng-3.6.4/
echo "SETUP COMPLETE!"
