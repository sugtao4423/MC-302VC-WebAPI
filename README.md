# MC-302VC-WebAPI
自宅の給湯器の MC-302VC をWebから操作できるようにするやつ  
MC-302VC は ECHONET Lite 規格なのでそれをGoから操作してあげる

* [MC-302VC の ECHONET Lite 仕様](https://echonet.jp/introduce/gz-000732-2/)

## Compile
基本的には `go build` でOK

ただし `assets/` 配下を変更した場合は `go generate ./public` をして内容を更新する必要がある  
なお、HTML等も一緒にバイナリ化されるため同ディレクトリに `assets/` を配置する必要はない

## Port
Port | Note
--- | ---
8080/tcp (default) | Webインターフェースのポート
3610/udp (require) | ECHONET Lite からのレスポンスを受け取るため必須

## Option
Arg | Note | Default
--- | --- | ---
`-addr` | MC-302VC のIPアドレス | require
`-port` | Webインターフェースのポート | `8080`
`-user` | API BasicAuth username | `空`
`-pass` | API BasicAuth password | `空`

`-user` と `-pass` の両方が空の場合は `/api` 配下への認証が不要になる  
そうでない場合は `/api` 配下のみに認証がかかる

## API
* Base url: `/api`
* Request header: `Content-Type: application/json`

#### エラー系
```
{ "error": "Error Reason" }
```

#### 情報取得
```
GET /status
===>
{
  "operation_status": true,
  "water_temperature": 35,
  "bath_temperature": 44,
  "bath_auto_timer_status": false,
  "bath_auto_timer_time": [
    7,
    0
  ],
  "bath_operation_status": false,
  "bath_auto_mode_status": false,
  "bath_additional_heating_status": false
}
```

#### 風呂予約タイマーON/OFF
```
POST /bathAutoTimer
{ "status": true / false }
===>
{ "ok": true }
```

#### 風呂予約タイマー時刻設定
```
POST /bathAutoTimer/time
{ "hour": 7, "minute": 0 }
===>
{ "ok": true }
```

#### 風呂自動ON/OFF
```
POST /bath/auto
{ "status": true / false }
===>
{ "ok": true }
```

#### 風呂追い焚きON/OFF
```
POST /bath/additionalHeating
{ "status": true / false }
===>
{ "ok": true }
```

## systemd
```
# /etc/systemd/system/MC-302VC-WebAPI.service
[Unit]
Description=MC-302VC WebAPI server
After=network.target

[Service]
Type=simple
ExecStart=/path/to/MC-302VC-WebAPI
Restart=always

[Install]
WantedBy=multi-user.target
```

## ReverseProxy with sub directory
```
location /mc-302vc-webapi/ {
    rewrite ^/mc-302vc-webapi(.*) /$1 break;
    proxy_pass http://ip-for-hosting-server:port$1;
}
```
