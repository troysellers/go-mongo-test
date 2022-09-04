# Mongo Sync

Simple go util that looks accepts an [IMDB list id](https://imdb-api.com/api/#IMDbList-header) and either inserts or updates the movie found in the sample Mongo Atlas db sample_mflix.movies. 

## usage 

get the code
```
> git clone git@github.com:troysellers/go-mongo-test.git
```

Create a .env file 
```
> cp .env.sample .env
```

Open and set the required environment

Run
```
> go run mongo.go --list=<imdb-List-Id>
```

