version="1.0.0-alpha"

echo "Start Production $version"
sudo go build ./cmd/main/main.go
echo "Build AresBot for MacOs"
mv ./main ./AresBot
mv ./AresBot ./build/aresbot_mac/
echo "Zipping AresBot for MacOs"
zip -r -X AresBotMac_$version.zip ./build/aresbot_mac/
sudo GOOS=windows GOARCH=amd64 go build ./cmd/main/main.go
echo "Build AresBot for Windows"
mv ./main.exe ./AresBot.exe
mv ./AresBot.exe ./build/aresbot_windows/
echo "Zipping AresBot for Windows"
zip -r -X AresBotWindows_$version.zip ./build/aresbot_windows/
echo "Finish Zipping"
echo "Clearing Build Dir"
rm -rf ./build/aresbot_mac/FirestormBot
rm -rf ./build/aresbot_windows/FirestormBot.exe

echo "Upload ready zip directorys and cleand up"