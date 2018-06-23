# Need to use an alpine image because go-plugin requires the use of tmp file directory
FROM alpine:3.7

COPY dist/sensor-controller /

# todo: specify the directory for all the signal plugins
ENV SIGNAL_PLUGIN=/signal-plugin
COPY dist/signal-plugin /

# todo: specify the directory for all the trigger plugins
ENV TRIGGER_PLUGIN=/trigger-plugin
COPY dist/trigger-plugin /

CMD [ "/sensor-controller" ]
