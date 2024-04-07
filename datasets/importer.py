import pymongo
import json
import configparser
from tqdm import tqdm

try:
    # Read Config file
    config = configparser.ConfigParser()
    config.read('config.ini')

    uri = config['mongodb']['uri']
    db_name = config['mongodb']['database']
    collection_name = config['mongodb']['collection']
    source_file_name = config['source']['file_path']

    print('Uri:' + uri)
    print('Database name:' + db_name)
    print('Collection name: ' + collection_name)
    print('File: ' + source_file_name)

    # Connect to Cluster
    client = pymongo.MongoClient(uri)
    print('Client connected')

    # Get the database and collection
    db = client[db_name]
    collection = db[collection_name]
    print('Db and collection initialized')

    # Read file
    with open(source_file_name, 'r') as json_file:
        data = json.load(json_file)

    # Insert data
    data_size = len(data)
    inp = input('Data size: ' + str(data_size) + '. Proceed?')

    with tqdm(total=data_size) as pbar:
        for document in data:
            collection.insert_one(document)
            print('Document ' + document.get('title') + ' inserted')
            pbar.update(1)
            print("\r", end="")

    print('Import done!')
except Exception as e:
    print('An error has ocurred during import: ' + e)

finally:
    client.close()

