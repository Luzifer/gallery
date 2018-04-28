[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/gallery)](https://goreportcard.com/report/github.com/Luzifer/gallery)
![](https://badges.fyi/github/license/Luzifer/gallery)
![](https://badges.fyi/github/downloads/Luzifer/gallery)
![](https://badges.fyi/github/latest-release/Luzifer/gallery)

# Luzifer / gallery

`gallery` is a pure HTML/Javascript gallery viewer with an uploader written in Go. It is intended to run inside S3/CloudFront or other static web hosting scenarios: The uploader resizes the original image to fit given image constraints and also creates a thumbnail image. Afterwards an `albums.json` file is written and uploaded with the frontend files and the images.

## Setup

- Create a S3 bucket
- Grant global read permission using a bucket policy
- Put the images for your album into a folder
- Upload them:
    ```console
    $ gallery --storage s3://io-luzifer-photos --update-frontend --album-id=2012020601 --album-title 'Frozen river' ./frozen-river
     28 / 28 [==============================] 100.00% 36s
    ```

### Storage engines

- S3 (`s3://bucket-name`)  
  Upload everything into the root of the bucket
- File (`file://path`)  
  Upload everything into the given folder (`file:///tmp/gallery` would "upload" to `/tmp/gallery/...`)
