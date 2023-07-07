const brokerBtn = document.getElementById("brokerBtn");

brokerBtn.addEventListener("click", async () => {
  const body = {
    method: "POST",
  };

  try {
    const response = await fetch(BROKER_URL, body);
    const data = await response.json();

    sentBox.textContent = "empty post request";
    receivedBox.textContent = JSON.stringify(data, undefined, 4);

    if (data.error) {
      throw new Error(data.message);
    }
    outputBox.innerHTML += `<br><strong>Response from broker service</strong>: ${data.message}`;
  } catch (error) {
    outputBox.innerHTML += "<br><strong>Error</strong>:" + error;
  }
});
