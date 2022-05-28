const express = require("express");

app = express();

app.get("/", (req, res) => {
  res.json({
    message: "Hello from json_node server",
  });
});

app.post("/post", (req, res) => {
  let customJson = res.body;

  res.status(200).json(customJson);
});

app.post("/postForm", (req, res) => {
  res.status(200).json(JSON.stringify(res.body));
});

app.listen(3000, () => {
  console.log("App is listening on port 3000");
});
