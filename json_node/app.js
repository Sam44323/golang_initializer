const express = require("express");

app = express();

app.get("/", (req, res) => {
  res.json({
    message: "Hello from json_node server",
  });
});

app.listen(3000, () => {
  console.log("App is listening on port 3000");
});
