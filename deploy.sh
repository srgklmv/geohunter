echo 'Deploy script started.';

echo $USER;
echo $REMOTE_ADDR;

ssh -i ./key -o UserKnownHostsFile=/github/home/.ssh/known_hosts $USER@$REMOTE_ADDR << EOF

  ls;
  cd ..;
  ls;

EOF

echo 'Deploy script ended.';