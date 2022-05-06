# Version Comparison Utility

This repository is a golang utility which accepts two input versions and compare them. It uses [hashicorp/go-version](https://github.com/hashicorp/go-version) library for version parsing and comparison.

## Comparison Rules

Following are the three followed during version comparison:
- If version1 > version2, return 1.
- If version1 < version2, return -1.
- Otherwise, return 0.

## Assumptions

This utility assumes that the version strings are non-empty and contain only digits and the 'dot' character. The 'dot' character does not represent a decimal point and is used to separate number sequences. 
For example, '2.5' is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

## Prerequisites

Before you start, please make sure you have Go installed in your system. If not, please use the following link to install Golang:
https://golang.org/doc/install

## Getting Started

Clone the git repository in your system and then cd into project root directory

```bash
$ git clone https://github.com/RashikaAggarwal/version-utility.git
$ cd version-utility
```

Build your tool by executing the following steps
```bash
$ cd utility
$ go build
```

## Sample Outputs

This utility takes two version inputs separated by white space as command line arguments. See below examples
```bash
$ ./version-utility.exe 1.2 1.3.4
-1
```

```bash
$ ./version-utility.exe 1.2.9.9 1.2
1
```

```bash
$ ./version-utility.exe 1.1 1.1
0
```

In case of failures(say empty input or invalid version numbers), it prints error message and gets exited.
```bash
$ ./version-utility.exe
Error: No Input found, please enter valid input
```

```bash
$ ./version-utility.exe 1.1 test
Error: Invalid input version
```

## Unit Testing

This repository contains unit test cases which provide the industry standard code coverage. The test cases can be executed using following commands:
```bash
$ cd version-utility/utility
$ go test
```

To get the code coverage, use the following command:
```bash
$ go test -cover
```
