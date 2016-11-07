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
  -t, --tee          Print STDIN to screen before posting (false)
  -s, --stream       Post messages to LINE Notify continuously (false)
      --config_file= Load the specified configuration file

Help Options:
  -h, --help         Show this help message
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

Executable Binaries
--

Those are on [GitHub Releases](https://github.com/moznion/linenotcat/releases)

Thanks
--

This tool is much inspired by [slackcat](https://github.com/vektorlab/slackcat) and some code is taken from that.

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

