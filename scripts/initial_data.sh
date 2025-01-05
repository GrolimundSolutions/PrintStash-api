#!/bin/bash

set -x # Aktiviert Debugging

HOST="http://127.0.0.1:8000"

#
## Manufacturers
#
PATH="/api/v1/manufacturers"

Die Fehlermeldung zeigt, dass curl nicht gefunden wird, obwohl es auf deinem System installiert ist. Das Problem liegt vermutlich daran, dass die Shell, in der das Skript läuft, curl nicht korrekt erkennt. Hier sind einige Lösungsansätze:

1. Vollständigen Pfad zu curl verwenden
Ändere dein Skript so, dass der vollständige Pfad zu curl genutzt wird:

bash
Code kopieren
#!/bin/bash

HOST=http://127.0.0.1:8000
PATH=/api/v1/manufacturers

/usr/bin/curl $HOST$PATH -X POST -d '{"name": "eSUN"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Prusa"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "3DFS"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Anycubic"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Bambu"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Sunlu"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "3DJake"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "PrintWithSmile"}'

