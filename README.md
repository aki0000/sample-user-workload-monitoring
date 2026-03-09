# sample-user-workload-monitoring

ユーザがアクセスすると、レスポンス内に latency を表示する最小の Go Web アプリです。

## ローカル実行

```bash
go run main.go
```

別ターミナルでアクセス:

```bash
curl http://localhost:8080/
```

出力例:

```text
hello from go on OCP
latency: 1.4µs
```

## OCP (OpenShift) デプロイ

1. イメージをビルドしてレジストリへPush（例: `quay.io`）
2. `openshift/app/app.yaml` の `image` をPush先に変更
3. 適用

```bash
oc apply -f openshift/app/app.yaml
```

Route URL 確認:

```bash
oc get route latency-app -n latency-app
```
