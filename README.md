# teleinfo-timescaledb


## How to dev

```bash
docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg14-latest
```

## Test commands

```
curl -d '{"ISOUSC":"9", "BASE":"1234", "IINST": "6", "IMAX": "9"}' -H "Content-Type: application/json" -X POST http://localhost:4000/v1/frames
```


## Data

EAST (Energie Active soutirée totale): Linky HP/HC  (/1000 pour avoir en KWH)

PCOUP (Puissance apparente de coupure): kVA

SINSTS (Puissance apparente instantanée): W 

SINSTI (Puissance app. Instantanée injectée): W