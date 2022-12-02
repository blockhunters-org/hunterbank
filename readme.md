# hunterbank
**hunterbank** is a lending blockchain built using Cosmos SDK and Tendermint. It can be used for borrowing and lending assets for the Game of Chains testnets.

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
hunterbankd init $MONIKER
```
 
## Available commands
```
Usage:
  hunterbankd tx loan [flags]
  hunterbankd tx loan [command]

Available Commands:
  request-loan      Request a new loan using GoC assets as collateral
  approve-loan      Provide requested assets for a commission
  repay-loan        Repay an outstanding loan
  liquidate-loan    Liquidate a loan that is past it's deadline
  cancel-loan       Cancel a loan request
  
Usage:
  hunterbankd query loan [flags]
  hunterbankd query loan [command]

Available Commands:
  list-loan         Show existing loans

```


### Request a new loan

```
Use:   "request-loan [amount-requested] [loan-fee] [collateral] [deadline-blocks]"
Example: "hunterbankd tx loan request-loan 1000prov 5prov 500bucks 1000 --from alice"
```

### Approve loan

```
Use:   "approve-loan [loan-id]"
Example: "hunterbankd tx loan approve-loan 11 --from alice"
``` 

### Repay loan

```
Use:   "repay-loan [loan-id]"
Example: "hunterbankd tx loan repay-loan 11 --from bob"
``` 

### Liquidate loan

```
Use:   "liquidate-loan [loan-id]"
Example: "hunterbankd tx loan liquidate-loan 11 --from alice"
``` 

### Cancel loan

```
Use:   "cancel-loan [loan-id]"
Example: "hunterbankd tx loan cancel-loan 11 --from bob"
``` 

### List existing loans

```
Use:   "list-loans"
Example: "hunterbankd query list-loan"
``` 

 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.hunterbank/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.hunterbank/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
hunterbankd add-consumer-section
```

## Launch node

```
hunterbankd start
```
