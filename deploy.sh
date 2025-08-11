echo 'Deploy script started.';

echo $USER;
echo $HOST;

cat ./key;

ssh -i ./key -o UserKnownHostsFile=./known_hosts -T $USER@$HOST << EOF

  echo TEST

EOF

echo 'Deploy script ended.';