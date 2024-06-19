# 复制node.ini 文件到
\cp ./gvb_server.ini /etc/supervisord.d

cd ../

./main -esload init/data/article_index_20240511.json
./main -esload init/data/full_text_index_20240511.json
supervisorctl reload