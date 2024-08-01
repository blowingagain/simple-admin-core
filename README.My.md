# 一、改动内容

## 1、修改token
> 需求：登录需要加上新字段：区域（regionId）

修改内容文件：
 - [login_logic.go](api%2Finternal%2Flogic%2Fpublicuser%2Flogin_logic.go)
 - [login_by_sms_logic.go](api%2Finternal%2Flogic%2Fpublicuser%2Flogin_by_sms_logic.go)
 - [login_by_email_logic.go](api%2Finternal%2Flogic%2Fpublicuser%2Flogin_by_email_logic.go)

代码如下：

新增：`jwt.WithOption("regionId", req.RegionId)`
```go
token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
	l.svcCtx.Config.Auth.AccessExpire, 
	jwt.WithOption("userId", user.Id), 
	jwt.WithOption("roleId",strings.Join(user.RoleCodes, ",")), 
	jwt.WithOption("deptId", user.DepartmentId),
	jwt.WithOption("regionId", req.RegionId))
if err != nil {
	return nil, err
}
```

## 2、新增Github CI

> 路径：.github/workflows/ci.yml


```yaml
name: CI

on: [ push ]


jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.22.5

      - name: Cache Go modules
        uses: actions/cache@v3
        with:
          path: |
            ~/.cache/go-build
            ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-

      - name: Build Go project
        run: |
          go env -w CGO_ENABLED=0
          go env -w GOPROXY=https://goproxy.cn,direct
          go env -w GOPRIVATE=github.com
          go mod tidy
          go build -trimpath -ldflags "-s -w" -o core_api api/core.go
          go build -trimpath -ldflags "-s -w" -o core_rpc rpc/core.go

      - name: Verify core_api exists
        run: |
          if [ ! -f core_api ]; then
            echo "core_api file not found!"
            exit 1
          fi
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Log in to Alibaba Cloud ACR
        env:
          ACR_REGISTRY: ${{ secrets.ACR_REGISTRY }}
          ACR_USERNAME: ${{ secrets.ACR_USERNAME }}
          ACR_PASSWORD: ${{ secrets.ACR_PASSWORD }}
        run: echo $ACR_PASSWORD | docker login $ACR_REGISTRY -u $ACR_USERNAME --password-stdin

      - name: Build and push Docker image - API
        env:
          ACR_REGISTRY: ${{ secrets.ACR_REGISTRY }}
        run: |
          docker build -t $ACR_REGISTRY/dkzx_test/core-api-docker:latest -f Dockerfile-api .
          docker push $ACR_REGISTRY/dkzx_test/core-api-docker:latest

      - name: Build and push Docker image - RPC
        env:
          ACR_REGISTRY: ${{ secrets.ACR_REGISTRY }}
        run: |
          docker build -t $ACR_REGISTRY/dkzx_test/core-rpc-docker:latest -f Dockerfile-rpc .
          docker push $ACR_REGISTRY/dkzx_test/core-rpc-docker:latest
```
上面的 actions 最后将 API RPC 一起打包最后发布到2个阿里云ACR 仓库