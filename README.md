# 🖼️ ImageColors API

**ImageColors** is a simple and fast Go-based API that accepts a base64-encoded image and returns the RGBA color values of specified pixels. It's designed for image analysis, debugging, or educational purposes.

---

## 🚀 Features

- Accepts base64-encoded image input
- Receives a list of pixel coordinates (e.g., `"100,200"`)
- Returns the RGBA values for each pixel
- Lightweight and easy to deploy
- Written in idiomatic Go

---

## 📦 API Usage

### `POST /getcolors`

**Request Body:**

```json
{
  "pixels": ["100,200", "300,400"],
  "base64image": "iVBORw0K..."
}
