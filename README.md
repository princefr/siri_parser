# SIRI Profile France Parser (Not Ready..yet)
[![License][license-badge]][6]
[![Go][go-badge]][0]
[![Python][python-badge]][1]
[![Python CI][python-test-badge]][3]
[![Golang CI][go-test-badge]][4]
[![Golang CI][rust-test-badge]][5]
#

A Rust library for parsing  SIRI (Service Interface for Real-time Information) messages according to the French national profile specification (SIRI Profile France).

## Overview
This library implements the French national adaptation of the SIRI standard, which is used for real-time public transport information exchange in France. It provides tools for parsing and working with SIRI messages that conform to the French profile specifications.


## Key Features

Full support for SIRI Profile France v2.0
XML parsing
Type-safe data structures
Support for all SIRI France service types:

- Stop Monitoring
- Estimated Timetable
- General Message
- Situation Exchange
- Production Timetable
- Facility Monitoring
- Connection Monitoring
- Vehicle Monitoring



## License
This project is licensed under either of:

Apache License, Version 2.0
MIT license


[0]: https://pkg.go.dev/github.com/princefr/siri_parser/go
[1]: https://pypi.org/project/siri-parser/
[2]: https:://crate.io
[3]: https://github.com/princefr/siri_parser/actions/workflows/python_workflow.yml
[4]: https://github.com/princefr/siri_parser/actions/workflows/CI_golang.yml
[5]: https://github.com/princefr/siri_parser/actions/workflows/CI_rust.yml
[6]: http://opensource.org/licenses/MIT


[crates-badge]: https://img.shields.io/crates/v/xmlschema.svg?style=for-the-badge 'Crates.io'
[divider]: https://raw.githubusercontent.com/sebastienrousseau/vault/main/assets/elements/divider.svg "divider"
[rust-test-badge]: https://img.shields.io/docsrs/xmlschema.svg?style=for-the-badge 'Rust CI'
[go-test-badge]: https://img.shields.io/docsrs/xmlschema.svg?style=for-the-badge 'GO CI'
[python-test-badge]: https://img.shields.io/docsrs/xmlschema.svg?style=for-the-badge 'Python CI'
[libs-badge]: https://img.shields.io/badge/lib.rs-v0.0.2-orange.svg?style=for-the-badge 'Lib.rs'
[license-badge]: https://img.shields.io/crates/l/xmlschema.svg?style=for-the-badge 'License'
[go-badge]: https://img.shields.io/crates/l/xmlschema.svg?style=for-the-badge 'Go Package'
[python-badge]: https://img.shields.io/crates/l/xmlschema.svg?style=for-the-badge 'Pypi'
[made-with-rust]: https://img.shields.io/badge/rust-f04041?style=for-the-badge&labelColor=c0282d&logo=rust 'Made With Rust'
