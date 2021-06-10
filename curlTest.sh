#!/bin/bash
for (( i = 0; i < 20; i++ ))
do
    curl --location --request POST '10.1.19.242:80?quque=grabbitTestQueue&callback=http://10.1.19.242:80/test' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "url":"http://ponhub.mycentra.ru/api/aggregation_switch?api_key=ponBaseReadApiKey&cluster.id=5",
        "method":"GET",
        "headers":"Host: ponbase.local\nConnection: keep-alive\nAccept: *\/*\nUser-Agent: Mozilla\/5.0 (X11; Linux x86_64) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/88.0.4324.182 Safari\/537.36\nX-Requested-With: XMLHttpRequest\nReferer: http:\/\/ponbase.local\/api\/\nAccept-Encoding: gzip, deflate\nAccept-Language: ru,ru-RU;q=0.9\nCookie: PHPSESSID=l2dophn4a1m2g4ash00pgrs8p6",
        "body":""
    }'
done
