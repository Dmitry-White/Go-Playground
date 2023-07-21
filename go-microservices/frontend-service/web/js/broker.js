const brokerBtn = document.getElementById("brokerBtn");
const authBrokerBtn = document.getElementById("authBrokerBtn");

brokerBtn.addEventListener("click", async () => {
  const { BASE, INDEX } = API.BROKER;

  const body = {
    method: "POST",
  };
  const URL = `${BASE}${INDEX}`;

  try {
    const response = await fetch(URL, body);
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

authBrokerBtn.addEventListener("click", async () => {
  const { BASE, PROCESS } = API.BROKER;

  const payload = {
    action: "Auth",
    auth: {
      email: "admin@example.com",
      password: "verysecret",
    },
  };

  const headers = new Headers({
    "Content-Type": "application/json",
  });

  const body = {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload),
  };
  const URL = `${BASE}${PROCESS}`;

  try {
    const response = await fetch(URL, body);
    const data = await response.json();

    sentBox.textContent = JSON.stringify(payload, undefined, 4);
    receivedBox.textContent = JSON.stringify(data, undefined, 4);

    if (data.error) {
      throw new Error(data.message);
    }
    outputBox.innerHTML += `<br><strong>Response from broker service</strong>: ${data.message}`;
  } catch (error) {
    outputBox.innerHTML += "<br><strong>Error</strong>:" + error;
  }
});
