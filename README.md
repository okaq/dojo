# dojo

Dojo, jp. "Sympathy, Compassion"

aomi - grid kernel methods

bomi - prototype terrain generator

noko - favicons
nokoa/ - json and png favicon store to disk

oko - import bitmap sample and encode
okoa/ - requires directories for png and json

kira - previous bitmap sample and encode incarnation
kira.go - server side file process, upload and disk store
kira.html - frontend sample canvas
kirad/ - input png 1024x1024 file dir
kirap/ - old png files already processed
kiraj/ - output json bitmap base64 binary encoded string

Pipeline for bitmap sample and encode
- store original 1024x1024 monochrome PNG files in kirad/
- go run kira.go
- open http://localhost:8080 in chrome
- browser canvas fetches file list
- rendering images and sampling pixel data
- uploads data to file handler on server
- finally stitching and saving to single json file in kiraj/
- post: manually move PNGs to kirap/ 

Naming convention for bitmap sample and encode
kirad/dragon_1.png -> kiraj/dragons.json


