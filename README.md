# Welcome to FormulaAPI!
FormulaAPI is a free API designed to serve data related to the FIA Formula One World Championship and its history.

### Table of Contents
[How-to](https://github.com/balinwarren/FormulaAPI/edit/issue5-update-readme/README.md#how-to-use)\
[Notice](https://github.com/balinwarren/FormulaAPI/edit/issue5-update-readme/README.md#notice)

## How to Use
FormulaAPI data is currently split into 3 categories, drivers, constructors, and circuit information. All data is returned as JSON.

**Driver Endpoints**

GET /drivers\
This endpoint returns a list of all drivers to start a race in Formula One.

GET /drivers/year/{year}\
This endpoint returns all drivers to start a race in a specified year of Formula One.

GET /drivers/name/{lastName}\
This endpoint returns all drivers with a specified last name.

GET /drivers/name/{lastName}/{firstName}\
This endpoint returns all drivers with a specified full name. Best endpoint to pull individual driver information.

## Notice
FormulaAPI is an unofficial project and is not associated in any way with the Formula 1 companies. F1, FORMULA ONE, FORMULA 1, FIA FORMULA ONE WORLD CHAMPIONSHIP, GRAND PRIX and related marks are trade marks of Formula One Licensing B.V.
