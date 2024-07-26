# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [1.3.0](https://github.com/Utconnect/coffer/compare/v1.2.0...v1.3.0) (2024-07-26)


### Features

* login with github token ([78fbdc0](https://github.com/Utconnect/coffer/commit/78fbdc02a391e3249a70015c870af4e474a1d338))


### CI/CD

* ignore messages from vault when run unseal script ([3262534](https://github.com/Utconnect/coffer/commit/32625348c7f1a6dfa107bb3307ec9cb9a5dda527))

## [1.2.0](https://github.com/Utconnect/coffer/compare/v1.1.3...v1.2.0) (2024-07-24)


### Features

* allow patching secrets by using kv2 ([827ef15](https://github.com/Utconnect/coffer/commit/827ef157432b4cbd884c48b86c3b2b39269c813a))


### CI/CD

* add script to auto-seal vault ([8a86d67](https://github.com/Utconnect/coffer/commit/8a86d671773d029fbba18f13a2f224b8116ea5b8))
* allow to scan all files in project ([d993f20](https://github.com/Utconnect/coffer/commit/d993f203a95af93defc4021078f524e2f539262e))

### [1.1.3](https://github.com/Utconnect/coffer/compare/v1.1.2...v1.1.3) (2024-07-22)


### CI/CD

* add static properties to sonar ([c00cd30](https://github.com/Utconnect/coffer/commit/c00cd30f618b3c1d677dc003471efc570c150329))

### [1.1.2](https://github.com/Utconnect/coffer/compare/v1.1.1...v1.1.2) (2024-07-21)


### CI/CD

* update sonar version when release ([95b73dc](https://github.com/Utconnect/coffer/commit/95b73dcea279ac36e1a101dc444006f03f754393))

### [1.1.1](https://github.com/Utconnect/coffer/compare/v1.1.0...v1.1.1) (2024-07-20)


### CI/CD

* copy only /src for build stage ([f0a4b3b](https://github.com/Utconnect/coffer/commit/f0a4b3b3c6e04c2e4637ddbd72491694c4adc40f))
* run image with non-root user ([23b8d2a](https://github.com/Utconnect/coffer/commit/23b8d2a00bdecd879f46006a226359d621985b0d))
* use `go-alpine` to reduce docker image size ([b3ed288](https://github.com/Utconnect/coffer/commit/b3ed288a395d1b9232932bac2cabada2273ecc27))

## 1.1.0 (2024-07-17)


### Features

* get secret ([5c95f98](https://github.com/Utconnect/coffer/commit/5c95f981c92f4ffa4034226dde6a3c68ae847a9e))


### Bug Fixes

* Can not up docker compose. ([1197e2f](https://github.com/Utconnect/coffer/commit/1197e2fdec3dcb9464f36bed690ceda1c79d7579))
* cannot use vault ui via docker ([68f254a](https://github.com/Utconnect/coffer/commit/68f254a77083e3893fad4af2f713295297a7de13))


### CI/CD

* add `release` workflow ([9d3e3de](https://github.com/Utconnect/coffer/commit/9d3e3dea9fef1a74856d637c53de35462c0da93c))
