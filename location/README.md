# Location 📍
Is a service to get all the absolute path of system directory paths and application paths.

## Path variables
| Variable     | Purpose                                                  |
|--------------|----------------------------------------------------------|
| Home         | User home directory                                      |
| DataHome     | Persistent user data (e.g., templates, cache-like data)  |
| DataDirs     | List of paths persistent application data                |
| ConfigHome   |                                                          |
| BinHome      |                                                          |


## Linux 🐧
Paths:

| Variable  | Directory                                  | Purpose                                                     |
|-----------|--------------------------------------------|-------------------------------------------------------------|
| Home      | `~/`                                       |                                                             |
| BinHome   | `~/.local/bin/{appName}/`                  | User-specific executables (should be in PATH)               |
|           | `~/.local/share/{appName}/`	               |                                                             |
|           | `~/.local/state/{appName}/`	               | Transient state data (lock files, PID files, runtime state) |
|           | `~/.cache/{appName}/`                      | Non-essential, temporary cache data                         |
| CacheFile | `~/.local/share/{appName}/{appName}.cache` |                                                             |

### desktop
The path to the desktop file: `~/.local/share/applications/`

To add desktop file:
```bash
update-desktop-database ~/.local/share/applications/
```

### Icon
The path is in the `.desktop` file.

Normally this is where it is located:
- `~/.local/share/{appName}/assets/{appName}.svg`
- `~/.local/share/{appName}/assets/{appName}.png`

> [!NOTE]
> The two file format [SVG](https://developer.mozilla.org/en-US/docs/Web/SVG) and [PNG](https://www.w3.org/TR/png/)
> need to be present.

### mime
The path to the mime configuration file: `~/.local/share/mime/packages/{typeName}.xml`

`/home/jn/.config/mimeapps.list` -

To update mime:
```bash
update-mime-database ~/.local/share/mime/
```

To associate file extension (e.g.: `.hery`):
```bash
xdg-mime default {appName}.desktop application/x-hery
```

## Windows 🪟

## Mac OS X 🍎

