build:
	echo "Build backend"
	cp -R resources dist
	go build -o dist/notans
	echo "Kindly check dist folder, and copy all files on that folder"

run:
	echo "Running server....."
	cd dist && ./notans