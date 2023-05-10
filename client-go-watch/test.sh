#!/bin/bash

readonly OK=0
readonly NONOK=1
readonly UNKNOWN=2

readonly DockerDisk='docker disk'

# Check systemd cmd present
if ! command -v du >/dev/null; then
  echo "Could not find cmd du"
  exit $UNKNOWN
fi

# Return success if service active (i.e. running)
disk=`df -hT / /var/lib/docker /grlocaldata /var/lib/etcd /data 2>/dev/null | awk '{print $6,$7}' | sed -n '1!p' | sort -u | sed 's/%//g'`

while read line; do
  a_line=($line)
  if [ ${a_line[0]} -gt 80 ]; then
    echo "节点$i上${a_line[1]}分区超过80%, 请检查! "
    exit $NONOK
  fi
done


exit $OK