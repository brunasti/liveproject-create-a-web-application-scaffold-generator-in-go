go build -o scaffold_gen

echo "Create Handler application"
./scaffold_gen -d ../temp/test_handler -n TestHandler -p 8081 -r repository

echo "Run Handler application"
cd ../temp/test_handler/handlers
go build -o handler_app
./handler_app &
sleep 2
cd ../../../liveproject-create-a-web-application-scaffold-generator-in-go


echo "TEST ----"
echo curl localhost:8081/healthcheck
curl localhost:8081/healthcheck
echo ""
echo "----"

echo curl localhost:8081
curl localhost:8081
echo ""
echo "----"


echo "-------------------------------------------"
echo "Create Static application"
./scaffold_gen -d ../temp/test_static -n TestStatic -p 8082 -r repository -s

echo "Run Static application"
cd ../temp/test_static/handlers
go build -o static_app
./static_app &
sleep 2
cd ../../liveproject-create-a-web-application-scaffold-generator-in-go

echo "TEST ----"
echo curl localhost:8082/healthcheck
curl localhost:8082/healthcheck
echo ""
echo "----"

echo curl localhost:8082
curl localhost:8082
echo ""
echo "----"


ps -ef | grep "./handler_app"
echo "----"
ps -ef | grep "./static_app"
echo "----"
