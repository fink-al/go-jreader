# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).


## [v1.0.0] - 2024-02-12

### Added

- Initial release
- Add JSONElement interface (Get, Value, BooleanValue, NumberValue, StringValue methods; see common.go for details)
- Add Load(data JSONData) (JSONElement, error) function
- Add jSONMap, jSONSlice, jSONMapSlice, jSONValue, nonExistent concrete types
