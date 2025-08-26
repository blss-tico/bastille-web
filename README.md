# Bastille-Web (This project is under development...)

![GitHub repo size](https://img.shields.io/github/repo-size/iuricode/README-template?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/iuricode/README-template?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/iuricode/README-template?style=for-the-badge)
![Bitbucket open issues](https://img.shields.io/bitbucket/issues/iuricode/README-template?style=for-the-badge)
![Bitbucket open pull requests](https://img.shields.io/bitbucket/pr-raw/iuricode/README-template?style=for-the-badge)

> A web interface for the FreeBSD bastille jails management. 

## 💻 Requisites

Before begin, you must to know that this project was tested with:

- FreeBSD 14;
- Golang 1.24;
- Bastille 1.0.1.250714;

## 🚀 Install bastille-web

To install the bastille-web, follow these steps:

FreeBSD:

```
git clone https://github.com/blss-tico/bastille-web.git
cd bastille-web
go build
```

## ☕ Usando bastille-web

To use bastille-web, with default ip/port option:
```
sudo ./bastille-web
```

To use bastille-web, with user defined ip/port option:
```
sudo ./bastille-web 10.0.0.1:80
```

## 📝 Licença

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.
