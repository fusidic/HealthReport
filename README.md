# HealthReport

每天自动上传健康信息

## 一般使用

### 拉取仓库

在服务器上运行：

```bash
$ git clone https://github.com/fusidic/HealthReport.git
```

### 注册个人信息

在 `./info/user.json` 中添加自己的信息，注意格式。

### 定时启动

通过 `crontab` 设置定时启动脚本，`LinuxGo` 适用于Linux amd64版本的机型，`WinGo` 适用于windows amd64 机型，其他机型使用 `go build` 编译

```bash
$ crontab -e
```

写入：

```bash
# 每天 2:10 AM 运行，注意机器所在时区
10 2 * * * {project_path}/report.sh
```

<br/>

## Docker 部署

### 获取镜像

在本文仓库路径下：

```bash
$ docker build -t <ID>/health-report .
```

构建完成后，可以看见本地镜像已经存在：

```bash
$ docker images                         
REPOSITORY                  TAG          IMAGE ID            CREATED             SIZE
fusidic/health-report       v0.1         a66c8a2c4ce5        43 minutes ago      13.8MB
```

也可以使用我上传的镜像：

```bash
$ docker pull fusidic/health-report:latest
```

不过比较遗憾的是，**暂时没法定时执行**，当然你也可以试试 Kubernetes 的 `CronJob`

<br/>

## CronJob

`CronJob` 为 Kubernetes 的一种资源类型，用于定期执行任务，好处是只用一个 `YAML` 文件就能准确高效且安全的指定定时任务。

### 如何使用

我在仓库中同样提供了一个名为 `cronjob.yaml` 的文件，如果你~~碰巧~~有 Kubernetes 的集群的话，直接使用：

```bash
$ kubectl create -f cronjob.yaml
```

