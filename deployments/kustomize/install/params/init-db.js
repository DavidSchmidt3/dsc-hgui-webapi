const mongoHost = process.env.HGUI_API_MONGODB_HOST;
const mongoPort = process.env.HGUI_API_MONGODB_PORT;

const mongoUser = process.env.HGUI_API_MONGODB_USERNAME;
const mongoPassword = process.env.HGUI_API_MONGODB_PASSWORD;

const database = process.env.HGUI_API_MONGODB_DATABASE;
const collection = process.env.HGUI_API_MONGODB_COLLECTION;

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while (true) {
  try {
    connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
    break;
  } catch (exception) {
    print(`Cannot connect to mongoDB: ${exception}`);
    print(`Will retry after ${retrySeconds} seconds`);
    sleep(retrySeconds * 1000);
  }
}

// if database and collection exists, exit with success - already initialized
const databases = connection.getDBNames();
if (databases.includes(database)) {
  const dbInstance = connection.getDB(database);
  collections = dbInstance.getCollectionNames();
  if (collections.includes(collection)) {
    print(`Collection '${collection}' already exists in database '${database}'`);
    process.exit(0);
  }
}

// initialize
// create database and collection
const db = connection.getDB(database);
db.createCollection(collection);

// create indexes
db[collection].createIndex({ "id": 1 });

let result = db[collection].insertMany([
  {
    "id": "774d8d43-3984-4bf0-9261-e360b61213ec",
    "patientName": "John Doe",
    "title": "Test problem",
    "description": "Test description",
    "severity": "low",
    "createdAt": new Date(),
    "email": "test@test.com"
  },
]);

if (result.writeError) {
  console.error(result);
  print(`Error when writing the data: ${result.errmsg}`);
}

// exit with success
process.exit(0);