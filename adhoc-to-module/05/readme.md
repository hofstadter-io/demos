# 05 - enhancing the api server

Generating clients for the app

- apikeys
- graceful shutdown
- others echo/gorm things?
- config (viper?), (CUE?) (watch / reload)
- extras
  - cli tool (cobra)
  - client/{go}
- seed data

## go cli / client

```sh

# run the server
./app serve
./app alive

# test the go client
./app seed data.cue
./app get user --email tony@hof.io
```


