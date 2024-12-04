package utils

var MESSAGE_QUEUE_NAME string  = "BlogQueue"
var CREATE_OPS string = "create"
var UPDATE_OPS string = "update"
var DELETE_OPS string = "delete"
var GET_OPS string = "fetch"
var BLOG_COLLECTION = "Blogs"

// collection names
var ALL_TAG string = "All Tags"
var BLOG_R_TAG string = "Blog_R_Tag"

// redis key names
// var REDIS_BLOG_COLLECTION = "Blogs"

// kafka
var KAFKA_TOPIC = "mongo-to-elastic-changes"
var KAFKA_BROKER = "localhost:9092";
var ELASTIC_INDEX_NAME = "mongo-events"
var NUMBER_OF_RETRIES = 2;