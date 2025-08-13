echo 'Deploy script started.';

ssh -i ./key -o UserKnownHostsFile=./known_hosts -T $USER@$HOST << EOF

  cd geohunter;
  git checkout dev;
  git pull;
  echo 'Repo pulled.';

  rm .env;
  echo '$REPO_ENV' >> .env;
  chmod u+r .env;

  docker compose up -d --build --force-recreate geohunter;
  echo 'docker compose done.';

  docker system prune -af;
  echo 'Cleanup completed.';

EOF

echo 'Deploy script ended.';