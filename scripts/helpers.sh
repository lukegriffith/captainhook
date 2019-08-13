kserve(){ 
  kill -SIGTERM $(ps -a | grep standalone | awk '{print $1}')
}

scurl(){
  curl localhost:8081/webhook/$(uuidgen) -X POST --data '{"test":1}' -H 'secret: 12341312'
}
