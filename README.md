# Bing Wallpaper for macOS 🌅

Hey! This is a simple, lightweight macOS app I wrote in Go. It does one thing and does it well: it grabs the daily Bing Wallpaper and sets it as your desktop background so you have something fresh to look at every day.

> [!NOTE]  
> To be honest, I mostly built this as a fun way to play around with Go and see what I could do with GitHub Actions. It's been a cool little project!

## What it does

- **Native & Light:** It's a proper macOS `.app` bundle, but under the hood, it's all Go.
- **Invisible:** It runs quietly in the background. No messy terminal windows or annoying Dock icons.
- **High Quality:** It pulls the high-res (`_UHD.jpg`) images directly from the Bing API, so they look great even on Retina displays.

## How it works
When you launch the app, it checks the Bing Image Archive. If you don't already have today's wallpaper saved in `~/Pictures/BingWallpapers`, it'll download it and use macOS System Events to refresh your desktop background across all your spaces.

## Using the app
I designed this to be "set it and forget it." Since it's a background process (`LSUIElement`), it won't clutter up your Dock.

1. **Permissions:** The very first time you run it, macOS will ask if `BingWallpaper` can control `System Events`. Just click **OK**—that's how it gets the permission it needs to change your wallpaper.
2. **Login Item:** The best way to use it is to add it as a login item. That way, every time you boot up your Mac, your wallpaper is already updated and waiting for you.
3. **Outcome:** Your desktop gets a beautiful new look, and the app closes itself immediately. Simple as that.

## Installation & Packaging

You can either let GitHub Actions do the heavy lifting or build it yourself if you're feeling hands-on.

### 1. Use the pre-built installer (The Easy Way)
Whenever I push changes, GitHub Actions automatically builds a `.pkg` installer. You can find it under the **Actions** tab. Just download it, run it, and it'll drop `BingWallpaper.app` right into your `/Applications` folder.

### 2. Build it yourself (The Dev Way)
If you've got Go installed (`brew install go`), you can build the app or the installer package directly from your terminal using the `Makefile`:

```bash
# To build just the .app bundle:
make app

# To generate the full .pkg installer:
make pkg

# To clean up the build files:
make clean
```

## License
This project is licensed under the MIT License—feel free to check out the [LICENSE](LICENSE) file if you're curious about the details.
