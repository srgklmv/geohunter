echo 'Deploy script started.';

echo $USER;
echo $HOST;

ssh -i ./key -o UserKnownHostsFile=./known_hosts $USER@$HOST << EOF

  ls;
  cd ..;
  ls;

EOF

echo 'Deploy script ended.';