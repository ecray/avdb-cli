# avdb-cli-go

# Host actions
> avdb-cli host get all

> avdb-cli host get infdcpdn01

> avdb-cli host get all -q colo=las1

> avdb-cli host add infdcpdns01 -d $(jo colo=las1)

> avdb-cli host update infdcpdns01 -d $(jo colo=aws)

> avdb-cli host delete infdcpdns01

# Group actions
> avdb-cli group get all

> avdb-cli group get las1-adm

> avdb-cli group get all -q host hosts=infdcpdns01

> avdb-cli group add infdcpdns01 -data $(jo colo=las1) -hosts infdcpdns01

> avdb-cli group update infdcpdns01 -d $(jo colo=aws) -hosts lvopsdcadm01

> avdb-cli group delete las1-adm

# Updating / Removing hosts, data

## removes infdcpdns01 from las10-adm group
> avdb-cli group update las1-adm -hosts -infdcpdns01

## removes item trivial from data in hosts
> avdb-cli host update infdcpdns01 trivial=
