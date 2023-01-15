# Go Template Online Parser
https://helm-go-template-online-production.up.railway.app/

# Description
GoTemplate online parser + helm parser online to make it easier to debug go-templates and helm templates

# Local Development

## UI Only
```bash
cd webui-vue
yarn serve
```

## Backend
```bash
./build.sh
docker run -it -e PORT=3000 -p 3000 docker.io/library/gotplonline
```
or 
```bash
./build.sh
go run main.go
```



