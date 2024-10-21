
# SIRI Profile France Parser (Rust) - Proof of concept

## Overview

The SIRI (Service Interface for Real-time Information) Profile France Parser is a Rust project aimed at creating a universal parser for the SIRI spécification, specifically tailored to the French public transport context. 
This parser is designed to be the "One parser to rule them all," capable of handling various aspects of the SIRI standard efficiently and comprehensively.

We carefully followed the SIRI France spécification here: [SIRI](https://normes.transport.data.gouv.fr/normes/siri/profil-france/)

## Purpose

The main goal of this parser is to provide a robust, efficient, and flexible solution for processing SIRI data in the French public transport ecosystem. 
By creating a unified parser, it aims to simplify the integration of real-time public transport information systems across different operators and regions in France.

## Services implemented

- Production Timetable
- Estimated Timetable
- Stop Monitoring
- Vehicle Monitoring
- Connection Monitoring
- General Message
- Facility Monitoring
- Situation Exchange

## Notification Specifities

This parser can only read SIRI notification that are inside XML  Envelope and Body like shown below.


```xml

<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
            <!-- Notifications go here -->
            <!-- NotifyStopMonitoring
            siri:NotifyEstimatedTimetable -->
            <!-- ETC... -->
   </soapenv:Body>
</soapenv:Envelope>

```


## Folder structure

We followed SIRI (Profil France) specification for our folder structures. 
Structures are held in a structures folder.
Enums are held in a enums folder.
Services offered by the parser are held in services folder.

---siri_parser
---------src
-------------enums
--------------- actions.rs
--------------- arrival_status.rs
--------------- etc...
-------------services
--------------- estimated_table_delivery.rs
--------------- stop_monitoring_delivery.rs
--------------- etc...
-------------structures
--------------- xxx_delivery.rs
--------------- estimated_calls.rs
--------------- recorded_calls.rs

Examples :

You can get the envelope of a xml file or a xml string by respectively calling from_file or from_str functions of the Envelope struct

from_file example: 

```rust

use parser::{Envelope, SIRI};

let siri = SIRI::from_file::<Envelope>("src/fixtures/siri_sm/siri-destineo-sm-cus-2-2024-04-04-13-02-25.xml");
match siri {
    Ok(siri) => {
        match siri.body.0 {
            SiriServiceType::EstimatedTimetable(siri) => println!("XML: {:?}", siri),
            SiriServiceType::StopMonitoring(siri) => println!("XML: {:?}", siri),
            _ => println!("XML: {:?}", siri),
        }
    },
    Err(e) => println!("Error: {:?}", e),
}

```

Envelope body can be from different types, it follow the available SIRI available service type


### 2. French Profile Specialization

- This repo is tailored to handle France-specific extensions and modifications to the SIRI standard
- Supports French-specific data formats and conventions


### 3. High Performance

- Utilizes Rust's performance benefits for fast and efficient parsing
- Designed with concurrency in mind to handle large volumes of real-time data



