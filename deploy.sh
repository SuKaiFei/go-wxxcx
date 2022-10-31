#!/bin/bash -v -e

PORT=0
#判断当前端口是否被占用，没被占用返回0，反之1
function Listening {
   TCPListeningnum=`netstat -an | grep ":$1 " | awk '$1 == "tcp" && $NF == "LISTEN" {print $0}' | wc -l`
   UDPListeningnum=`netstat -an | grep ":$1 " | awk '$1 == "udp" && $NF == "0.0.0.0:*" {print $0}' | wc -l`
   (( Listeningnum = TCPListeningnum + UDPListeningnum ))
   if [ $Listeningnum == 0 ]; then
       echo "0"
   else
       echo "1"
   fi
}

#指定区间随机数
function random_range {
   shuf -i $1-$2 -n1
}

#得到随机端口
function get_random_port {
   templ=0
   while [ $PORT == 0 ]; do
       temp1=`random_range $1 $2`
       if [ `Listening $temp1` == 0 ] ; then
              PORT=$temp1
       fi
   done
   echo "port=$PORT"
}
get_random_port 8000 8010; #这里指定了1~10000区间，从中任取一个未占用端口号

nextPort=${1:-$PORT}
curPort=${2:-$(cat /app/wxxcx/pid)}

echo 当前运行端口：$curPort
echo 下一次部署后的端口：$nextPort

sed -i "s/$curPort/$nextPort/" /app/wxxcx/config.yaml

nohup /app/wxxcx/go-wxxcx -conf /app/wxxcx/config.yaml >/app/wxxcx/go-wxxcx$(date '+%Y%m%d')-$nextPort.log 2>&1 &

#设置变量，url为你需要检测的目标网站的网址（IP或域名）
url=http://localhost:$nextPort/wxxcx/bqb/ping
 
#定义函数check_http：
#使用curl命令检查http服务器的状态
#-m设置curl不管访问成功或失败，最大消耗的时间为5秒，5秒连接服务为相应则视为无法连接
#-s设置静默连接，不显示连接时的连接速度、时间消耗等信息
#-o将curl下载的页面内容导出到/dev/null(默认会在屏幕显示页面内容)
#-w设置curl命令需要显示的内容%{http_code}，指定curl返回服务器的状态码
check_http(){
    status_code=$(curl -m 5 -s -o /dev/null -w %{http_code} $url)
}
 
while :
do
       check_http
       date=$(date +%Y%m%d-%H:%M:%S) 
#生成报警邮件的内容
       echo "当前时间为:$date
       $url服务器异常,状态码为${status_code}"
       
#指定测试服务器状态的函数，并根据返回码决定是发送邮件报警还是将正常信息写入日志
       if [ $status_code -eq 200 ];then
              echo "$url连接正常"
              break
       fi
       sleep 1
done

echo $nextPort > /app/wxxcx/pid

sed -i "s/$curPort/$nextPort/" /etc/nginx/sites-enabled/api.wxxcx.top
nginx -s reload

nextPortPid=$(netstat -anp|grep $curPort|awk '{printf $7}'|cut -d/ -f1)
nextPortPid=$(echo $nextPortPid | sed -e 's/-//g')
echo nextPortPid: $nextPortPid
kill $nextPortPid