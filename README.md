# log4go

## Description

This repository is reconstructed from  jeanphorn/log4go, which is a logging package similar to log4j for the Go programming language.

## Features
## Usage

First, get the code from this repo. 

```go get github.com/spafka/log4go```

Then import it to your project.

```import log "github.com/spafka/log4go"```


## Examples

The config file is optional, if you don't set the config file, it would use the default console config. && console must be true

Here is a Json config example:
```json
{
  "console": {
    "enable": true,
    "level": "DEBUG",
    "pattern": "[%D %T]  %M"
  },
  "files": [
    {
      "enable": true,
      "level": "DEBUG",
      "filename": "console.log",
      "pattern": "[%D %T]  %M",
      "category": "DEFAULT",
      "rotate": true,
      "maxsize": "1024M",
      "maxlines": "10K",
      "daily": true
    },
    {
      "enable": true,
      "level": "DEBUG",
      "filename": "RemoteControl.log",
      "pattern": "[%D %T]  %M",
      "category": "RemoteControl",
      "rotate": true,
      "maxsize": "1024M",
      "maxlines": "10K",
      "daily": true
    }
  ]
}
```

Code example:

```
package main

import (
	log "github.com/spafka/log4go"
)

func main() {
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log.LoadConfiguration("./example.json")

	log.LOGGER("Test").Info("category Test info test ...")
	log.LOGGER("Test").Info("category Test info test message: %s", "new test msg")
	log.LOGGER("Test").Debug("category Test debug test ...")

	// Other category not exist, test
	log.LOGGER("Other").Debug("category Other debug test ...")

	// socket log test
	log.LOGGER("TestSocket").Debug("category TestSocket debug test ...")

	// original log4go test
	log.Info("normal info test ...")
	log.Debug("normal debug test ...")

	log.Close()
}

```

The output like:

## Thanks

Thanks alecthomas for providing the [original resource](https://github.com/alecthomas/log4go).

In old
   log.Info which default category is nil, && i make it to "DEFAULT",so U can Define "DEFAULt" FILEAppler which can log this in a file && use daily log funcntion rather than
   echo it to  >> nohup.out or so

todo add additivity which definded in Log4j2