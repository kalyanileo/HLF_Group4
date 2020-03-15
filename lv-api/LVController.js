var express = require("express");
var router = express.Router();
var bodyParser = require("body-parser");

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

const txSubmit = require("./invoke");
const txFetch = require("./query");

//var TFBC = require("./FabricHelper");

// Create Bag 
router.post("/createBag", async function(req, res) {
  try {
    let result = await txSubmit.invoke("createBag", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Assign distributor ID 
router.post("/assignDistributor", async function(req, res) {
  try {
    let result = await txSubmit.invoke("assignDistributor", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});


// Get Bag 
router.post("/getBag", async function(req, res) {
  //TFBC.getLC(req, res); req.body.lcId
  try {
    let result = await txFetch.query("getBag", req.body.tagId);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// Get Bag History
router.post("/getBagHistory", async function(req, res) {
  try {
    let result = await txFetch.query("getBagHistory", req.body.tagId);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
