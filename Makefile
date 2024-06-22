FOOTBALL_BINARY=footballDataApp

football_data:
	@echo "Building binary..."
	go build -o ${FOOTBALL_BINARY} ./
	@echo "Done!"

football_data_amd:
	@echo "Building binary..."
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${FOOTBALL_BINARY} ./
	@echo "Done!"
