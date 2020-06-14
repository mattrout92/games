yarn build
aws s3 sync build/ s3://games-ui
s3cmd modify --add-header="Cache-Control:no-cache" s3://games-ui/index.html
