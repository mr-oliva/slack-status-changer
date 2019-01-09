# slack-status-changer

## 使い方

1. go get github.com/bookun/slack-status-changer

2. ホームディレクトリに `.slack.yml` を配置する

``` ~/.slack.yml
internal_webapp_url: <オフィスからしか閲覧できないURLを記載>
tokens:
  # workspackeのtokenを記載
  - xoxp-xxxxx
```
[.slack.yml.sample](https://github.com/bookun/slack-status-changer/blob/master/.slack.yml.sample)を参考にしてください。
web版slackからログインしたことのあるworkspaceについては[ここ](https://api.slack.com/custom-integrations/legacy-tokens)からtokenを発行できるようです。

3. slack-status-changer を実行。
internal_webapp_urlで記載したURLからのレスポンスコードが200以外の場合はslack上のステータスが :house: に、200が来た場合 :office: になります。
