#!/bin/bash

go clean
go build -tags=ebitenginedebug
EBITENGINE_INTERNAL_IMAGES_KEY=P ./PixelPirates