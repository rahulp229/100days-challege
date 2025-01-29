build : rm day4-app1-inventory-app-for-grocery-store go build -race   ./day4-app1-inventory-app-for-grocery-store
newapp : mkdir $1 && touch main.go && go mod init $1