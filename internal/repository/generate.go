package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ChatRepo -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i MessageRepo -o ./mocks/ -s "_minimock.go"
