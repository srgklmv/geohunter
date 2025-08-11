echo 'Deploy script started.';

ssh -i ./key $USER@$REMOTE_ADDR << EOF

  ls;
  cd ..;
  ls;

EOF

echo 'Deploy script ended.';