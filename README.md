# gorbac-sample

# casbin-sample
・[Casbin](https://casbin.org/en/)

## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．アプリケーションの起動

① アプリケーションコンテナ内へ移動

```
$ docker exec -it gorbac-sample bash
```

② アプリケーションの起動

```
root@fe385569a625:/go/src/app# go run main.go
```