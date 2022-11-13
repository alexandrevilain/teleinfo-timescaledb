# teleinfo-timescaledb

🇬🇧🇬🇧🇬🇧🇬🇧

Info: As this repository will only work with French Electricity providers, the README is in French. Sorry for that.

Note that code base is in English.

🇬🇧🇬🇧🇬🇧 🇬🇧


## Comment développer ?

### Coté serveur

```bash
docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg14-latest
go run ./cmd/server --config=config.example.yaml
```

### Coté client

Le client doit être connecté au module teleinfo du linky via USB.
```
go run ./cmd/client
```


## Données utilisées

Données sauvegardées en brut: 

- EAST (Energie active soutirée totale): Wh => total_active_energy
- PCOUP (Puissance apparente de coupure): kVA => apparent_breaking_power 
- SINSTS (Puissance apparente instantanée soutirée): VA => instant_apparent_power
- SINSTI (Puissance apparente instantanée injectée): VA => instant_apparent_power_injected

Donée à calculer:

- Puissance active Instantanée: Différence entre deux index qui se suivent au prorata du temps. Ne disposant pas du déphasage entre U et I il est impossible de le caculer autrement.