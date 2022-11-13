# teleinfo-timescaledb

🇬🇧🇬🇧🇬🇧🇬🇧
Info: As this repository will only work with French Electricity providers, the README is in French. Sorry for that.
Note that code base is in English.
🇬🇧🇬🇧🇬🇧 🇬🇧


## How to dev

```bash
docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg14-latest
```

## Data

Données sauvegardées en brut: 

- EAST (Energie active soutirée totale): Wh => total_active_energy
- PCOUP (Puissance apparente de coupure): kVA => apparent_breaking_power 
- SINSTS (Puissance apparente instantanée soutirée): VA => instant_apparent_power
- SINSTI (Puissance apparente instantanée injectée): VA => instant_apparent_power_injected

Donée à calculer:

- Puissance active Instantanée: Différence entre deux index qui se suivent au prorata du temps. Ne disposant pas du déphasage entre U et I il est impossible de le caculer autrement.


Données recupérées:

Adresse Secondaire du Compteur: ADSC
Version de la TIC: VTIC
Date et heure courante: DATE
Nom du calendrier tarifaire fournisseur: NGTF
Libellé tarif fournisseur en cours: LTARF
Energie active soutirée totale: EAST (Unité: Wh)
Energie active soutirée Fournisseur, index 01: EASF01 (Unité: Wh)
Energie active soutirée Fournisseur, index 02: EASF02 (Unité: Wh)
Energie active soutirée Fournisseur, index 03: EASF03 (Unité: Wh)
Energie active soutirée Fournisseur, index 04: EASF04 (Unité: Wh)
Energie active soutirée Fournisseur, index 05: EASF05 (Unité: Wh)
Energie active soutirée Fournisseur, index 06: EASF06 (Unité: Wh)
Energie active soutirée Fournisseur, index 07: EASF07 (Unité: Wh)
Energie active soutirée Fournisseur, index 08: EASF08 (Unité: Wh)
Energie active soutirée Fournisseur, index 09: EASF09 (Unité: Wh)
Energie active soutirée Fournisseur, index 10: EASF10 (Unité: Wh)
Energie active soutirée Distributeur, index 01: EASD01 (Unité: Wh)
Energie active soutirée Distributeur, index 02: EASD02 (Unité: Wh)
Energie active soutirée Distributeur, index 03: EASD03 (Unité: Wh)
Energie active soutirée Distributeur, index 04: EASD04 (Unité: Wh)
Energie active injectée totale: EAIT (Unité Wh)
Energie réactive Q1 totale: ERQ1 (Unité: VArh)
Energie réactive Q2 totale: ERQ2 (Unité: VArh)
Energie réactive Q3 totale: ERQ3 (Unité: VArh)
Energie réactive Q4 totale: ERQ4 (Unité: VArh)
Courant efficace, phase 1: IRMS1 (Unité: A)
Courant efficace, phase 2: IRMS2 (Unité: A)
Courant efficace, phase 3: IRMS3 (Unité: A)
Tension efficace, phase 1: URMS1 (Unité: V)
Tension efficace, phase 2: URMS2 (Unité: V)
Tension efficace, phase 3: URMS3 (Unité: V)
Puissance app. de référence: PREF (Unité: kVA)
Puissance app. de coupure: PCOUP (Unité: kVA)
Puissance app. Instantanée soutirée: SINSTS (Unité: VA)
Puissance app. Instantanée soutirée phase 1: SINSTS1 (Unité: VA)
Puissance app. instantanée soutirée phase 2: SINSTS2 (Unité: VA)
Puissance app. instantanée soutirée phase 3: SINSTS3 (Unité: VA)
Puissance app. max. soutirée: SMAXSN (Unité: VA)
Puissance app. max. soutirée n phase 1: SMAXSN1 (Unité: VA)
Puissance app. max. soutirée n phase 2: SMAXSN2 (Unité: VA)
Puissance app. max. soutirée n phase 3: SMAXSN3 (Unité: VA)
Puissance app max. soutirée n-1: SMAXSN-1 (Unité: VA)
Puissance app max. soutirée n-1 phase 1 SMAXSN1-1 (Unité: VA)
Puissance app max. soutirée n-1 phase 2 SMAXSN2-1 (Unité: VA)
Puissance app max. soutirée n-1 phase 3 SMAXSN3-1 (Unité: VA)
Puissance app. Instantanée injectée: SINSTI (Unité: VA)
Puissance app. max. injectée n: SMAXIN (Unité: VA)
Puissance app max. injectée n-1: SMAXIN-1 (Unité: VA)
Point n de la courbe de charge active soutirée: CCASN (Unité: W)
Point n-1 de la courbe de charge active soutirée: CCASN-1 (Unité: W)
Point n de la courbe de charge active injectée: CCAIN (Unité: W)
Point n-1 de la courbe de charge active injectée: CCAIN-1 (Unité: W)
Tension moy. ph. 1: UMOY1 (Unité: V)
Tension moy. ph. 2: UMOY2 (Unité: V)
Tension moy. ph. 3: UMOY3 (Unité: V)
Registre de Statuts: STGE
Début Pointe Mobile 1: DPM1
Fin Pointe Mobile 1: FPM1
Début Pointe Mobile 2: DPM2
Fin Pointe Mobile 2: FPM2
Début Pointe Mobile 3: DPM3
Fin Pointe Mobile 3: FPM3
Message court: MSG1 
Message Ultra court: MSG2
PRM: PRM
Relais: RELAIS
Numéro de l’index tarifaire en cours: NTARF
Numéro du jour en courscalendrier fournisseur: NJOURF
Numéro du prochain jour calendrier fournisseur: NJOURF+1
Profil du prochain jour calendrier fournisseur: PJOURF+1
Profil du prochain jour de pointe: PPOINTE