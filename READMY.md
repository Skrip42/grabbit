# grabbitLayer
это маленькая обвязка поверх rabbitMQ написанная на го
по сути предоставляет http интерфейс для реализации RPC
т.е. вы посылаете запрос в грабит, от туда он попадает в очередь, с очереди возвращается в грабит, и тот делает запрос.
результать помещается в обратную очередь, и по возвращению в грабит делается запрос на адрес возврата

![схема](https://www.rabbitmq.com/img/tutorials/python-six.png)

Подробнее о том как работает RPC можно почитать [тут](https://www.rabbitmq.com/tutorials/tutorial-six-go.html)



## развертывание
- установите golang >= 1.15
- затяние репозиторий

```shell
git clone https://github.com/Skrip42/grabbitLayer.git
```

- сбидите проект

```shell go build cmd/grabbitlayer/main.go```
- откопируйте и отредактируйте config/config.yaml.example в config/config.yaml
- откопируйте и отредактируйте grabbitlayer grabbitlayer.service.example в /etc/systemd/system/grabbitlayer.service
- включите сервис и настройте автозапуск:

```shell
sudo systemctl daemon-reload
sudo systemctl enable grabbitlayer
sudo systemctl start grabbitlayer
```

## использование
в виде курл запроса:
```shell
    curl --location --request POST 'grabbitlayer.host:80?queue=queue_name&callback=callback.url' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "url":"http://ponhub.mycentra.ru/api/aggregation_switch?api_key=ponBaseReadApiKey&cluster.id=5",
        "method":"GET",
        "headers":"Connection: keep-alive\nAccept: *\/*\nUser-Agent: Mozilla\/5.0 (X11; Linux x86_64) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/88.0.4324.182 Safari\/537.36\nX-Requested-With: XMLHttpRequest\nAccept-Encoding: gzip, deflate\nAccept-Language: ru,ru-RU;q=0.9\nCookie: PHPSESSID=l2dophn4a1m2g4ash00pgrs8p6",
        "body":""
    }'
```
как можно увидеть параметры целевого запроса просто передаются в теле как json вида
```json
{
    "url":"target.url",
    "method":"TARGET_METHOD",
    "headers":"Target:Headers"
    "body":"target body"
}
```
сам же запрос имеет 2 параметра
- queue - префикс имени очередей через которые будет гулять запрос
- callback - куда будет направлен ответ

для php/symfony рекомендую готовый [сервис](https://gist.github.com/Skrip42/003f733c5ba218cd55961eb91627bf51)
