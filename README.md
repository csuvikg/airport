#  World of Airports
The application pulls data from a Cloudant database of airports based on the provided
latitude, longitude and radius and lists them on a simplistic UI based on their distance
to the provided center.

## Running the application
The application can be run with either Docker or Go tools.
To run with docker on port `3003`, execute the following commands from the project root:
```bash
docker build -t airports .
docker run -dp 3003:3003 airports
```

## Using the application
The running application can be accessed [in the browser](http://localhost:3003). Fill in
the latitude and longitude in degrees, for example 47.503914 and 19.061767 for the
IBM Budapest Lab. Provide a search radius in km. On clicking Search or hitting enter,
the results are shown in a table.
