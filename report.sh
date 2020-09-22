#!/bin/bash
exit_script(){
    exit 1
}

touch response.log
mydate=$(date +"%y%m%d%a")
report="REPORT "
fail="WARNING!!!"
theme=$report$mydate
date>>./response.log
./HealthReport>>./response.log

exit_script
EOF