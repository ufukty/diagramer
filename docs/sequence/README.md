# Sequence Diagrams

## Syntax

### Definition

```bnf
DIAGRAM         ::= "sequenceDiagram" ["autoNumber"] { LIFELINE | MESSAGE }

NAME            ::= WORD

LIFELINE        ::= ( "participant" | "actor" ) NAME [ "as" ALIAS ]
ALIAS           ::= (WORD | DIGIT | SPACE)+

MESSAGE         ::= NAME "->>" NAME [ ":" MESSAGE_CONTENT ]
MESSAGE_CONTENT ::= (WORD | DIGIT | SPACE)+
```

### Example

```txt
sequenceDiagram
autoNumber

participant db as Server Database
participant rd as Server Redis
participant ep as Server Endpoint
participant a as Client App
participant ds as Client DataSource
participant ui as Client UI
actor u as User

u->>a: opens the website
a->>ui: initializes
a->>ds: initializes
ds->>ep: creates a socket, subscribes to event stream

u->>ui: moves a "task"
ui->>ds: notifies change
ds->>ep: notifies change
ep->>db: update table(s)
ep->>rd: invalidates cache(s)

ep->>ds: pushes notification: "new placement", "new parent (?)", "new order (?)"
ds->>ep: asks for new placement / task details (if it is still necessary / in-view)
ep->>db: compute placement
ep->>rd: save to redis
ep->>ds: pushes data: "new placement"

ds->>ui: updates config
ui->>ui: diffs configs
ui->>ui: updates html, if necessary
```
