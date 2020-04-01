#!/bin/bash

echo "Importing new data"
cd COVID-19
git fetch origin
git rebase
cd ..
cp COVID-19/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv ./confirmed.csv 
cp COVID-19/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv ./recovered.csv 
cp COVID-19/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv ./deaths.csv
go run cvsreader.go >covidconverted.csv 
cd csv-to-influxdb
python3 csv-to-influxdb.py --create --dbname covid19 --input ../covidconverted.csv -d ';' --tagcolumns "country,province" --fieldcolumns "lat,lon,confirmed,recovered,current,deaths,confirmed-increase" --timecolumn date --timeformat '%m/%d/%y' --metricname confirmed
cd ..
echo "Done"
