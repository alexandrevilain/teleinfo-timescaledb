# teleinfo-timescaledb

🇬🇧🇬🇧🇬🇧🇬🇧

Info: As this repository will only work with French Electricity providers, the README is in French. Sorry for that.

Note that code base is in English.

🇬🇧🇬🇧🇬🇧 🇬🇧

## Comment ça marche ?

Ce projet fournis 2 binaires:
- teleinfo-client 
- teleinfo-server

Le binaire `teleinfo-client` est fait pour tourner sur un RaspberryPi (ou n'importe quelle autre machine linux). Il doit être relié a un module teleinfo USB (comme [celui-ci](https://www.cartelectronic.fr/teleinfo-compteur-enedis/127-teleinfo-1-compteur-usb-lc.html)). 
Le binaire `teleinfo-server` est fait pour tourner coté serveur et doit être connecté a une base de donnée postgres avec l'extension timescaledb.
Le `teleinfo-client` envoie via HTTP les trames teleinfo au `teleinfo-server`.

HTTP a été choisi pour cette v1 mais sera remplacé par MQTT dans le futur.


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
- SINSTS (Puissance apparente instantanée soutirée): VA => instant_apparent_power
- SINSTI (Puissance apparente instantanée injectée): VA => instant_apparent_power_injected

Donée à calculer:

- Puissance active Instantanée: Différence entre deux index qui se suivent au prorata du temps. Ne disposant pas du déphasage entre U et I il est impossible de le caculer autrement.
