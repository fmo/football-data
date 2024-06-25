# Football Data

Football Data command line tool. Getting data from API Football and publishing it to the Kafka.

## Environment Variables

The application requires the following environment variables:

- `RAPID_API_KEY`: Your RapidAPI key.
- `ENVIRONMENT`: The environment in which the application is running. For local development, this should be `local`.
- `KAFKA_USERNAME`: Your Kafka username.
- `KAFKA_PASSWORD`: Your Kafka password.
- `KAFKA_HOST`: The host address of your Kafka instance.
- `KAFKA_TOPIC_FIXTURE`: The Kafka topic for fixtures.
- `KAFKA_TOPIC_TEAMS`: The Kafka topic for teams.
- `KAFKA_TOPIC_STANDING`: The Kafka topic for standings.

You can set these environment variables in a `.env` file in the root of the project. Use the provided `.env.dist` file as a template.

## Download binary

Download the binary from releases https://github.com/fmo/football-data/releases

## Run

To run the application, first ensure that all environment variables are set, then execute the binary:

```
./footballDataApp games --leagueId 203 --teamId 611
```

## Response

```
{
    "awayTeam":"İstanbulspor",
    "gameDate":"2024-05-26T16:00:00+00:00",
    "homeTeam":"Fenerbahce",
    "leagueName":"Süper Lig",
    "level":"info",
    "msg":"Game fetched!",
    "round":"Regular Season - 38",
    "score":"6-0",
    "time":"2024-06-25T21:13:58+02:00"
}
```

## Leagues

| League               | Id  |
|----------------------|-----|
| turkish-super-league | 203 |
| ucl                  | 2   |
| premier-league       | 39  |
| bundesliga           | 78  |
| la-liga              | 140 |
| serie-a              | 135 |
| uefa-conference      | 848 |
| uefa-europa-league   | 3   |
| euro-championship    | 4   |
| friendlies           | 10  |


## Teams

| Team        | Id  |
|-------------|-----|
| Real Madrid | 541 |
| Fenerbahce  | 611 |
| Turkiye     | 777 |
| Germany     | 25  |
| England     | 10  |
