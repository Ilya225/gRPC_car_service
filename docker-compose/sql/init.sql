CREATE TABLE car
(
    source_id              VARCHAR(500),
    source                 VARCHAR(500),
    link                   VARCHAR(500),
    make                   VARCHAR(255),
    model                  VARCHAR(255),
    version                VARCHAR(255),
    price                  INT,
    currency               VARCHAR(8),
    mileage                INT,
    mileage_unit           VARCHAR(8),
    first_registration     BIGINT,
    power                  SMALLINT,
    power_unit             VARCHAR(8),
    offer_type             VARCHAR(255),
    previous_owners_number SMALLINT,
    transmission_type      VARCHAR(126),
    fuel_type              VARCHAR(126),
    fuel_consumption       SMALLINT,
    co2_emission           SMALLINT,
    PRIMARY KEY (source_id, source)
);