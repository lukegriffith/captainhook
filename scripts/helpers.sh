
signal() {
  ps -e | grep main | awk '{print $1}' | xargs kill $1
}

kserve(){ 
  signal -SIGTERM
}

rserve(){
  signal -SIGUSR1
}

dumpconf() {
 signal -SIGUSR2 
}



scurl(){
  curl localhost:8081/webhook/$(uuidgen) -X POST --data '{"test":1}' -H 'secret: 12341312'
}
