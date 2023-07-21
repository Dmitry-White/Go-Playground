const outputBox = document.getElementById("output");
const sentBox = document.getElementById("payload");
const receivedBox = document.getElementById("received");

const API = {
    BROKER: {
        BASE: "http://localhost:8081",
        INDEX: "/",
        PROCESS: "/process"
    }
}

