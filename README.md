# motdservice
MOTD Service

docker build -t adamgilmore/motdservice .
docker tag adamgilmore/motdservice:latest adamgilmore/motdservice:1.2
docker push adamgilmore/motdservice:1.2

docker run --name motdservice -p 127.0.0.1:8091:8091 -it adamgilmore/motdservice

