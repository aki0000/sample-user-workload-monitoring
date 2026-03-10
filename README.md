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
   - ノードが `amd64` の場合は以下（タグは毎回更新する）

```bash
podman build --platform linux/amd64 -t quay.io/akhino/latency-app:v1.0.1 ./
podman push quay.io/akhino/latency-app:v1.0.1
```

1. `openshift/app/app.yaml` の `image` をPush先に変更
2. 適用

```bash
oc apply -f openshift/app/app.yaml
```

Route URL 確認:

```bash
oc get route latency-app -n latency-app
```

## OpenShift User Workload Monitoring

User Workload Monitoring を有効化:

```bash
oc apply -f openshift/user-workload-monitoring/cluster-monitoring-config-map.yaml
```

ServiceMonitor を作成:

```bash
oc apply -f openshift/user-workload-monitoring/latency-app-service-monitor.yaml
```

Prometheus で取得できるメトリクス:

- `http_request_duration_seconds_bucket`
- `http_request_duration_seconds_sum`
- `http_request_duration_seconds_count`
