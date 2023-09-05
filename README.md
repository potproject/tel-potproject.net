# tel-potproject.net

Telnet Server Homepage.

```bash
$ telnet potproject.net
Connected to potproject.net.
Escape character is '^]'.

----------
You must not use Shift_JIS encoding to view this page!!!
Please use UTF-8 encoding.
----------
□□□□ あなたは 00036 人目の訪問者です □□□□
----------
 ∧＿＿∧ 
( ´・ω・）   _                   _           _                    _
|ﾉ| ﾉ)ﾉ)|ゝ | |                 (_)         | |                  | |
ﾒ＿ > ❙ ﾚﾉ_ | |_ _ __  _ __ ___  _  ___  ___| |_       _ __   ___| |_
| '_ \ / _ \| __| '_ \| '__/ _ \| |/ _ \/ __| __|     | '_ \ / _ \ __|
| |_) | (_) | |_| |_) | | | (_) | |  __/ (__| |_   _  | | | |  __/ |_
| .__/ \___/ \__| .__/|_|  \___/| |\___|\___|\__| (_) |_| |_|\___|\__|
| |             | |            _/ |
|_|             |_|           |__/

# Profile
Name: potpro (ぽとぷろ)
Blog: blog.potproject.net (https://blog.potproject.net)
Github: @potproject (https://github.com/potproject)
X(Twitter): @potpro (https://twitter.com/potpro)
Mastodon: https://mastodon.potproject.net/@potpro
Bluesky: https://bsky.app/profile/potp.ro
----------

Connection closed by foreign host.
```

## Runnning
### Local
```bash
# Ja: ポート23での起動は拒否される可能性が高いので、1023以上のポートに変更すると良いです
# En: It is likely to be rejected by starting at port 23, so it is good to change it to a port above 1023
go run main.go
```


### Docker
```bash
docker-compose build
docker-compose up -d
```

## License
MIT