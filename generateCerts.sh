rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile VuittonOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile VuittonOrgChannel -outputCreateChannelTx ./config/vuittonchannel.tx -channelID vuittonchannel
