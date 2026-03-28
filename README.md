# Bing Wallpaper for macOS 🌅

A simple macOS app written in Go that fetches the daily Bing Wallpaper and sets it as your desktop background.

> [!NOTE]  
> This was created mostly as an excuse to play with go and github actions.

## Features

- Built as a native macOS `.app` bundle containing a Go executable.
- Runs invisibly in the background.
- Downloads the high-resolution (`_UHD.jpg`) image from the Bing API.

## How it works
When opened, the app reads the Bing Image Archive API. If today's wallpaper isn't already saved in `~/Pictures/BingWallpapers`, it downloads the image and uses macOS System Events to update the background of all desktop spaces.

## User Experience 
When you run the app, it works quietly in the background without interrupting you. This is so it can be added as a login item and have the wallpaper changed on boot. 
1. **No Dock icons or popups:** The app is configured as a background process (`LSUIElement`), so you won't see a terminal window or a Dock icon.
2. **One-Time Permissions Check:** The first time it runs, macOS will ask `"BingWallpaper" would like to control "System Events".` You just need to click **OK** so it has permission to change your desktop.
3. **Result:** Your desktop background will update to the latest Bing image, and the app will exit immediately.

## Installation & Packaging

You can rely on the automated GitHub Actions to build the `.pkg` installer for you, or you can build it yourself using the included `Makefile`.

### 1. Build using GitHub Actions (Automated)
Once this repository is pushed to GitHub, navigate to the **Actions** tab. The automated workflow will instantly build a macOS `.pkg` installer that you can download and run. It installs `BingWallpaper.app` perfectly into your `/Applications` folder!

### 2. Build it locally using Make
If you have the Go compiler installed locally (`brew install go`), you can create the app bundle or the full `.pkg` right from your terminal:

```bash
# To just build the .app bundle that you can double click:
make app

# To generate the full installer PKG (useful for distributing to friends):
make pkg

# To clean up compiling artifacts:
make clean
```

## License
This project is officially licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
