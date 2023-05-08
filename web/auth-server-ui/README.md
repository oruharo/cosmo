# cosmo-auth-server login UI

## How to create this project

```
git clone https://github.com/cosmo-workspace/cosmo
cd cosmo/

WEBUI=$(pwd)/web/auth-server-ui

cd web/

npm create vite@latest -- auth-server-ui --template react-ts

cd $WEBUI

yarn add \
  @mui/material @emotion/react @emotion/styled \
  @mui/icons-material \
  react-error-boundary \
  @bufbuild/connect-web @bufbuild/protobuf

```

## How to start

```
cd web/auth-server-ui
yarn install && yarn dev
```

## How to server test

```
# build
cd web/auth-server-ui
yarn build --base=/server-driver-test --outDir=build_test

../../hack/download-certs.sh dashboard-server-cert cosmo-system

COSMO_USER_NAME=xxxxxxxx
{
pkill -ef "main echo-server" ; \
go run ../../hack/echo-server/main.go echo-server &  \
go run ../../hack/server-driver/main.go --port=9999 --target-port=8888 --user=${COSMO_USER_NAME} --auth-ui=./build_test/ --auth-url=https://cosmo-dashboard.cosmo-system.svc.cluster.local:8443 ; \
pkill -ef "main echo-server"
}

COSMO_USER_NAME=xxxxxxxx
curl  -v  -H "Content-Type: application/json" -d '{"userName":"'${COSMO_USER_NAME}'","password":"yyyyyyy"}' \
    http://localhost:9999/server-driver-test/authproxy.v1alpha1.AuthProxyService/Login
```