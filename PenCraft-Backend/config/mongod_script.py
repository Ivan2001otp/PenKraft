import subprocess
import sys
import time
import os

MONGO_DATA_DIR = "E:/GO - Hub[Backend]/data/db"
MONGO_LOG_DIR = "E:/GO - Hub[Backend]/data/logs"
MONGO_PORT1 = 27017
MONGO_PORT2 = 27018
MONGO_PORT3 = 27019
REPLICA_SET_NAME = "rs0"

def create_directories() :
    print("Creating directories for mongodb data and dlogs !")
   
    os.makedirs(f"{MONGO_DATA_DIR}/db1",exist_ok=True)
    os.makedirs(f"{MONGO_DATA_DIR}/db2",exist_ok=True)
    os.makedirs(f"{MONGO_DATA_DIR}/db3",exist_ok=True)
    os.makedirs(MONGO_LOG_DIR, exist_ok=True)


def create_mongo_config_files() :
    print("Creating db configf files...")


    with open("mongod1.conf","w") as f:
        f.write(f""" 
net:
  port: {MONGO_PORT1}
  bindIp: 127.0.0.1

storage:
  dbPath: {MONGO_DATA_DIR}/db1

replication:
  replSetName: {REPLICA_SET_NAME}
""")
        
        #configuration for second instance.(secondary)
    with open("mongod2.conf","w") as f:
            f.write(f"""
net:
  port:{MONGO_PORT2}
  bindIp: 127.0.0.1

storage:
  dbPath: {MONGO_DATA_DIR}/db2

replication:
  replSetName: {REPLICA_SET_NAME}
""")
            
            #configuring for third instance
    with open("mongod3.conf","w") as f:
                f.write(f"""
net:
  port: {MONGO_PORT3}
  bindIp: 127.0.0.1

storage:
  dbPath: {MONGO_DATA_DIR}/db3

replication:
  replSetName: {REPLICA_SET_NAME}
""")
        
# start the mongodb instances
def start_mongo_instances() :
    print("starting mongo db instances")
    # removing --fork flag,because its not supported in windows unlike  macos, unix os.
    #start instance1
    try :
            p1 = subprocess.run(["mongod","--config","mongod1.conf","--logpath",f"{MONGO_LOG_DIR}/log1/mongod1.log"],check=True)
            print("instance1 started!")
            print(f"p1 Success : {p1.stdout}")


            #start instance2
            p2 = subprocess.run(["mongod","--config","mongod2.conf","--logpath",f"{MONGO_LOG_DIR}/log2/mongod2.log"],check=True)
            print("instance2 started!")
            print(f"p2 Success : {p2.stdout}")

            #start instance3
            p3 = subprocess.run(["mongod","--config","mongod3.conf","--logpath",f"{MONGO_LOG_DIR}/log3/mongod3.log"],check=True)
            print("instance3 started!")
            print(f"p3 Success : {p3.stdout}")

            print("all 3 replica sets  started successfully !")
    except subprocess.CalledProcessError as e:
         print("stderr: ",e.stderr)
         print("stdout: ",e.stdout)
         print(f"Error : {e}")

def initialize_replica_sets() :
    print("Initializing mongodb replica sets..")
    #give some time for mongodb to start
    time.sleep(10)

    #connect the first inst, and initialize replicaset
    result = subprocess.run(["mongo","--port",str(MONGO_PORT1),"--eval","""
rs.initiate();
rs.add('localhost:27018');
rs.add('localhost:27019');
print('Replica Set Initialized:');
print(rs.status());
"""],capture_output=True,text=True)
    
    if result.returncode != 0:
        print(f"Error initializing replica sets: {result.stderr}")
        return
    print("Replica set initialized successfully !")


# checkthe replica set status
def check_replicaset_status() :
    print("Checking replica set status")
    result = subprocess.run(["mongo","--port",str(MONGO_PORT1),"--eval","--printjson(rs.status())"],capture_output=True,text=True)
    if result.returncode != 0:
        print(f"Error checking replica set status : {result.stderr}")
        return
    print(result.stdout)

def main() :
    #create directories
    create_directories()

    #create mongodb config files
    create_mongo_config_files()

    #start mongodb instance
    start_mongo_instances()

    #initialize the replica set
    initialize_replica_sets()

    check_replicaset_status()

    print("mongo db replica set setup is successfully completed.")

if __name__ == "__main__" :
    main()