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

// Assign Retailer ID
router.post("/assignRetailer", async function(req, res) {
  try {
    let result = await txSubmit.invoke("assignRetailer", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Sell Bag
router.post("/sellBag", async function(req, res) {
  try {
    let result = await txSubmit.invoke("assignOwner", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// P2P Resell Bag
router.post("/resellBagToCust", async function(req, res) {
  try {
    let result = await txSubmit.invoke("assignOwner", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});


// Sell Bag
router.post("/resellBag", async function(req, res) {
  try {
    let result = await txSubmit.invoke("resellBag", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/rentBag", async function(req, res) {
  try {
    let result = await txSubmit.invoke("rentBag", JSON.stringify(req.body));
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
