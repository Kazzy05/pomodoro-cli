# pomodoro-cli

## How to use

After you build and install it (`go build -o pomo main.go && mv pomo /usr/local/bin/`), you can use it in several ways:

- Default (25/5):

```sh
pomo
```

- Custom (e.g., 50 min work, 10 min break):

```sh
pomo -w 50 -b 10
```

- Check Help:

```sh
pomo --help
```
