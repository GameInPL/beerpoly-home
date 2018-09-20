#!/bin/bash
set -e

cat > config/config_dev.json << EndOfConfig
{
  "mailer": {
    "smtp": {
      "addr": "$SMTP_HOST",
      "auth": {
        "username": "$SMTP_USERNAME",
        "password": "$SMTP_PASSWORD",
        "identity": "$SMTP_IDENTITY"
      }
    }
  },
  "translate": {
    "langs": "en, pl",
    "default": "en"
  },
  "router": {
    "host": ":80",
    "static": {
      "path": "./web/dist/",
      "prefix": "/static/"
    }
  },
  "database": {
    "engine": "sqlite3",
    "url": "file:main.db?cache=shared"
  },
  "app": {
    "baseURL": "$APP_BASE_URL"
  },
  "oauth": {
    "github": {
      "app": "$OAUTH_GITHUB_APP",
      "secret": "$OAUTH_GITHUB_SECRET"
    }
  },
  "menu": {
    "OpenSourceLink": "$MENU_OPEN_SOURCE_LINK",
    "GameLink": "$MENU_GAME_LINK",
    "DownloadLink": "$MENU_DOWNLOAD_LINK"
  }
}
EndOfConfig

cp config/config_dev.json config/config_prod.json
cp config/config_dev.json config/config_test.json

### Build database
rm main.db ||:
./bphome dbbuild

### Add secrets
for i in `env | grep -E "^SLOT_.*_SECRET"`
do
    name=$(echo $i| cut -d'_' -f 2)
    eval value=\${f$i}
    ./bphome slot:secrets:add --secret=$value --slot=$name
    echo "Secret $value for $name slot... added "
done

### Add users
for i in `env | grep -E "^USER_.*_USERNAME="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval firstname=\${${baseKey}FIRSTNAME}
    eval lastname=\${${baseKey}LASTNAME}
    ./bphome user:add --interactive=false --username=$username --email=$email --firstname=$firstname --lastname=$lastname
    echo "User $username ($email)... added "
done

### Update users roles
for i in `env | grep -E "^USER_.*_ROLES="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval roles=\${${baseKey}ROLES}
    ./bphome user:roles:update --by=$email --roles=$roles
    echo "User $username ($email)... roles updated"
done

### Update users passwords
for i in `env | grep -E "^USER_.*_PASSWORD="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval password=\${${baseKey}PASSWORD}
    ./bphome user:password:update --by=$username --password=$password
    echo "User $username ($email)... password updated"
done

### Connect user to remote accounts
for i in `env | grep -E "^USER_.*_CONNECT_.*="`
do
    userBaseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    connectBaseKey="USER_$(echo $i| cut -d'_' -f 2)_CONNECT_$(echo $i| cut -d'_' -f 4)_"
    service="$(echo $i| cut -d'_' -f 4 | cut -d'=' -f 1 | tr '[:upper:]' '[:lower:]')"
    eval username=\${${userBaseKey}USERNAME}
    eval email=\${${userBaseKey}EMAIL}
    eval remote=\${${i}}
    ./bphome user:connected:add --local=$email --remote=$remote --service=$service
    echo "Connect $username ($email) user to $service remote service... success"
done

### RUN
./bphome run
