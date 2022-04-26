# K8S Go API

## 使用方法

### Dockerイメージのビルド

clone後のルートディレクトリに移動し、以下コマンドを実行

```bash
docker image build -t go-sample-api:dev .
```

### kubectlのインストール

1. Docker Desktopにて、k8sを使える状態にする
2. 以下コマンドを実行

```bash
brew install kubectl
```

### kubectlの実行

- deploymentの実行

```bash
cd k8s
kubectl apply -f deployment.yml
```

- serviceの実行

```bash
cd k8s
kubectl apply -f service.yml
```

- podの確認

```bash
$ kubectl get po  
NAME                           READY   STATUS    RESTARTS   AGE
go-sample-api-dd4ccdc7-2zhzr   1/1     Running   0          33m
go-sample-api-dd4ccdc7-9jtc5   1/1     Running   0          33m
go-sample-api-dd4ccdc7-hpdlb   1/1     Running   0          33m
go-sample-api-dd4ccdc7-r2sxc   1/1     Running   0          33m
go-sample-api-dd4ccdc7-vz65q   1/1     Running   0          33m
```

- curlでAPIを叩く

```bash
$ curl localhost:80/user/ja
こんにちは、user!
$ curl localhost:80/ryoichi/fr
Bonjour ryoichi!
```

- logの確認

```bash
$ kubectl logs go-sample-api-dd4ccdc7-2zhzr
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.helloWorld (3 handlers)
[GIN-debug] GET    /:user                    --> main.helloUser (3 handlers)
[GIN-debug] GET    /:user/:region            --> main.greedingUser (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080

$ kubectl logs go-sample-api-dd4ccdc7-9jtc5 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.helloWorld (3 handlers)
[GIN-debug] GET    /:user                    --> main.helloUser (3 handlers)
[GIN-debug] GET    /:user/:region            --> main.greedingUser (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/04/26 - 14:09:47 | 200 |     943.542µs |    192.168.65.3 | GET      "/user/ja"
[GIN] 2022/04/26 - 14:09:57 | 200 |     115.417µs |    192.168.65.3 | GET      "/ryoichi/fr"
```

## 改善点

1. alpineベースのGoイメージを使えばDockerfileをもっと簡単にできる
