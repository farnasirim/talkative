# Talkative
Talkative is a very simple web application to ease testing certain kubernetes features

### Usage
```
docker pull colonelmo/talkative
```
Run the container and 
```
curl -L $SERVER_IP/
curl -L $SERVER_IP/count
curl -L $SERVER_IP/count/reset
curl -L $SERVER_IP/count/reset?value=42
```
