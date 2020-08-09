FROM gospatial/tegola:v0.12.1

COPY config.toml /opt/tegola_config/

ENTRYPOINT [ "/opt/tegola", "serve", "--config", "/opt/tegola_config/config.toml"]
