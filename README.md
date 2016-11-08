linenotcat
==

A command line tool to send messages to [LINE Notify](https://notify-bot.line.me/).

linenotcat = LINE NOTify + cat

Getting Started
--

### 1. Put configuration file

```sh
$ echo 'YOUR_ACCESS_TOKEN' > $HOME/.linenotcat
```

It is deadly simple; this configuration file contains only access token.

Default, `linenotcat` reads configuration file that is on `$HOME` directory.
If you want to force it to read arbitrary file, please run the command with `--config_file` option.

### 2. Send message via STDIN

```sh
$ echo 'Hello world!' | linenotcat
```

### 2'. Send contents of file

```sh
$ linenotcat ./your/awesome/file.txt
```

Options
--

```
Application Options:
  -m, --message=     Send a text message directly
  -i, --image=       Send an image file
  -t, --tee          Print STDIN to screen before posting (false)
  -s, --stream       Post messages to LINE Notify continuously (false)
      --config_file= Load the specified configuration file
      --status       Show connection status that belongs to the token (false)

Help Options:
  -h, --help         Show this help message
```

### -m, --message

Send a text message directly.

e.g.

```sh
$ linenotcat -m 'Hello world!'
```

Then `Hello world!` text will be sent via LINE Notify.

### -i, --image

Send an image file.

e.g.

```sh
$ linenotcat -i /path/to/your/awesome/image.png
```

Then `image.png` image will be sent via LINE Notify with default message (default message=`Image file`).

If you want to send an image with arbitrary message, please use with `-m` option.

e.g.

```sh
$ linenotcat -i /path/to/your/awesome/image.png -m 'Super duper image!'
```

### -t, --tee

Print STDIN to screen before posting.

### -s, --stream

Stream messages to LINE Notify continuously.

e.g.

```sh
$ tail -f /your/awesome/error.log | linenotcat --stream
```

Then contents of `error.log` are notified continuously to LINE Notify until process is died.

### --config_file

e.g.

```sh
$ echo 'Hello world!' | linenotcat --config_file="/your/path/to/config"
```

Then this command loads token information from your specified configuration file.

### --status

Show connection status that belongs to the token.

e.g.

```sh
$ linenotcat --status
{"status":200,"message":"ok","targetType":"USER","target":"moznion"}
```

Executable Binaries
--

Those are on [GitHub Releases](https://github.com/moznion/linenotcat/releases)

Thanks
--

This tool is much inspired by [slackcat](https://github.com/vektorlab/slackcat) and some code is taken from that.

For developers
--

This project depends on [glide](https://github.com/Masterminds/glide).
So if you want to build this project, please build with glide.

Example:

```sh
$ glide install
$ go build cmd/linenotcat/linenotcat.go
```

Or you can build with `make` command

```sh
$ make VERSION=1.2.3
```

See Also
--

- [LINE Notify](https://notify-bot.line.me/)

Author
--

moznion (<moznion@gmail.com>)

License
--

```
The MIT License (MIT)
Copyright © 2016 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

