# khaiii-api: khaiii API server

![khaiii_logo](_asset/khaiii_logo_256.webp)

## Build

Update docs:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

Build container image
```sh
make build_image
```
