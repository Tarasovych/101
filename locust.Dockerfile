FROM locustio/locust:2.4.1

COPY locustfile.py /mnt/locust/locustfile.py

CMD ["-f", "/mnt/locust/locustfile.py"]
