
var ws = new WebSocket("ws://localhost:8080/echo");

ws.onopen = function() {
  // Web Socket is connected, send data using send()
  ws.send("Message to send");
  console.log("Message is sent...");
};

ws.onmessage = function (evt) {
  var received_msg = evt.data;
  console.log("Message is received...");
  console.log(received_msg);
};
