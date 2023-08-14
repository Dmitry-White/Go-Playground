const buttons = document.getElementById("buttons");

const { BASE, PROCESS } = API.BROKER;

const SERVICES = {
  brokerBtn: {
    url: `${BASE}`,
    payload: {},
  },
  authBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "Auth",
      auth: {
        email: "admin@example.com",
        password: "verysecret",
      },
    },
  },
  logBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "Log",
      log: {
        name: "event",
        data: "some kind of data",
      },
    },
  },
  mailBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "Mail",
      mail: {
        from: "testFrom@example.com",
        to: "testTo@example.com",
        subject: "Test Subject",
        message: "Test Message",
      },
    },
  },
  asyncBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "Async",
      async: {
        name: "Async Test Name",
        data: "Async Test Data",
      },
    },
  },
  rpcBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "LogRPC",
      logRPC: {
        name: "RPC Test Name",
        data: "RPC Test Data",
      },
    },
  },
  grpcBrokerBtn: {
    url: `${BASE}${PROCESS}`,
    payload: {
      action: "LogGRPC",
      logGRPC: {
        name: "GRPC Test Name",
        data: "GRPC Test Data",
      },
    },
  },
};

const makeRequest = async (service) => {
  const { url, payload } = service;

  const headers = new Headers({
    "Content-Type": "application/json",
  });

  const body = {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload),
  };

  try {
    const response = await fetch(url, body);
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
};

buttons.addEventListener("click", async (event) => {
  const target = event.target;
  console.log("Target: ", target.id);

  const service = SERVICES[target.id];

  if (!service) {
    return;
  }

  try {
    await makeRequest(service);
  } catch (error) {
    console.log("Error: ", error);
  }
});
