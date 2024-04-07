To use the importer, first set the config.ini file as follows:

uri: Your mongodb connection string
database: Your mongodb database name
collection: The collection you want to use
file_path: The path of the file you will be loading

Currently, the importer only supports loading one file at a time
That file must be contain a valid json array, where each item is a json object itself.
Each of that json objects will be loaded in the collection as a separate document
