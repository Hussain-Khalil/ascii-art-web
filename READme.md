# ASCII Art Web

## Description
This project provides a web interface to generate ASCII art from text using various predefined banners. Users can select different styles like 'standard', 'shadow', and 'thinkertoy'.

## Authors
- Hussain Khalil
- Hussain Malalla
- Salman Alghazal

## Usage

Run the complete project from its root directory:

```sh
go run .
```

Then open `http://localhost:8000` in a browser, enter text, choose a banner,
and select **Generate**.

## Implementation details

The `GET /` handler renders the HTML form. The form sends the text and selected
banner to `POST /ascii-art`.