# avdb-cli-go

# Set up 
> avdb-cli -s http://avdb.server.com:3000 -t XxXtestXxX group get all

## or export variables
export AVDB_TOKEN=XxXtestXxX
export AVDB_SERVER=http://avdb.server.com:3000


# Host actions
> avdb-cli host get all

> avdb-cli host get tacotruck01

> avdb-cli host get all -q food=burritos

> avdb-cli host add tacotruck01 -d $(jo bebidos=vino)

> avdb-cli host update tacotruck01 -d $(jo bebidos=jarros)

> avdb-cli host delete tacotruck01

# Group actions
> avdb-cli group get all

> avdb-cli group get foodtrucks

> avdb-cli group get all -q hosts=tacotruck01,cash=yes // for now multiple query parameters are an AND operation

> avdb-cli group add foodtrucks -data $(jo colo=las1) -hosts tacotruck01

> avdb-cli group update foodtrucks -d $(jo cash=no) -hosts tacotruck02,kimcheetruck02

> avdb-cli group delete foodtrucks

# Tag actions
> avdb-cli tag get all

> avdb-cli tag add WEB_CLUSTER web01

> avdb-cli tag get WEB_CLUSTER

> avdb-cli tag get all -q host=web01

> avdb-cli tag delete WEB_CLUSTER web01

# Updating / Removing hosts, data

## removes tacotruck01 from foodtrucks group
> avdb-cli group update foodtrucks -hosts -tacotruck01

## removes item roaches from data in hosts
> avdb-cli host update tacotruck01 roaches=
