package utils

// mongodb URL
var MONGO_DB_CONN_URL = "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0"

// number of limits per page
var LIMIT int = 3;

// CRUD consts
var MESSAGE_QUEUE_NAME string  = "BlogQueue"
var CREATE_OPS string = "create"
var UPDATE_OPS string = "update"
var DELETE_OPS string = "delete"
var GET_OPS string = "fetch"

// Tag available
var Marvel_tag string = "Marvel"
var Rpg_tag string = "RPG"
var Fps_tag  string = "FPS"
var Sony_tag string = "Sony"
var Ps5_tag string = "PS5"
var Dc_tag string = "DC"

//elasticsearch const
var TITLE string = "title";
var EXCERPT string ="excerpt";
var ELASTIC_PORT string = "http://localhost:9200"
var ES_BLOG string = "blog";


// collection names
var ALL_TAG string = "All Tags"
var BLOG_R_TAG string = "Blog_R_Tag"
var BLOG_COLLECTION = "Blogs"
var PS5_COLLECTION = "PS5-Collection"
var RPG_COLLECTION = "RPG-Collection"
var FPS_COLLECTION = "FPS-Collection"
var SONY_COLLECTION = "SONY-Collection"


// kafka
var KAFKA_TOPIC = "mongo-to-elastic-changes"
var KAFKA_BROKER = "localhost:9096";
var ELASTIC_INDEX_NAME = "mongo-events"
var NUMBER_OF_RETRIES = 2;