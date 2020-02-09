echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=LVMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lv.vuitton.com/users/Admin@lv.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.lv.vuitton.com:7051" cli peer channel create -o orderer.vuitton.com:7050 -c vuittonchannel -f /etc/hyperledger/configtx/vuittonchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.lv.vuitton.com joining the channel..."
# Join peer0.lv.vuitton.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=LVMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lv.vuitton.com/users/Admin@lv.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.lv.vuitton.com:7051" cli peer channel join -b vuittonchannel.block

echo "peer0.lv.vuitton.com joined the channel"

echo "peer0.distributor.vuitton.com joining the channel..."

# Join peer0.distributor.vuitton.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=DistributorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/distributor.vuitton.com/users/Admin@distributor.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.distributor.vuitton.com:7051" cli peer channel join -b vuittonchannel.block

echo "peer0.distributor.vuitton.com joined the channel"

echo "peer0.retailer.vuitton.com joining the channel..."
# Join peer0.retailer.vuitton.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=RetailerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/retailer.vuitton.com/users/Admin@retailer.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.retailer.vuitton.com:7051" cli peer channel join -b vuittonchannel.block
sleep 5

echo "peer0.retailer.vuitton.com joined the channel"

echo "Installing vuitton chaincode to peer0.lv.vuitton.com..."

# install chaincode
# Install code on lv peer
docker exec -e "CORE_PEER_LOCALMSPID=LVMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lv.vuitton.com/users/Admin@lv.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.lv.vuitton.com:7051" cli peer chaincode install -n vuittoncc -v 1.0 -p github.com/vuitton/go -l golang

echo "Installed vuitton chaincode to peer0.lv.vuitton.com"

echo "Installing vuitton chaincode to peer0.distributor.vuitton.com...."

# Install code on distributor peer
docker exec -e "CORE_PEER_LOCALMSPID=DistributorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/distributor.vuitton.com/users/Admin@distributor.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.distributor.vuitton.com:7051" cli peer chaincode install -n vuittoncc -v 1.0 -p github.com/vuitton/go -l golang

echo "Installed vuitton chaincode to peer0.distributor.vuitton.com"

echo "Installing vuitton chaincode to peer0.retailer.vuitton.com..."
# Install code on retailer peer
docker exec -e "CORE_PEER_LOCALMSPID=RetailerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/retailer.vuitton.com/users/Admin@retailer.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.retailer.vuitton.com:7051" cli peer chaincode install -n vuittoncc -v 1.0 -p github.com/vuitton/go -l golang

sleep 5

echo "Installed vuitton chaincode to peer0.distributor.vuitton.com"

echo "Instantiating vuitton chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=LVMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lv.vuitton.com/users/Admin@lv.vuitton.com/msp" -e "CORE_PEER_ADDRESS=peer0.lv.vuitton.com:7051" cli peer chaincode instantiate -o orderer.vuitton.com:7050 -C vuittonchannel -n vuittoncc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('LVMSP.member','DistributorMSP.member','RetailerMSP.member')"

echo "Instantiated vuitton chaincode."

echo "Following is the docker network....."

