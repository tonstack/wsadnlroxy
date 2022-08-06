DOLLAR="$" envsubst < data/nginx-tmpl/app.tmpl.conf         > data/nginx/app.conf
DOLLAR="$" envsubst < data/nginx-tmpl/app-no-ssl.tmpl.conf  > data/nginx/app-no-ssl.conf
