#!/bin/sh

# Define eventlog and syslog-ng versions to be installed
EL_VER=0.2.13
SL_VER=3.6.4

# Exit staging process if file is missing
file_exists(){
    local FILE=$1
    if [ ! -e $FILE ]; then
        sleep 2s
        echo ""
        echo "ERROR: $FILE COULD NOT BE FOUND; EXITING STAGING PROCESS"
        echo "PLEASE TRY STAGING AGAIN"
        echo ""
        sleep 3s
        abort_staging
    fi
}

# In the case of a non-recoverable error, we will remove 'bash_start.sh' and
# exit setup. This will lead to staging failure which will save the user's time.
abort_staging(){
    rm -f /app/bash_start.sh
    exit 1
}

# Begin staging process
cd /app
echo "DOWNLOADING SYSLOG-NG AND EVENTLOG SOURCE..."
wget https://my.balabit.com/downloads/syslog-ng/open-source-edition/3.4.8/source/eventlog_$EL_VER.tar.gz
file_exists /app/eventlog_$EL_VER.tar.gz
wget https://my.balabit.com/downloads/syslog-ng/open-source-edition/3.6.4/source/syslog-ng_$SL_VER.tar.gz
file_exists /app/syslog-ng_$SL_VER.tar.gz
echo "DONE!"
echo "UNPACKING SOURCE..."
tar -xvzf eventlog_$EL_VER.tar.gz
tar -xvzf syslog-ng_$SL_VER.tar.gz
echo "DONE!"
echo "BUILDING EVENTLOG..."
cd /app/eventlog-$EL_VER
./configure --prefix=/app/eventlog/
make
make install
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/app/eventlog/lib/pkgconfig
echo "DONE!"
echo "BUILDING SYSLOG-NG..."
cd /app/syslog-ng-$SL_VER
./configure --prefix=/app/syslog-ng/
make
make install
echo "DONE!"
echo "INSTALLING SYSLOG-NG CONFIG..."
cp /app/syslog-ng.conf.template /app/syslog-ng/etc/syslog-ng.conf
echo "DONE!"
echo "CLEANING UP..."
rm -rf /app/eventlog_$EL_VER.tar.gz
rm -rf /app/syslog-ng_$SL_VER.tar.gz
rm -rf /app/eventlog-$EL_VER
rm -rf /app/syslog-ng-$SL_VER
echo "DONE!"
echo "SETUP COMPLETE!"
