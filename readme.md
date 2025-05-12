# realfin
**realfin** is a blockchain built using Cosmos SDK and Tendermint.

## Get started

```
realfin chain
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`.

### Web Frontend
These commands can be run within your scaffolded blockchain project. 

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/realfin@latest! | sudo bash
```
`username/realfin` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).


./realfind tx oracle submit-price BTC 65000 --from alice
./realfind query oracle list-price
./realfind query oracle show-price 0