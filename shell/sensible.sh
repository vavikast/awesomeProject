#!/bin/bash
#Author: Felix
#Date: `date +"%F %T"`
#DESCRIPTION
#敏感信息文件和存放位置
tempfile=`mktemp`
echo $tempfile
read -p "Please 输入敏感信息文件夹位置，例如 /root/mingan.txt："  sensible
for FILE in $(find / \( -path "/mnt" -o -path "/media" -o -path "/etc" -o -path "/dev" -o -path "/boot" -o -path "/bin" -o -path "/sys" -o -path "/proc" -o -path "/sys" -o -path "/lib" -o -path "/lib64"  -o -path "/run" \) -prune -o -type f 2>/dev/null ); do
    sensibletxt=`cat $sensible`
	for TXT in $sensibletxt;do
		num=$(grep -c $TXT $FILE 2>/dev/null)
		echo $num
		if [ $num  -gt 2 ];then
		echo "$TXT-->$(grep -c $TXT $FILE)-->$FILE" > $tempfile
    fi
	done
done