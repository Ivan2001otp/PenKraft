import subprocess
import sys
import time


def start_mongo() :
	# // db path : E:\GO - Hub[Backend]\Full stack Projects\PenKraft\PenCraft-Backend

    try:
        command = ["mongod","--replSet","rs0", "--bind_ip","localhost","--port","27017","--dbpath","E:\GO - Hub[Backend]\data\db","--logpath","E:\GO - Hub[Backend]\data\logs\mongod.log"]

        print("Starting mongoDB from scripting...")
        process = subprocess.Popen(command, stdout=subprocess.PIPE,stderr=subprocess.PIPE)

        #waiting mongodb to start..
        time.sleep(4)

        print("Communicating...")
        stdout, stderr = process.communicate()
        print("Done with communication..")
        
        if process.returncode != 0 :
            print(f"Error starting mongoDB : {stderr.decode()}")
        else:
            print(f"Mongodb started successfully.{stdout.decode()}")

    except Exception as e:
        print(f"An error occured while trying to start MongoDB : {e}")


if __name__ == "__main__" :
    start_mongo()