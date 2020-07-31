create table CovidInfo
(
    Country        varchar(255)                       not null,
    CountryCode    varchar(255)                       null,
    Slug           varchar(255)                       null,
    NewConfirmed   int                                null,
    TotalConfirmed int                                null,
    NewDeaths      int                                null,
    TotalDeaths    int                                null,
    NewRecovered   int                                null,
    TotalRecovered int                                null,
    last_updated   datetime default CURRENT_TIMESTAMP null
);


