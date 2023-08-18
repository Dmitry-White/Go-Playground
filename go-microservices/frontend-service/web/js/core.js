const outputBox = document.getElementById("output");
const sentBox = document.getElementById("payload");
const receivedBox = document.getElementById("received");

const MODES = {
  local: "http://localhost:8081",
  swarm: "http://backend:80",
};

const API = {
  BROKER: {
    BASE: MODES[mode],
    INDEX: "/",
    PROCESS: "/process",
  },
};

console.log(MODES[mode]);
