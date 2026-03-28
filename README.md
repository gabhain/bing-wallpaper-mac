# Bing Wallpaper for macOS 🌅

A lightweight, completely native macOS application written in Go that automatically fetches the daily Ultra High-Resolution (UHD) Bing Wallpaper, saves it, and seamlessly sets it as your desktop background without dropping a single icon in your dock!

## Features

- **Native macOS Execution:** The compiled Go binary is packaged cleanly into a `.app` bundle.
- **Headless Action:** Defined as an `LSUIElement`, clicking the app executes the logic completely invisibly in the background without popping up terminals or Dock icons.
- **Highest Quality:** Automatically grabs the `_UHD.jpg` endpoint straight from the Microsoft Image API for maximum crispness.

## How it works
On launch, the tool pulls the JSON payload from the Bing Image Archive API. It checks if the wallpaper has already been downloaded to `~/Pictures/BingWallpapers` today. If not, it saves it cleanly and uses macOS `System Events` (via `osascript`) to set it across all active desktop spaces.

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
