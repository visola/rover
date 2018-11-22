for d in 100 70 50 30 10 0 -10 -30 -50 -70 -100; do
    json="{\"YAxis\":$d}"
    echo $json

    curl -X PUT -d $json http://raspberrypi.local:8080/wheels
    sleep 2
done

curl -X PUT -d '{"YAxis":0}' http://raspberrypi.local:8080/wheels
