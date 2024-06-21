#!/usr/bin/bash
if ./yes-no.sh "initializae a new database and user account"
then
sudo -u postgres createuser $USER
sudo -u postgres createdb $USER
echo You will need to set a password as some tools do not allow blank database passwords.
psql -c "\password"
echo The database will now be built.
sudo -u postgres psql -c "grant all privileges on database $USER to $USER;"
fi
echo "Filling database"
psql < schema.sql
psql < data.sql
