# proxy-skill [![Go Report Card](https://goreportcard.com/badge/github.com/azzzak/proxy-skill)](https://goreportcard.com/report/github.com/azzzak/proxy-skill)

Позволяет разработчику навыков для голосового помощника [Алиса](https://alice.yandex.ru) тестировать навыки с помощью смартфона или колонки.

## Использование

Для полноценного использования необходимо создать приватный навык и пройти его модерацию. После этого запустить прокси на вебхуке, который присвоен приватному навыку, перенаправив трафик на нужный адрес.

`docker run -d -p 3000:3000 -e "FORWARD_TO=https://a1b2c3d4.ngrok.io" azzzak/proxy-skill`
