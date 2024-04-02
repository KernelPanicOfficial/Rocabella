# Rocabella

Sniffer files generator.

<p align="center">
  <img width="350" height="350" src="/Pictures/Rocabella-Logo.png"><br /><br />
  <img alt="Static Badge" src="https://img.shields.io/badge/License-MIT-green?link=https%3A%2F%2Fgithub.com%2Fnickvourd%2FRocabella%2Fblob%2Fmain%2FLICENSE">
  <img alt="Static Badge" src="https://img.shields.io/badge/Version-1.0%20(Death Star)-red?link=https%3A%2F%2Fgithub.com%2Fnickvourd%2FRocabella%2Freleases"><br /><br />
  <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/nickvourd/Rocabella?logoColor=yellow">
  <img alt="GitHub forks" src="https://img.shields.io/github/forks/nickvourd/Rocabella?logoColor=red">
  <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/nickvourd/Rocabella?logoColor=blue">
</p>

## Description

Rocabella is an open-source Go tool that generates sniffing files.

It supports a variety of sniffing file formats, including:

- searchConnector-ms
- library-ms
- LNK shortcuts
- URL shortcuts

Rocabella is written in Golang, a cross-platform language, enabling its use on both Windows and Linux systems.

For command-line usage and examples, please refer to our [Wiki](https://github.com/nickvourd/Rocabella/wiki).

> If you find any bugs, don’t hesitate to [report them](https://github.com/nickvourd/Rocabella/issues). Your feedback is valuable in improving the quality of this project!

## Disclaimer

The author and contributors of this project are not liable for any illegal use of the tool. It is intended for educational purposes only. Users are responsible for ensuring lawful usage.

## Table of Contents
- [Rocabella](#rocabella)
  - [Description](#description)
  - [Disclaimer](#disclaimer)
  - [Table of Contents](#table-of-contents)
  - [Acknowledgement](#acknowledgement)
  - [Installation](#installation)
  - [Usage](#usage)
  - [References](#references)
 
## Acknowledgement

Special thanks to my brothers [@S1ckB0y1337](https://twitter.com/S1ckB0y1337) and [@MikeAngelowtt](https://twitter.com/MikeAngelowtt), who provided invaluable assistance during the beta testing phase of the tool.

## Installation

You can use the [precompiled binaries](https://github.com/nickvourd/Rocabella/releases), or you can manually install Rocabella by following the next steps:

1) Clone the repository by executing the following command:

```
git clone https://github.com/nickvourd/Rocabella.git
```

2) Once the repository is cloned, navigate into the Rocabella directory:

```
cd Rocabella
```

3) Install the third-party dependencies:

```
go mod download
```

4) Build Rocabella with the following command:

```
go build Rocabella
```

## Usage

:information_source: Please refer to the [Rocabella Wiki](https://github.com/nickvourd/Rocabella/wiki) for detailed usage instructions and examples of commands.

## References

- [Farming for Red Teams: Harvesting NetNTLM by MDSec](https://www.mdsec.co.uk/2021/02/farming-for-red-teams-harvesting-netntlm/)
- [The Latest in Software Functionality Abuse: URL Internet Shortcut Files Abused to Deliver Malware by Cofense](https://cofense.com/blog/latest-software-functionality-abuse-url-internet-shortcut-files-abused-deliver-malware/)
- [SearchConnector-ms by Filesec](https://filesec.io/searchConnector-ms)
- [Library-ms by Filesec](https://filesec.io/library-ms)
- [LNK by Filesec](https://filesec.io/lnk)
- [URL by Filesec](https://filesec.io/url)
