![chamber of understanding](https://i.imgur.com/sVtZKfo.gif)

## Architecture

#### Adapters
Adapters are client side that track & apply changes to a specific database.


#### Core
Core is a client side manager that will communicate with the sync engine and apply any changes to a database.


#### Sync Engine
Sync Engine communicates with Core, manages changesets to create parity across all databases, and broadcast updates.

## TODO
- [ ] Create Adapter for SQLite 
  - [ ] Implement `TrackChanges`
  - [ ] Implement `ApplyChanges`
