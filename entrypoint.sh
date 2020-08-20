cd ./go-test-web
echo "Starting Static build"
ng build
cd ..
echo "deleting Static directory"
rm -rf ./static
echo "creating Static directory"
mv ./go-test-web/dist/ang-static/ ./static/
go run main.go