rm -f /usr/bin/genmetrics
cp genmetrics /usr/bin/

if [ ! -e /usr/bin/genmetrics ];then
        echo "genmetrics install failed!"
        return 1
fi

echo "genmetrics install ok."