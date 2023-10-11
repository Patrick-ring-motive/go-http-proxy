cd api
go build -o handler vercel.go
chmod +x ./handler
cd main
go build -o indexgo index.go
chmod +x ./indexgo
cd ../..