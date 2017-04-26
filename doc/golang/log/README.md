[*back to doc*](https://github.com/malw2020/learn/tree/master/doc#contents)<br>

### Seelog

Seelog is a powerful and easy-to-learn logging framework that provides functionality for flexible dispatching, filtering, and formatting log messages. It is natively written in the Go programming language.

from: [*https://github.com/cihub/seelog*](https://github.com/cihub/seelog)



[↑ top](#contents)
<br><br>


## Quick-start

	package main

	import (
		seelog "github.com/cihub/seelog"
	)

	func main() {
		logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
		if err != nil {
			seelog.Critical("err parsing config log file, ", err)
			return
		}

		seelog.ReplaceLogger(logger)
		defer seelog.Flush()

		seelog.Info("tester start...")
	}


[↑ top](#contents)
<br><br>

## Installation
	go get -u github.com/cihub/seelog

[↑ top](#contents)
<br><br>

## Documentation

  Seelog has github wiki pages, which contain detailed how-tos references: [*https://github.com/cihub/seelog/wiki*](https://github.com/cihub/seelog/wiki)


[↑ top](#contents)
<br><br>






