# sound-notifier

sound-notifier is a small sound controller that sends system notifications. That's it.

## Documentation
```
$> go run sound-notifier.go --help
usage: sound-notifier [<flags>] <command> [<args> ...]

sound controller for Linux that triggers notifications

Flags:
  --help    Show context-sensitive help (also try --help-long and --help-man).
  --silent  Hide notification when volume changes

Commands:
  help [<command>...]
    Show help.

  set <volume>
    Manually set volume (between 0 and 100)

  up <volume>
    Increase volume

  down <volume>
    Decrease volume

  mute
    Mute switch (automatic On/Off)
```
