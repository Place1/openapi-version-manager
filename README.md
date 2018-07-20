# openapi-version-manager

## What's this?
`openapi-version-manager` is a tiny CLI tool to help you manage different versions of `openapi-generator` on your local
machine.

```bash
$ openapi-version-manager --help
NAME:
   openapi-version-manager - A new cli application

USAGE:
   openapi-version-manager-darwin-amd64 [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     current  show the current openapi codegen version
     list     list available openapi codegen versions
     use      use the specified openapi-generator-version
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

**list available openapi versions**
```bash
$ openapi-version-manager list
3.0.0-rc1
3.0.0-rc0
2.3.1
2.3.0
2.2.3
2.2.2
2.2.1
2.2.0
2.1.6
2.1.5
```

**use a specific openapi-generator version**
```bash
$ openapi-version-manager use 2.2.3
downloading http://search.maven.org/remotecontent?filepath=io/openapi/openapi-generator-cli/2.2.3/openapi-generator-cli-2.2.3.jar
 3.25 MiB / 13.19 MiB [===================>                                                                     ]  25% 00m01
```
```bash
$ openapi-generator version
2.2.3
```

## Installation

### MacOS
```bash
curl -Lo openapi-version-manager https://github.com/Place1/openapi-version-manager/releases/download/v1.0.0/openapi-version-manager-darwin-amd64
chmod +x ./openapi-version-manager
mv ./openapi-version-manager /usr/local/bin/
openapi-version-manager --help
```

### Linux
```bash
curl -Lo openapi-version-manager https://github.com/Place1/openapi-version-manager/releases/download/v1.0.0/openapi-version-manager-linux-amd64
chmod +x ./openapi-version-manager
mv ./openapi-version-manager /usr/local/bin/
openapi-version-manager --help
```

### Windows
1. Download the windows binary from the release page
2. rename it to `openapi-version-manager.exe`
3. execute it using powershell
    - `openapi-version-manager` required admin privileges on windows because it writes the `openapi-generator`
      executable to `C:\Windows\System32\openapi-generator`. I'm not familiar with windows but i'd love
      a PR that changes this behaviour so that admin privileges are not required.
