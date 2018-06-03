# get-license

## Description

Easily get Open Source Licenses with this simple CLI tool. It works by scraping the [choosealicense.com](https://choosealicense.com/) site and outputting the license text to STDOUT.

You should definitely checkout [choosealicense.com](https://choosealicense.com/) if you're unsure what license is the best fit for your site.

## Usage

Currently the following licenses are available:

- *apache* - [Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)
- *mit* - [MIT](https://choosealicense.com/licenses/mit/)
- *mozilla* - [Mozilla Public License 2.0](https://choosealicense.com/licenses/mpl-2.0/)
- *gpl* - [GNU GPL 3.0](https://choosealicense.com/licenses/gpl-3.0/)

```shell
# get apache license
getlicense -name Apache >> LICENSE
```

## Installation

If you have Go installed then you can run the following:

```shell
go get -u github.com/davyj0mes/get-license
```

Binaries for Windows, Linux and OSX can be found in [releases](./releases)

## License

[MIT](./LICENSE)
