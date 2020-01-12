# Crypto Qrcode Generator

[![Build Status](https://secure.travis-ci.org/yassun/crypto-qrcode-generator.png?branch=master)](http://travis-ci.org/yassun/crypto-qrcode-generator)

http://crypto-qr-generator.site/

Sample Address : 3HXM3iSGWzuXwHYzyquhmtMn2A5uXdLgQh

<img width="240" alt="スクリーンショット 2019-09-01 23 02 42" src="https://user-images.githubusercontent.com/2255617/64077588-d6e2cc00-cd0c-11e9-9572-d02a8d3b1c7f.png"> <img width="240" src="https://user-images.githubusercontent.com/2255617/64077835-0bf01e00-cd0f-11e9-8d08-c4676a17df94.png">

# About
- This WebApp is simple crypto qrcode generator. 
- Hosted on Kubernetes(GKE). 
- The URI based on the BIP-21 standard(only BTC).

# Configuration

http://yasun.hatenablog.jp/entry/2019/09/03/145514

![Crypto Qrcode Generator Architecture](https://user-images.githubusercontent.com/2255617/64079033-5c6e7800-cd1d-11e9-84d8-fbd264d600ac.png)

- Backend API: Golang/Echo
- Frontend: Vue.js/Vuetiff
- Production Environment: Kubernetes(GKE)
- Development Environment: docker-compose
- CI/CD: TravisCI
